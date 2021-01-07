package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/ioj/sqlty/compiler"
	"github.com/ioj/sqlty/db"
	"github.com/ioj/sqlty/generator"
	"github.com/ioj/sqlty/stmt"
)

func compiledir(connString, dirPath string) error {
	sqlFiles, err := filepath.Glob(path.Join(dirPath, "*.sql"))
	if err != nil {
		return err
	}

	if len(sqlFiles) == 0 {
		return fmt.Errorf("no *.sql files to compile in %v", dirPath)
	}

	ctx := context.Background()
	resolver, err := db.NewResolver(ctx, connString)
	if err != nil {
		return err
	}
	defer resolver.Close()

	gen, err := generator.New("./templates/*.go.tpl")
	if err != nil {
		return err
	}

	enums := &stmt.Enums{PackageName: "sql", Enums: resolver.Enums()}
	if err := gen.Enums(dirPath, enums); err != nil {
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

		params, returnvals, err := resolver.ResolveTypes(ctx, q.PreparedQuery(), q.NotNullArray())
		if err != nil {
			return err
		}

		t3 := time.Now()
		resolveTime += t3.Sub(t2)

		stmtq, err := q.StmtQuery("sql", params, returnvals)
		if err != nil {
			return err
		}

		fnameOut := strings.TrimSuffix(fname, path.Ext(fname)) + ".gen.go"
		if err := gen.Query(fnameOut, stmtq); err != nil {
			return err
		}

		t4 := time.Now()
		generateTime += t4.Sub(t3)
	}

	t := time.Now()
	goimports := exec.Command("goimports", "-w", ".")
	goimports.Dir = path.Dir(dirPath)
	if output, err := goimports.CombinedOutput(); err != nil {
		return fmt.Errorf("goimports error: %v, %v", err, string(output))
	}
	goimportsTime := time.Now().Sub(t)

	fmt.Printf("Processed files: %d\n", len(sqlFiles))
	fmt.Printf("Compile time: %.02fs\n", compileTime.Seconds())
	fmt.Printf("Resolve time: %.02fs\n", resolveTime.Seconds())
	fmt.Printf("Generate time: %.02fs\n", generateTime.Seconds())
	fmt.Printf("Goimports time: %.02fs\n", goimportsTime.Seconds())

	return nil
}

func main() {
	connString := "postgres://localhost/stonky"
	dirPath := "/home/ioj/projects/stonks/database/sql"
	if err := compiledir(connString, dirPath); err != nil {
		log.Fatal(err)
	}
}
