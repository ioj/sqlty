package generator

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

const cachefilename = "cache.yaml"

// cache is used to check if a given output file needs to be
// regenerated. For every render request, generator calculates
// a checksum of passed parameters (stmt.Query) and regenerates
// the file only if the previous checksum is different.
type cache struct {
	Items map[string]string `yaml:"items"`
}

func newCacheFromFile(dir string) (*cache, error) {
	fname := path.Join(dir, cachefilename)
	f, err := os.Open(fname)
	if os.IsNotExist(err) {
		return &cache{Items: make(map[string]string)}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to open cache file: %w", err)
	}
	defer f.Close()

	c := &cache{}
	if err := yaml.NewDecoder(f).Decode(c); err != nil {
		return nil, fmt.Errorf("failed to decode cache file: %w", err)
	}

	if c.Items == nil {
		c.Items = make(map[string]string)
	}

	return c, nil
}

func (c *cache) save(dir string) error {
	if c == nil {
		return nil
	}

	// Use MkdirAll to handle nested paths
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	fname := path.Join(dir, cachefilename)
	tmpFile := fname + ".tmp"

	// Write to temp file first for atomic operation
	f, err := os.Create(tmpFile)
	if err != nil {
		return fmt.Errorf("failed to create temp cache file: %w", err)
	}

	if err := yaml.NewEncoder(f).Encode(c); err != nil {
		f.Close()
		os.Remove(tmpFile)
		return fmt.Errorf("failed to encode cache: %w", err)
	}

	if err := f.Close(); err != nil {
		os.Remove(tmpFile)
		return fmt.Errorf("failed to write cache file: %w", err)
	}

	// Atomic rename
	if err := os.Rename(tmpFile, fname); err != nil {
		os.Remove(tmpFile)
		return fmt.Errorf("failed to save cache file: %w", err)
	}

	return nil
}

// update updates the hash for given params. Returns true if hash is different than
// the old one.
func (c *cache) update(key string, params any) (bool, error) {
	if c == nil {
		// Always invalidate when cache is nil
		return true, nil
	}

	old := c.Items[key]

	bytes, err := json.Marshal(params)
	if err != nil {
		return false, err
	}

	h := md5.Sum(bytes)
	new := hex.EncodeToString(h[:])

	if old == new {
		return false, nil
	}

	c.Items[key] = new
	return true, nil
}
