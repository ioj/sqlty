package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCacheFromFile_Empty(t *testing.T) {
	dir := t.TempDir()

	c, err := newCacheFromFile(dir)
	require.NoError(t, err)
	require.NotNil(t, c)
	assert.Empty(t, c.Items)
}

func TestNewCacheFromFile_Existing(t *testing.T) {
	dir := t.TempDir()
	content := `items:
  test.go: abc123
  other.go: def456
`
	err := os.WriteFile(filepath.Join(dir, cachefilename), []byte(content), 0644)
	require.NoError(t, err)

	c, err := newCacheFromFile(dir)
	require.NoError(t, err)
	assert.Equal(t, "abc123", c.Items["test.go"])
	assert.Equal(t, "def456", c.Items["other.go"])
}

func TestNewCacheFromFile_EmptyFile(t *testing.T) {
	dir := t.TempDir()
	err := os.WriteFile(filepath.Join(dir, cachefilename), []byte(""), 0644)
	require.NoError(t, err)

	// Empty YAML file causes decode error - this is expected behavior
	_, err = newCacheFromFile(dir)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode cache file")
}

func TestCacheSave(t *testing.T) {
	dir := t.TempDir()
	c := &cache{Items: map[string]string{"test.go": "hash123"}}

	err := c.save(dir)
	require.NoError(t, err)

	// Verify file was created
	data, err := os.ReadFile(filepath.Join(dir, cachefilename))
	require.NoError(t, err)
	assert.Contains(t, string(data), "test.go")
	assert.Contains(t, string(data), "hash123")
}

func TestCacheSave_CreatesDirectory(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "nested", "path")
	c := &cache{Items: map[string]string{"test.go": "hash"}}

	err := c.save(dir)
	require.NoError(t, err)

	_, err = os.Stat(filepath.Join(dir, cachefilename))
	require.NoError(t, err)
}

func TestCacheSave_Nil(t *testing.T) {
	var c *cache
	err := c.save(t.TempDir())
	require.NoError(t, err) // Should not error on nil cache
}

func TestCacheUpdate_NewEntry(t *testing.T) {
	c := &cache{Items: make(map[string]string)}
	params := map[string]string{"name": "test"}

	changed, err := c.update("test.go", params)
	require.NoError(t, err)
	assert.True(t, changed)
	assert.NotEmpty(t, c.Items["test.go"])
}

func TestCacheUpdate_SameParams(t *testing.T) {
	c := &cache{Items: make(map[string]string)}
	params := map[string]string{"name": "test"}

	// First update
	_, err := c.update("test.go", params)
	require.NoError(t, err)

	// Same params again
	changed, err := c.update("test.go", params)
	require.NoError(t, err)
	assert.False(t, changed)
}

func TestCacheUpdate_DifferentParams(t *testing.T) {
	c := &cache{Items: make(map[string]string)}

	// First update
	_, err := c.update("test.go", map[string]string{"name": "test1"})
	require.NoError(t, err)

	// Different params
	changed, err := c.update("test.go", map[string]string{"name": "test2"})
	require.NoError(t, err)
	assert.True(t, changed)
}

func TestCacheUpdate_NilCache(t *testing.T) {
	var c *cache

	// Nil cache should always report changed
	changed, err := c.update("test.go", "anything")
	require.NoError(t, err)
	assert.True(t, changed)
}

func TestCacheRoundTrip(t *testing.T) {
	dir := t.TempDir()

	// Create and save
	c1 := &cache{Items: map[string]string{
		"file1.go": "hash1",
		"file2.go": "hash2",
	}}
	err := c1.save(dir)
	require.NoError(t, err)

	// Load
	c2, err := newCacheFromFile(dir)
	require.NoError(t, err)

	assert.Equal(t, c1.Items, c2.Items)
}
