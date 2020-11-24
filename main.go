package main

import (
	"context"
	"log"
	"path"
	"path/filepath"
	"strings"

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

	for _, fname := range sqlFiles {
		q, err := compiler.CompileFile(fname)
		if err != nil {
			return err
		}

		params, returnvals, err := resolver.ResolveTypes(ctx, q.PreparedQuery(), q.NotNullArray())
		if err != nil {
			return err
		}

		stmtq, err := q.StmtQuery("sql", params, returnvals)
		if err != nil {
			return err
		}

		fnameOut := strings.TrimRight(fname, path.Ext(fname)) + ".gen.go"
		if err := gen.Query(fnameOut, stmtq); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	connString := "postgres://localhost/sqlty_test"
	dirPath := "/home/ioj/projects/sqlty-gen/sql"
	if err := compiledir(connString, dirPath); err != nil {
		log.Fatal(err)
	}
}
