# Watcher Package

The `watcher` package provides file system monitoring for SQL files. It watches
a directory for changes and debounces events to avoid excessive recompilation
during rapid file modifications.

## Overview

```
File System Events (fsnotify)
        │
        ▼
┌─────────────────┐
│     Filter      │  Only .sql files, ignores temp files
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│    Debouncer    │  Batches events within 100ms window
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Event Channel  │  Emits batched file change events
└─────────────────┘
```

The watcher:

- Monitors a directory for SQL file changes using `fsnotify`
- Filters events to only `.sql` files (ignores temp files like `.file.sql` or
  `file.sql~`)
- Debounces rapid changes to batch multiple saves into a single event
- Provides clean shutdown via context cancellation

## Files

| File              | Purpose                                   |
| ----------------- | ----------------------------------------- |
| `watcher.go`      | Core watcher logic and event handling     |
| `watcher_test.go` | Unit tests for file watching and debounce |

## Usage

```go
// Create watcher with 100ms debounce interval
w, err := watcher.New("/path/to/sql/files", 100*time.Millisecond)
if err != nil {
    return err
}
defer w.Close()

// Start watching with context for cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

events, errors := w.Start(ctx)

// Process events
for {
    select {
    case event := <-events:
        fmt.Printf("Files changed: %v\n", event.Files)
        // Recompile changed files...

    case err := <-errors:
        fmt.Printf("Watch error: %v\n", err)

    case <-ctx.Done():
        return nil
    }
}
```

## Types

### Event

Represents a batch of file changes after debouncing.

```go
type Event struct {
    Files []string  // Paths of changed SQL files
}
```

### Watcher

Monitors a directory for SQL file changes.

```go
type Watcher struct {
    // ... internal fields
}

func New(sourcePath string, debounceInterval time.Duration) (*Watcher, error)
func (w *Watcher) Start(ctx context.Context) (<-chan Event, <-chan error)
func (w *Watcher) Close() error
```

## Event Filtering

The watcher only processes events that match these criteria:

| Criterion       | Rule                                               |
| --------------- | -------------------------------------------------- |
| File extension  | Must end with `.sql`                               |
| Temp files      | Ignores files starting with `.` or ending with `~` |
| Event type      | Only `Write`, `Create`, and `Remove` events        |

## Debouncing

When files change rapidly (e.g., editor auto-save, multiple files saved at
once), the debouncer batches events:

```
t=0ms    file1.sql modified  → start 100ms timer
t=20ms   file2.sql modified  → reset timer, add to batch
t=50ms   file1.sql modified  → reset timer (already in batch)
t=150ms  timer fires         → emit Event{Files: [file1.sql, file2.sql]}
```

This prevents:

- Compilation storms during rapid edits
- Partial file reads from editors still writing
- Duplicate recompilation of the same file

## Concurrency Safety

The watcher uses:

- **Mutex** to protect the pending file set during concurrent access
- **Sequence numbers** to invalidate stale timer callbacks
- **Buffered channels** to prevent goroutine leaks on startup failure

## Integration with SQLty

The watcher is used by SQLty's `--watch` mode:

```bash
sqlty --watch
```

This enables automatic recompilation when SQL files change:

1. Initial full compilation runs
2. Watcher monitors the source directory
3. On file changes, only changed files are recompiled
4. Shared files (enums, composite types) are always regenerated
5. `goimports` formats the output
