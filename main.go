package main

import (
	"bytes"
	"context"
	_ "embed"
	"flag"
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

// Version information (can be set via ldflags)
var (
	version = "dev"
	commit  = "unknown"
)

//go:embed db/types.yaml
var defaultTypes []byte

var (
	showVersion = flag.Bool("version", false, "print version and exit")
	showHelp    = flag.Bool("help", false, "show help")
	verbose     = flag.Bool("verbose", false, "enable verbose output")
	configFile  = flag.String("config", "sqlty.yaml", "config file path")
	timeout     = flag.Duration("timeout", 30*time.Second, "database connection timeout")
)

func compiledir(cfg *config.Config, verbose bool, timeout time.Duration) error {
	sqlFiles, err := filepath.Glob(path.Join(cfg.Paths.Source, "*.sql"))
	if err != nil {
		return fmt.Errorf("failed to find SQL files: %w", err)
	}

	if len(sqlFiles) == 0 {
		return fmt.Errorf("no *.sql files to compile in %v", cfg.Paths.Source)
	}

	if verbose {
		fmt.Printf("Found %d SQL files to process\n", len(sqlFiles))
	}

	defaulttypes := &db.PGTypeTranslationsFile{}
	if cfg.DefaultTypes == "" {
		if err := yaml.NewDecoder(bytes.NewReader(defaultTypes)).Decode(defaulttypes); err != nil {
			return fmt.Errorf("failed to decode default types: %w", err)
		}
	} else {
		f, err := os.Open(cfg.DefaultTypes)
		if err != nil {
			return fmt.Errorf("failed to open custom types file %s: %w", cfg.DefaultTypes, err)
		}
		defer f.Close()

		if err := yaml.NewDecoder(f).Decode(defaulttypes); err != nil {
			return fmt.Errorf("failed to decode custom types file %s: %w", cfg.DefaultTypes, err)
		}
	}

	types := append(defaulttypes.Types, cfg.Types...)

	// Use context with timeout for database operations
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if verbose {
		fmt.Println("Connecting to database...")
	}

	resolver, err := db.NewResolver(ctx, cfg.DBURL, types)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer resolver.Close()

	gen, err := generator.New(cfg.Paths.Templates, cfg.Paths.Cache)
	if err != nil {
		return fmt.Errorf("failed to initialize generator: %w", err)
	}
	defer gen.Close()

	enums := &stmt.Enums{PackageName: cfg.PackageName, Enums: resolver.Enums()}
	if err := gen.Enums(cfg.Paths.Output, enums); err != nil {
		return fmt.Errorf("failed to generate enums: %w", err)
	}

	if err := gen.DB(cfg.Paths.Output, &stmt.DB{PackageName: cfg.PackageName}); err != nil {
		return fmt.Errorf("failed to generate db utilities: %w", err)
	}

	var compileTime, resolveTime, generateTime time.Duration

	for i, fname := range sqlFiles {
		if verbose {
			fmt.Printf("[%d/%d] Processing %s\n", i+1, len(sqlFiles), filepath.Base(fname))
		}

		t1 := time.Now()
		q, err := compiler.CompileFile(fname)
		if err != nil {
			if err == compiler.ErrEmptyFile {
				fmt.Printf("warn: %v is empty, ignoring\n", fname)
				continue
			}
			return fmt.Errorf("%s: %w", fname, err)
		}

		t2 := time.Now()
		compileTime += t2.Sub(t1)

		params, returns, err := resolver.ResolveTypes(ctx, q.PreparedQuery(), q.NotNullArray())
		if err != nil {
			return fmt.Errorf("%s: failed to resolve types: %w", fname, err)
		}

		t3 := time.Now()
		resolveTime += t3.Sub(t2)

		stmtq, err := q.StmtQuery(cfg.PackageName, params, returns)
		if err != nil {
			return fmt.Errorf("%s: %w", fname, err)
		}

		if (returns == nil || len(returns.Params) == 0) && stmtq.ExecMode != stmt.ExecModeExec {
			fmt.Printf("warn: %s has no return values, changing exec mode to @exec\n", stmtq.Name)
			stmtq.ExecMode = stmt.ExecModeExec
		}

		// source/filename.sql -> output/filename.gen.go
		fnameOut := strings.TrimSuffix(path.Base(fname), path.Ext(fname)) + ".gen.go"
		fnameOut = path.Join(cfg.Paths.Output, fnameOut)

		if err := gen.Query(q.Template(), fnameOut, stmtq); err != nil {
			return fmt.Errorf("%s: failed to generate query: %w", fname, err)
		}

		t4 := time.Now()
		generateTime += t4.Sub(t3)
	}

	ct, err := resolver.CompositeTypes(ctx)
	if err != nil {
		return fmt.Errorf("failed to resolve composite types: %w", err)
	}

	compositeTypes := &stmt.CompositeTypes{PackageName: cfg.PackageName, Types: ct}
	if err := gen.CompositeTypes(cfg.Paths.Output, compositeTypes); err != nil {
		return fmt.Errorf("failed to generate composite types: %w", err)
	}

	// Run goimports (non-fatal if it fails or is not installed)
	if _, err := exec.LookPath("goimports"); err != nil {
		fmt.Println("warn: goimports not found, generated code may need manual formatting")
	} else {
		goimports := exec.Command("goimports", "-w", ".")
		goimports.Dir = cfg.Paths.Output
		if output, err := goimports.CombinedOutput(); err != nil {
			fmt.Printf("warn: goimports failed: %v\n%s\n", err, string(output))
			fmt.Println("Generated code may need manual formatting")
		}
	}

	if verbose {
		fmt.Printf("\nTiming: compile=%v, resolve=%v, generate=%v, total=%v\n",
			compileTime.Round(time.Millisecond),
			resolveTime.Round(time.Millisecond),
			generateTime.Round(time.Millisecond),
			(compileTime + resolveTime + generateTime).Round(time.Millisecond))
		fmt.Printf("Generated %d query files\n", len(sqlFiles))
	}

	return nil
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "SQLty - SQL code generator for Go with PostgreSQL\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nConfiguration:\n")
		fmt.Fprintf(os.Stderr, "  Create a sqlty.yaml file with:\n")
		fmt.Fprintf(os.Stderr, "    dburl: postgres://user:pass@localhost:5432/mydb\n")
		fmt.Fprintf(os.Stderr, "    paths:\n")
		fmt.Fprintf(os.Stderr, "      source: ./queries\n")
		fmt.Fprintf(os.Stderr, "      output: ./db\n")
		fmt.Fprintf(os.Stderr, "    package: db\n")
	}

	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("sqlty version %s (commit %s)\n", version, commit)
		os.Exit(0)
	}

	// Check if config file exists
	if _, err := os.Stat(*configFile); os.IsNotExist(err) {
		if *configFile == "sqlty.yaml" {
			// Default config file doesn't exist - show usage
			fmt.Fprintf(os.Stderr, "No configuration file found.\n\n")
			flag.Usage()
			os.Exit(0)
		}
		// User specified a config file that doesn't exist
		log.Fatalf("Configuration file not found: %s", *configFile)
	}

	cfg, err := config.LoadFrom(*configFile)
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	if err := compiledir(cfg, *verbose, *timeout); err != nil {
		log.Fatalf("Error: %v", err)
	}

	if *verbose {
		fmt.Println("Done!")
	}
}
