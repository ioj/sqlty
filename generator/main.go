package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"text/template"
)

func main() {
	tmplfn := template.FuncMap{
		"firstParamTypeName": func(params []Param) string {
			if len(params) == 0 {
				return ""
			}

			return params[0].Type.Name
		},

		"firstParamNilReturnValue": func(params []Param) string {
			if len(params) == 0 {
				return ""
			}

			if params[0].Type.Nullable {
				return "nil"
			}

			return params[0].Type.ZeroValue
		},

		"firstParamZeroReturnValue": func(params []Param) string {
			if len(params) == 0 {
				return ""
			}

			return params[0].Type.ZeroValue
		},

		"needsPrintf": func(def *StatementDef) bool {
			return len(def.Params.Spread) > 0 || def.Params.StructPick.Name != ""
		},

		"hasParams": func(def *StatementDef) bool {
			return !def.Params.None()
		},
	}

	tmpl, err := template.New("").Funcs(tmplfn).ParseGlob("../templates/*.go.tpl")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("/home/ioj/projects/sqlty-gen/generated.go")
	if err != nil {
		log.Fatal(err)
	}

	/*
			params := &StatementDef{
				PackageName: "sql",
				Name:        "WhateverBleh",
				ExecMode:    "many",
				Statement: `SELECT *
		FROM temp;`,
				Comments: []string{
					"SampleFunction is a sample function.",
					"It's nice and does some things.",
				},
				Params: Params{
					AsStruct: true,
					Basic: []Param{
						{"ID", Type{"pgtype.UUID", "&pgtype.UUID{}", true}},
						{"Limit", Type{"int", "0", false}},
					},
					Spread: []Param{
						{"Cats", Type{"string", "\"\"", false}},
						{"Dogs", Type{"string", "\"\"", false}},
					},
					StructPick: StructPick{
						Name: "BlehBloh",
						Params: []Param{
							{"ID", Type{"pgtype.UUID", "&pgtype.UUID{}", true}},
							{"Name", Type{"string", "\"\"", false}},
							{"Cat", Type{"string", "\"\"", false}},
							{"Dog", Type{"string", "\"\"", false}},
						},
					},
				},
				Returns: []Param{
					{"ID", Type{"pgtype.UUID", "&pgtype.UUID{}", true}},
					{"Name", Type{"string", "\"\"", false}},
				},
			}*/
	params := &StatementDef{
		PackageName: "sql",
		Name:        "AddStocks",
		ExecMode:    "exec",
		Statement:   "insert into stocks (symbol, name, market, currency, enabled) values %[1]v",
		Params: Params{
			AsStruct: false,
			/*
				Basic: []Param{
					{"ID", Type{"pgtype.UUID", "pgtype.UUID{}", false}},
					{"Limit", Type{"int", "0", false}},
				},*/
			StructPick: StructPick{
				Name: "Stock",
				Params: []Param{
					{"Symbol", Type{"string", "\"\"", false}},
					{"Name", Type{"string", "\"\"", false}},
					{"Market", Type{"string", "\"\"", false}},
					{"Currency", Type{"string", "\"\"", false}},
					{"Enabled", Type{"bool", "false", false}},
				},
			},
		},
		/*
			Returns: []Param{
				{"ID", Type{"pgtype.UUID", "pgtype.UUID{}", false}},
				{"Name", Type{"string", "\"\"", false}},
				{"Active", Type{"bool", "\"\"", false}},
			},*/
	}

	if err := params.Validate(); err != nil {
		log.Fatal(err)
	}

	if err := tmpl.Lookup("query.go.tpl").Funcs(tmplfn).Execute(f, params); err != nil {
		log.Fatal(err)
	}

	f.Close()

	goimports := exec.Command("goimports", "-w", f.Name())
	goimports.Dir = path.Dir(f.Name())
	if output, err := goimports.CombinedOutput(); err != nil {
		fmt.Println("-- goimports error --")
		fmt.Println(string(output))
		fmt.Println(err)
		fmt.Printf("-- goimports error end --\n\n")
	}

	/*
		if output, err := exec.Command("gofmt", "-w", f.Name()).CombinedOutput(); err != nil {
			fmt.Println("-- gofmt error --")
			fmt.Println(string(output))
			fmt.Println(err)
			fmt.Printf("-- gofmt error end --\n\n")
		}*/

	f, err = os.Open(f.Name())
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	fmt.Println(string(data))
	fmt.Printf("\n\n(%v)\n", f.Name())

	/*
		if err := os.Remove(f.Name()); err != nil {
			log.Fatal(err)
		}*/
}
