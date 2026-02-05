// Package watcher provides file system monitoring for SQL files.
// It watches a directory for changes and debounces events to avoid
// excessive recompilation during rapid file modifications.
package watcher

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Event represents a batch of file changes after debouncing.
type Event struct {
	// Files contains the paths of changed SQL files.
	Files []string
}

// Watcher monitors a directory for SQL file changes.
type Watcher struct {
	sourcePath string
	fsWatcher  *fsnotify.Watcher

	debounceInterval time.Duration
	mu               sync.Mutex
	pending          map[string]struct{}
	timer            *time.Timer
	timerSeq         uint64 // sequence number to invalidate stale timer callbacks

	events chan Event
	errors chan error
}

// New creates a new Watcher for the given source directory.
func New(sourcePath string, debounceInterval time.Duration) (*Watcher, error) {
	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("failed to create file watcher: %w", err)
	}

	w := &Watcher{
		sourcePath:       sourcePath,
		fsWatcher:        fsWatcher,
		debounceInterval: debounceInterval,
		pending:          make(map[string]struct{}),
		events:           make(chan Event, 1), // buffered to prevent blocking
		errors:           make(chan error, 1), // buffered to prevent goroutine leak
	}

	return w, nil
}

// Start begins watching for file changes. Returns channels for events and errors.
// The watcher will stop when the context is cancelled.
func (w *Watcher) Start(ctx context.Context) (<-chan Event, <-chan error) {
	if err := w.fsWatcher.Add(w.sourcePath); err != nil {
		// Use buffered channels to prevent goroutine leak
		errCh := make(chan error, 1)
		errCh <- fmt.Errorf("failed to watch directory %s: %w", w.sourcePath, err)
		close(errCh)

		eventCh := make(chan Event)
		close(eventCh)

		return eventCh, errCh
	}

	go w.run(ctx)

	return w.events, w.errors
}

func (w *Watcher) run(ctx context.Context) {
	defer close(w.events)
	defer close(w.errors)
	defer w.fsWatcher.Close()

	for {
		select {
		case <-ctx.Done():
			// Stop any pending timer before flushing
			w.mu.Lock()
			if w.timer != nil {
				w.timer.Stop()
			}
			w.timerSeq++ // invalidate any pending callback
			w.mu.Unlock()

			w.flushLocked()
			return

		case event, ok := <-w.fsWatcher.Events:
			if !ok {
				return
			}

			if !w.shouldProcess(event) {
				continue
			}

			w.addPending(event.Name)

		case err, ok := <-w.fsWatcher.Errors:
			if !ok {
				return
			}
			select {
			case w.errors <- fmt.Errorf("file watcher error: %w", err):
			case <-ctx.Done():
				return
			}
		}
	}
}

// shouldProcess returns true if the event should trigger recompilation.
func (w *Watcher) shouldProcess(event fsnotify.Event) bool {
	// Only watch .sql files
	if !strings.HasSuffix(event.Name, ".sql") {
		return false
	}

	// Ignore temporary files from editors
	base := filepath.Base(event.Name)
	if strings.HasPrefix(base, ".") || strings.HasSuffix(base, "~") {
		return false
	}

	// Only process write, create, and remove events
	return event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove) != 0
}

// addPending adds a file to the pending set and resets the debounce timer.
func (w *Watcher) addPending(path string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.pending[path] = struct{}{}

	if w.timer != nil {
		w.timer.Stop()
	}

	// Increment sequence to invalidate any stale callbacks
	w.timerSeq++
	seq := w.timerSeq

	w.timer = time.AfterFunc(w.debounceInterval, func() {
		w.flushIfSeq(seq)
	})
}

// flushIfSeq sends pending files only if the sequence number matches.
// This prevents stale timer callbacks from sending duplicate events.
func (w *Watcher) flushIfSeq(seq uint64) {
	w.mu.Lock()
	if w.timerSeq != seq {
		w.mu.Unlock()
		return
	}
	w.mu.Unlock()

	w.flushLocked()
}

// flushLocked sends the pending files as an event and clears the pending set.
func (w *Watcher) flushLocked() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if len(w.pending) == 0 {
		return
	}

	files := make([]string, 0, len(w.pending))
	for path := range w.pending {
		files = append(files, path)
	}

	w.pending = make(map[string]struct{})

	// Non-blocking send with buffered channel
	select {
	case w.events <- Event{Files: files}:
	default:
		// Channel is full, this shouldn't happen with buffered channel
		// but log if it does to avoid silent data loss
		fmt.Printf("warn: watcher event channel full, some changes may be missed\n")
	}
}

// Close stops the watcher and releases resources.
func (w *Watcher) Close() error {
	return w.fsWatcher.Close()
}
