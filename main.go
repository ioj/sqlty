package main

import (
	"bytes"
	"context"
	_ "embed"
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
	"github.com/ioj/sqlty/config"
	"github.com/ioj/sqlty/db"
	"github.com/ioj/sqlty/generator"
	"github.com/ioj/sqlty/stmt"
)

//go:embed db/types.yaml
var defaultTypes []byte

func compiledir(cfg *config.Config) error {
	sqlFiles, err := filepath.Glob(path.Join(cfg.Paths.Source, "*.sql"))
	if err != nil {
		return err
	}

	if len(sqlFiles) == 0 {
		return fmt.Errorf("no *.sql files to compile in %v", cfg.Paths.Source)
	}

	defaulttypes := &db.PGTypeTranslationsFile{}
	if cfg.DefaultTypes == "" {
		if err := yaml.NewDecoder(bytes.NewReader(defaultTypes)).Decode(defaulttypes); err != nil {
			return err
		}
	} else {
		f, err := os.Open(cfg.DefaultTypes)
		if err != nil {
			return err
		}
		defer f.Close()

		if err := yaml.NewDecoder(f).Decode(defaulttypes); err != nil {
			return err
		}
	}

	types := append(defaulttypes.Types, cfg.Types...)

	ctx := context.Background()
	resolver, err := db.NewResolver(ctx, cfg.DBURL, types)
	if err != nil {
		return err
	}
	defer resolver.Close()

	gen, err := generator.New(cfg.Paths.Templates, cfg.Paths.Cache)
	if err != nil {
		return err
	}
	defer gen.Close()

	enums := &stmt.Enums{PackageName: cfg.PackageName, Enums: resolver.Enums()}
	if err := gen.Enums(cfg.Paths.Output, enums); err != nil {
		return err
	}

	if err := gen.DB(cfg.Paths.Output, &stmt.DB{PackageName: cfg.PackageName}); err != nil {
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

		stmtq, err := q.StmtQuery(cfg.PackageName, params, returns)
		if err != nil {
			return err
		}

		if (returns == nil || len(returns.Params) == 0) && stmtq.ExecMode != stmt.ExecModeExec {
			fmt.Printf("warn: overriding %v exec mode to @exec\n", stmtq.Name)
			stmtq.ExecMode = stmt.ExecModeExec
		}

		// source/filename.sql -> output/filename.gen.go
		fnameOut := strings.TrimSuffix(path.Base(fname), path.Ext(fname)) + ".gen.go"
		fnameOut = path.Join(cfg.Paths.Output, fnameOut)

		if err := gen.Query(q.Template(), fnameOut, stmtq); err != nil {
			return err
		}

		t4 := time.Now()
		generateTime += t4.Sub(t3)
	}

	ct, err := resolver.CompositeTypes(ctx)
	if err != nil {
		return err
	}

	compositeTypes := &stmt.CompositeTypes{PackageName: cfg.PackageName, Types: ct}
	if err := gen.CompositeTypes(cfg.Paths.Output, compositeTypes); err != nil {
		return err
	}

	goimports := exec.Command("goimports", "-w", ".")
	goimports.Dir = path.Dir(cfg.Paths.Output)
	if output, err := goimports.CombinedOutput(); err != nil {
		return fmt.Errorf("goimports error: %v, %v", err, string(output))
	}

	return nil
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	if err := compiledir(cfg); err != nil {
		log.Fatal(err)
	}
}
