package main

import (
	"fmt"
	"log"

	"github.com/ioj/sqlty/compiler"
)

func main() {
	/*
		ctx := context.Background()
		conn, err := pgconn.Connect(ctx, "postgres://localhost/test1")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close(context.Background())

		// res, err := conn.Prepare(ctx, "", "SELECT * FROM ng_private.workspace where id = $1 LIMIT 5", []uint32{})
		res, err := conn.Prepare(ctx, "", `insert into bleh (col1, col2) values ($1, $2) returning col1`, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", res)
	*/

	// queries, err := compiler.CompileFile("compiler/_queries/select_1.sql")
	queries, err := compiler.CompileString("q", `
		/*
			@name whatever
			@notNullParams  id, dupa.PAKA, bleh
			@param dupa ((oko, paka, blada, sraka )...)
			@param bleh (...)
			@paramStructName oczko
			@returnValueName dupsko
			@many
		*/
		-- ble ble bla
		--
		-- bla bla bla
		-- bo bo
		SELECT * FROM users where id = :id AND :dupa AND :bleh OR whatever LIKE '%pipsko%' AND :id = 5;
	`)
	if err != nil {
		if perr, ok := err.(*compiler.ErrCompilationFailed); ok {
			for _, err := range perr.Errors {
				fmt.Println(err.Sprintf())
			}
			return
		} else {
			log.Fatal(err)
		}
	}

	for _, q := range queries {
		fmt.Println(q)
		fmt.Println(q.Statement())
	}
}
