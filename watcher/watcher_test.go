package watcher

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWatcher_DetectsFileChanges(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()

	// Create initial SQL file
	sqlFile := filepath.Join(tempDir, "test.sql")
	err := os.WriteFile(sqlFile, []byte("SELECT 1"), 0644)
	require.NoError(t, err)

	// Create watcher
	w, err := New(tempDir, 50*time.Millisecond)
	require.NoError(t, err)
	defer w.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	events, errs := w.Start(ctx)

	// Modify the file
	time.Sleep(100 * time.Millisecond)
	err = os.WriteFile(sqlFile, []byte("SELECT 2"), 0644)
	require.NoError(t, err)

	// Wait for event
	select {
	case event := <-events:
		assert.Len(t, event.Files, 1)
		assert.Equal(t, sqlFile, event.Files[0])
	case err := <-errs:
		t.Fatalf("unexpected error: %v", err)
	case <-ctx.Done():
		t.Fatal("timeout waiting for event")
	}
}

func TestWatcher_IgnoresNonSQLFiles(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()

	// Create watcher
	w, err := New(tempDir, 50*time.Millisecond)
	require.NoError(t, err)
	defer w.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	events, _ := w.Start(ctx)

	// Create a non-SQL file
	time.Sleep(100 * time.Millisecond)
	txtFile := filepath.Join(tempDir, "test.txt")
	err = os.WriteFile(txtFile, []byte("hello"), 0644)
	require.NoError(t, err)

	// Should not receive any events
	select {
	case event := <-events:
		t.Fatalf("unexpected event for non-SQL file: %v", event)
	case <-ctx.Done():
		// Expected - no events for non-SQL files
	}
}

func TestWatcher_DebouncesBatchedChanges(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()

	// Create initial SQL files
	file1 := filepath.Join(tempDir, "test1.sql")
	file2 := filepath.Join(tempDir, "test2.sql")
	err := os.WriteFile(file1, []byte("SELECT 1"), 0644)
	require.NoError(t, err)
	err = os.WriteFile(file2, []byte("SELECT 2"), 0644)
	require.NoError(t, err)

	// Create watcher with longer debounce
	w, err := New(tempDir, 100*time.Millisecond)
	require.NoError(t, err)
	defer w.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	events, errs := w.Start(ctx)

	// Modify both files quickly
	time.Sleep(50 * time.Millisecond)
	err = os.WriteFile(file1, []byte("SELECT 1 updated"), 0644)
	require.NoError(t, err)
	time.Sleep(10 * time.Millisecond)
	err = os.WriteFile(file2, []byte("SELECT 2 updated"), 0644)
	require.NoError(t, err)

	// Should receive one batched event
	select {
	case event := <-events:
		assert.Len(t, event.Files, 2)
	case err := <-errs:
		t.Fatalf("unexpected error: %v", err)
	case <-ctx.Done():
		t.Fatal("timeout waiting for event")
	}
}

func TestWatcher_DetectsNewFiles(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()

	// Create watcher
	w, err := New(tempDir, 50*time.Millisecond)
	require.NoError(t, err)
	defer w.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	events, errs := w.Start(ctx)

	// Create a new SQL file
	time.Sleep(100 * time.Millisecond)
	newFile := filepath.Join(tempDir, "new.sql")
	err = os.WriteFile(newFile, []byte("SELECT 1"), 0644)
	require.NoError(t, err)

	// Wait for event
	select {
	case event := <-events:
		assert.Len(t, event.Files, 1)
		assert.Equal(t, newFile, event.Files[0])
	case err := <-errs:
		t.Fatalf("unexpected error: %v", err)
	case <-ctx.Done():
		t.Fatal("timeout waiting for event")
	}
}

func TestWatcher_IgnoresTempFiles(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()

	// Create watcher
	w, err := New(tempDir, 50*time.Millisecond)
	require.NoError(t, err)
	defer w.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	events, _ := w.Start(ctx)

	// Create temp files (should be ignored)
	time.Sleep(100 * time.Millisecond)
	err = os.WriteFile(filepath.Join(tempDir, ".test.sql"), []byte("hidden"), 0644)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(tempDir, "test.sql~"), []byte("backup"), 0644)
	require.NoError(t, err)

	// Should not receive any events
	select {
	case event := <-events:
		t.Fatalf("unexpected event for temp file: %v", event)
	case <-ctx.Done():
		// Expected - no events for temp files
	}
}
