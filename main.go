package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ioj/sqlty/compiler"
	"github.com/ioj/sqlty/db"
	"github.com/ioj/sqlty/generator"
	"github.com/ioj/sqlty/stmt"
)

func main() {
	// queries, err := compiler.CompileFile("compiler/_queries/select_1.sql")
	queries, err := compiler.CompileString("q", `
		/*
			@name foobar_by_id
			@notNullParams id
			@many
		*/
		SELECT * FROM test2 WHERE id = :id;

		/*
			@name foobar_by_col
			@param timestamps ((ts1, ts2)...)
			@exec
		*/
		INSERT INTO test2 (ts1, ts2) VALUES :timestamps;
	`)
	if err != nil {
		if perr, ok := err.(*compiler.ErrCompilationFailed); ok {
			for _, err := range perr.Errors {
				fmt.Println(err.Sprintf())
			}
			return
		}

		log.Fatal(err)
	}

	resolver, err := db.NewResolver(context.Background(), "postgres://localhost/sqlty_test")
	if err != nil {
		log.Fatal(err)
	}
	defer resolver.Close()

	stmtqueries := &stmt.Queries{PackageName: "sql"}

	for _, q := range queries {
		fmt.Println(q.PreparedQuery())
		params, returnvals, err := resolver.ResolveTypes(context.Background(),
			q.PreparedQuery(), q.NotNullArray())
		if err != nil {
			log.Fatal(err)
		}

		for n, p := range params {
			fmt.Printf("$%v: %+v\n", n+1, p)
		}

		fmt.Println("---")

		for _, rv := range returnvals {
			fmt.Printf("%v: %+v\n", rv.Name, rv.Type)
		}

		stmtq, err := q.StmtQuery(params, returnvals)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", *stmtq)

		stmtqueries.Queries = append(stmtqueries.Queries, stmtq)
	}

	gen, err := generator.New("./templates/*.go.tpl")
	if err != nil {
		log.Fatal(err)
	}

	enums := &stmt.Enums{PackageName: "sql", Enums: resolver.Enums()}
	if err := gen.Enums("/home/ioj/projects/sqlty-gen/sql", enums); err != nil {
		log.Fatal(err)
	}

	if err := gen.Queries("/home/ioj/projects/sqlty-gen/sql/generated.go", stmtqueries); err != nil {
		log.Fatal(err)
	}
}
