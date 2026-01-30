package config

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"path/filepath"
	"strings"

	"github.com/ioj/sqlty/db"
	"github.com/spf13/viper"
)

type Paths struct {
	Source    string
	Output    string
	Cache     string
	Templates string
}

type Config struct {
	DBURL        string
	Paths        Paths
	DefaultTypes string
	PackageName  string
	Types        []db.PGTypeTranslation
}

// Load loads configuration from the default sqlty.yaml file.
func Load() (*Config, error) {
	return LoadFrom("")
}

// LoadFrom loads configuration from the specified file path.
// If configPath is empty, it searches for sqlty.yaml in the current directory.
func LoadFrom(configPath string) (*Config, error) {
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("sqlty")
		viper.SetConfigType("yaml")
	}

	viper.SetDefault("Paths.Source", ".")
	viper.SetDefault("Paths.Output", ".")
	viper.SetDefault("Paths.Cache", ".sqlty")
	viper.SetDefault("Paths.Templates", "")
	viper.SetDefault("DefaultTypes", "")
	viper.SetDefault("PackageName", "")
	viper.SetDefault("DBURL", "")

	viper.SetEnvPrefix("sqlty")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	cfg := &Config{}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

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
