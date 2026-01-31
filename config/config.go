package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ioj/sqlty/db"
	"gopkg.in/yaml.v3"
)

type Paths struct {
	Source    string `yaml:"source"`
	Output    string `yaml:"output"`
	Cache     string `yaml:"cache"`
	Templates string `yaml:"templates"`
}

type Config struct {
	DBURL        string                 `yaml:"dbUrl"`
	Paths        Paths                  `yaml:"paths"`
	DefaultTypes string                 `yaml:"defaultTypes"`
	PackageName  string                 `yaml:"packageName"`
	Types        []db.PGTypeTranslation `yaml:"types"`
}

// Load loads configuration from the default sqlty.yaml file.
func Load() (*Config, error) {
	return LoadFrom("")
}

// LoadFrom loads configuration from the specified file path.
// If configPath is empty, it searches for sqlty.yaml in the current directory.
func LoadFrom(configPath string) (*Config, error) {
	if configPath == "" {
		configPath = "sqlty.yaml"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	cfg := &Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	// Apply defaults
	if cfg.Paths.Source == "" {
		cfg.Paths.Source = "."
	}
	if cfg.Paths.Output == "" {
		cfg.Paths.Output = "."
	}
	if cfg.Paths.Cache == "" {
		cfg.Paths.Cache = ".sqlty"
	}

	// Apply environment variable overrides (SQLTY_ prefix)
	applyEnvOverrides(cfg)

	if cfg.DBURL == "" {
		return nil, errors.New("dbUrl is required")
	}

	// Validate database URL format
	parsedURL, err := url.Parse(cfg.DBURL)
	if err != nil {
		return nil, fmt.Errorf("invalid dbUrl: %w", err)
	}
	if parsedURL.Scheme != "postgres" && parsedURL.Scheme != "postgresql" {
		return nil, errors.New("dbUrl must be a PostgreSQL connection string (postgres:// or postgresql://)")
	}

	if cfg.PackageName == "" {
		absOutput, err := filepath.Abs(cfg.Paths.Output)
		if err != nil {
			return nil, err
		}

		cfg.PackageName = strings.ToLower(path.Base(absOutput))
	}

	return cfg, nil
}

// applyEnvOverrides applies SQLTY_* environment variable overrides to the config.
func applyEnvOverrides(cfg *Config) {
	if v := os.Getenv("SQLTY_DBURL"); v != "" {
		cfg.DBURL = v
	}
	if v := os.Getenv("SQLTY_PATHS_SOURCE"); v != "" {
		cfg.Paths.Source = v
	}
	if v := os.Getenv("SQLTY_PATHS_OUTPUT"); v != "" {
		cfg.Paths.Output = v
	}
	if v := os.Getenv("SQLTY_PATHS_CACHE"); v != "" {
		cfg.Paths.Cache = v
	}
	if v := os.Getenv("SQLTY_PATHS_TEMPLATES"); v != "" {
		cfg.Paths.Templates = v
	}
	if v := os.Getenv("SQLTY_PACKAGENAME"); v != "" {
		cfg.PackageName = v
	}
	if v := os.Getenv("SQLTY_DEFAULTTYPES"); v != "" {
		cfg.DefaultTypes = v
	}
}
