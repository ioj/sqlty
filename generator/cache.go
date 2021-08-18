package generator

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
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

	c := &cache{}
	if err := yaml.NewDecoder(f).Decode(c); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *cache) save(dir string) error {
	if c == nil {
		return nil
	}

	err := os.Mkdir(dir, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	fname := path.Join(dir, cachefilename)
	f, err := os.Create(fname)
	if err != nil {
		return err
	}

	return yaml.NewEncoder(f).Encode(c)
}

// update updates the hash for given params. Returns true if hash is different than
// the old one.
func (c *cache) update(key string, params interface{}) (bool, error) {
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
