package config

import (
	"errors"
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

func Load() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("sqlty")
	viper.SetConfigType("yaml")

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
		return nil, err
	}

	cfg := &Config{}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if cfg.DBURL == "" {
		return nil, errors.New("dbUrl is required")
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
