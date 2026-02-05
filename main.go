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
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/ioj/sqlty/compiler"
	"github.com/ioj/sqlty/config"
	"github.com/ioj/sqlty/db"
	"github.com/ioj/sqlty/generator"
	"github.com/ioj/sqlty/stmt"
	"github.com/ioj/sqlty/watcher"
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
	watchMode   = flag.Bool("watch", false, "watch for file changes and recompile automatically")
)

// compilationContext holds the resources needed for compilation.
type compilationContext struct {
	cfg      *config.Config
	resolver *db.Resolver
	gen      *generator.Generator
	ctx      context.Context
	verbose  bool
}

// newCompilationContext creates a new compilation context with connected resources.
func newCompilationContext(ctx context.Context, cfg *config.Config, verbose bool) (*compilationContext, error) {
	types, err := loadTypes(cfg)
	if err != nil {
		return nil, err
	}

	if verbose {
		fmt.Println("Connecting to database...")
	}

	resolver, err := db.NewResolver(ctx, cfg.DBURL, types)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	gen, err := generator.New(cfg.Paths.Templates, cfg.Paths.Cache)
	if err != nil {
		resolver.Close()
		return nil, fmt.Errorf("failed to initialize generator: %w", err)
	}

	return &compilationContext{
		cfg:      cfg,
		resolver: resolver,
		gen:      gen,
		ctx:      ctx,
		verbose:  verbose,
	}, nil
}

// Close releases the compilation context resources.
func (cc *compilationContext) Close() {
	cc.gen.Close()
	cc.resolver.Close()
}

// loadTypes loads the type mappings from default or custom file.
func loadTypes(cfg *config.Config) ([]db.PGTypeTranslation, error) {
	defaulttypes := &db.PGTypeTranslationsFile{}
	if cfg.DefaultTypes == "" {
		if err := yaml.NewDecoder(bytes.NewReader(defaultTypes)).Decode(defaulttypes); err != nil {
			return nil, fmt.Errorf("failed to decode default types: %w", err)
		}
	} else {
		f, err := os.Open(cfg.DefaultTypes)
		if err != nil {
			return nil, fmt.Errorf("failed to open custom types file %s: %w", cfg.DefaultTypes, err)
		}
		defer f.Close()

		if err := yaml.NewDecoder(f).Decode(defaulttypes); err != nil {
			return nil, fmt.Errorf("failed to decode custom types file %s: %w", cfg.DefaultTypes, err)
		}
	}

	return append(defaulttypes.Types, cfg.Types...), nil
}

// getOutputFilename converts a source SQL file path to its generated Go file path.
func getOutputFilename(sourcePath, outputDir string) string {
	base := filepath.Base(sourcePath)
	withoutExt := strings.TrimSuffix(base, filepath.Ext(base))
	return filepath.Join(outputDir, withoutExt+".gen.go")
}

// compileSingleFile compiles a single SQL file and generates its output.
func compileSingleFile(cc *compilationContext, fname string) error {
	q, err := compiler.CompileFile(fname)
	if err != nil {
		if err == compiler.ErrEmptyFile {
			fmt.Printf("warn: %v is empty, ignoring\n", fname)
			return nil
		}
		return fmt.Errorf("%s: %w", fname, err)
	}

	params, returns, err := cc.resolver.ResolveTypes(cc.ctx, q.PreparedQuery(), q.NotNullArray())
	if err != nil {
		return fmt.Errorf("%s: failed to resolve types: %w", fname, err)
	}

	stmtq, err := q.StmtQuery(cc.cfg.PackageName, params, returns)
	if err != nil {
		return fmt.Errorf("%s: %w", fname, err)
	}

	if (returns == nil || len(returns.Params) == 0) && stmtq.ExecMode != stmt.ExecModeExec {
		fmt.Printf("warn: %s has no return values, changing exec mode to @exec\n", stmtq.Name)
		stmtq.ExecMode = stmt.ExecModeExec
	}

	fnameOut := getOutputFilename(fname, cc.cfg.Paths.Output)

	if err := cc.gen.Query(q.Template(), fnameOut, stmtq); err != nil {
		return fmt.Errorf("%s: failed to generate query: %w", fname, err)
	}

	return nil
}

// regenerateSharedFiles generates the shared files (enums, composite types, db utilities).
func regenerateSharedFiles(cc *compilationContext) error {
	enums := &stmt.Enums{PackageName: cc.cfg.PackageName, Enums: cc.resolver.Enums()}
	if err := cc.gen.Enums(cc.cfg.Paths.Output, enums); err != nil {
		return fmt.Errorf("failed to generate enums: %w", err)
	}

	if err := cc.gen.DB(cc.cfg.Paths.Output, &stmt.DB{PackageName: cc.cfg.PackageName}); err != nil {
		return fmt.Errorf("failed to generate db utilities: %w", err)
	}

	ct, err := cc.resolver.CompositeTypes(cc.ctx)
	if err != nil {
		return fmt.Errorf("failed to resolve composite types: %w", err)
	}

	compositeTypes := &stmt.CompositeTypes{PackageName: cc.cfg.PackageName, Types: ct}
	if err := cc.gen.CompositeTypes(cc.cfg.Paths.Output, compositeTypes); err != nil {
		return fmt.Errorf("failed to generate composite types: %w", err)
	}

	return nil
}

// runGoimports runs goimports on the output directory.
func runGoimports(outputDir string) {
	if _, err := exec.LookPath("goimports"); err != nil {
		fmt.Println("warn: goimports not found, generated code may need manual formatting")
		return
	}

	goimports := exec.Command("goimports", "-w", ".")
	goimports.Dir = outputDir
	if output, err := goimports.CombinedOutput(); err != nil {
		fmt.Printf("warn: goimports failed: %v\n%s\n", err, string(output))
		fmt.Println("Generated code may need manual formatting")
	}
}

// findSQLFiles returns all SQL files in the source directory.
func findSQLFiles(sourceDir string) ([]string, error) {
	sqlFiles, err := filepath.Glob(filepath.Join(sourceDir, "*.sql"))
	if err != nil {
		return nil, fmt.Errorf("failed to find SQL files: %w", err)
	}
	return sqlFiles, nil
}

func compiledir(cfg *config.Config, verbose bool, timeout time.Duration) error {
	sqlFiles, err := findSQLFiles(cfg.Paths.Source)
	if err != nil {
		return err
	}

	if len(sqlFiles) == 0 {
		return fmt.Errorf("no *.sql files to compile in %v", cfg.Paths.Source)
	}

	if verbose {
		fmt.Printf("Found %d SQL files to process\n", len(sqlFiles))
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cc, err := newCompilationContext(ctx, cfg, verbose)
	if err != nil {
		return err
	}
	defer cc.Close()

	if err := regenerateSharedFiles(cc); err != nil {
		return err
	}

	var compileTime time.Duration

	for i, fname := range sqlFiles {
		if verbose {
			fmt.Printf("[%d/%d] Processing %s\n", i+1, len(sqlFiles), filepath.Base(fname))
		}

		t1 := time.Now()
		if err := compileSingleFile(cc, fname); err != nil {
			return err
		}
		compileTime += time.Since(t1)
	}

	runGoimports(cfg.Paths.Output)

	if verbose {
		fmt.Printf("\nTiming: total=%v\n", compileTime.Round(time.Millisecond))
		fmt.Printf("Generated %d query files\n", len(sqlFiles))
	}

	return nil
}

func watch(cfg *config.Config, verbose bool, timeout time.Duration) error {
	fmt.Println("Starting watch mode...")

	// Initial full compilation
	if err := compiledir(cfg, verbose, timeout); err != nil {
		fmt.Printf("Initial compilation error: %v\n", err)
		fmt.Println("Continuing to watch for changes...")
	} else {
		fmt.Println("Initial compilation complete.")
	}

	// Create watcher with 100ms debounce
	w, err := watcher.New(cfg.Paths.Source, 100*time.Millisecond)
	if err != nil {
		return fmt.Errorf("failed to create watcher: %w", err)
	}
	defer w.Close()

	// Setup context with signal handling
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle interrupt signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("\nShutting down...")
		cancel()
	}()

	// Create persistent compilation context for watch mode
	cc, err := newCompilationContext(ctx, cfg, verbose)
	if err != nil {
		return fmt.Errorf("failed to setup compilation context: %w", err)
	}
	defer cc.Close()

	events, errors := w.Start(ctx)

	fmt.Printf("Watching for changes in %s...\n", cfg.Paths.Source)

	for {
		select {
		case <-ctx.Done():
			return nil

		case event, ok := <-events:
			if !ok {
				return nil
			}

			handleFileChanges(cc, event.Files)

		case err, ok := <-errors:
			if !ok {
				return nil
			}
			fmt.Printf("Watch error: %v\n", err)
		}
	}
}

// handleFileChanges processes a batch of changed files.
func handleFileChanges(cc *compilationContext, files []string) {
	fmt.Printf("Recompiling %d file(s)...", len(files))

	hasError := false

	for _, fname := range files {
		// Check if file still exists (might have been deleted)
		if _, err := os.Stat(fname); os.IsNotExist(err) {
			// File was deleted, remove the generated file
			fnameOut := getOutputFilename(fname, cc.cfg.Paths.Output)
			if err := os.Remove(fnameOut); err != nil && !os.IsNotExist(err) {
				fmt.Printf("Error removing %s: %v\n", fnameOut, err)
			} else if cc.verbose {
				fmt.Printf("Removed %s\n", fnameOut)
			}
			continue
		}

		if cc.verbose {
			fmt.Printf("Processing %s\n", filepath.Base(fname))
		}

		if err := compileSingleFile(cc, fname); err != nil {
			fmt.Printf("Error: %v\n", err)
			hasError = true
		}
	}

	// Always regenerate shared files
	if err := regenerateSharedFiles(cc); err != nil {
		fmt.Printf("Error generating shared files: %v\n", err)
		hasError = true
	}

	runGoimports(cc.cfg.Paths.Output)

	if !hasError {
		fmt.Println("Recompilation complete.")
	} else {
		fmt.Println("Recompilation completed with errors.")
	}
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

	if *watchMode {
		if err := watch(cfg, *verbose, *timeout); err != nil {
			log.Fatalf("Error: %v", err)
		}
	} else {
		if err := compiledir(cfg, *verbose, *timeout); err != nil {
			log.Fatalf("Error: %v", err)
		}

		if *verbose {
			fmt.Println("Done!")
		}
	}
}
