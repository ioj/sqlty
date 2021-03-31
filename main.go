package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/ioj/sqlty/compiler"
	"github.com/ioj/sqlty/db"
	"github.com/ioj/sqlty/generator"
	"github.com/ioj/sqlty/stmt"
)

type Config struct {
	DBURL        string                 `yaml:"dbUrl"`
	Dir          string                 `yaml:"dir"`
	TemplateDir  string                 `yaml:"templateDir"`
	DefaultTypes string                 `yaml:"defaultTypes"`
	Types        []db.PGTypeTranslation `yaml:"types"`
}

func (c *Config) Validate() error {
	if c.DBURL == "" {
		return errors.New("dbUrl is required")
	}

	if c.Dir == "" {
		c.Dir = "."
	}

	return nil
}

func compiledir(cfg *Config) error {
	sqlFiles, err := filepath.Glob(path.Join(cfg.Dir, "*.sql"))
	if err != nil {
		return err
	}

	if len(sqlFiles) == 0 {
		return fmt.Errorf("no *.sql files to compile in %v", cfg.Dir)
	}

	defaulttypesfile, err := os.Open(cfg.DefaultTypes)
	if err != nil {
		return err
	}

	defaulttypes := &db.PGTypeTranslationsFile{}
	if err := yaml.NewDecoder(defaulttypesfile).Decode(defaulttypes); err != nil {
		return err
	}

	types := append(defaulttypes.Types, cfg.Types...)

	ctx := context.Background()
	resolver, err := db.NewResolver(ctx, cfg.DBURL, types)
	if err != nil {
		return err
	}
	defer resolver.Close()

	gen, err := generator.New(path.Join(cfg.TemplateDir, "*.go.tpl"))
	if err != nil {
		return err
	}

	enums := &stmt.Enums{PackageName: "sql", Enums: resolver.Enums()}
	if err := gen.Enums(cfg.Dir, enums); err != nil {
		return err
	}

	var compileTime, resolveTime, generateTime time.Duration

	for _, fname := range sqlFiles {
		t1 := time.Now()
		q, err := compiler.CompileFile(fname)
		if err != nil {
			return err
		}

		t2 := time.Now()
		compileTime += t2.Sub(t1)

		params, returns, err := resolver.ResolveTypes(ctx, q.PreparedQuery(), q.NotNullArray())
		if err != nil {
			return fmt.Errorf("%v: %v", fname, err.Error())
		}

		t3 := time.Now()
		resolveTime += t3.Sub(t2)

		stmtq, err := q.StmtQuery("sql", params, returns)
		if err != nil {
			return err
		}

		if (returns == nil || len(returns.Params) == 0) && stmtq.ExecMode != stmt.ExecModeExec {
			fmt.Printf("warn: overriding %v exec mode to @exec\n", stmtq.Name)
			stmtq.ExecMode = stmt.ExecModeExec
		}

		fnameOut := strings.TrimSuffix(fname, path.Ext(fname)) + ".gen.go"
		if err := gen.Query(fnameOut, stmtq); err != nil {
			return err
		}

		t4 := time.Now()
		generateTime += t4.Sub(t3)
	}

	ct, err := resolver.CompositeTypes(ctx)
	if err != nil {
		return err
	}

	compositeTypes := &stmt.CompositeTypes{PackageName: "sql", Types: ct}
	if err := gen.CompositeTypes(cfg.Dir, compositeTypes); err != nil {
		return err
	}

	// t := time.Now()
	goimports := exec.Command("goimports", "-w", ".")
	goimports.Dir = path.Dir(cfg.Dir)
	if output, err := goimports.CombinedOutput(); err != nil {
		return fmt.Errorf("goimports error: %v, %v", err, string(output))
	}

	/*
		goimportsTime := time.Now().Sub(t)
		fmt.Printf("Processed files: %d\n", len(sqlFiles))
		fmt.Printf("Compile time: %.02fs\n", compileTime.Seconds())
		fmt.Printf("Resolve time: %.02fs\n", resolveTime.Seconds())
		fmt.Printf("Generate time: %.02fs\n", generateTime.Seconds())
		fmt.Printf("Goimports time: %.02fs\n", goimportsTime.Seconds())
	*/

	return nil
}

func main() {
	f, err := os.Open(".sqlty.yaml")
	if err != nil {
		log.Fatal(err)
	}

	cfg := &Config{}
	if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
		log.Fatal(err)
	}

	if err := cfg.Validate(); err != nil {
		log.Fatal(err)
	}

	if err := compiledir(cfg); err != nil {
		log.Fatal(err)
	}
}
