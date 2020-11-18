package generator

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"text/template"

	"github.com/ioj/sqlty/stmt"
	"github.com/serenize/snaker"
)

// Matches all characters that can't be used in golang's identifiers
// https://golang.org/ref/spec#Identifiers
var identFix = regexp.MustCompile(`[^\pL\pN_]`)

type Generator struct {
	tmpl *template.Template
}

var tmplfn = template.FuncMap{
	"firstParamTypeName": func(params []stmt.Param) string {
		if len(params) == 0 {
			return ""
		}

		return params[0].Type.Name
	},

	"firstParamNilReturnValue": func(params []stmt.Param) string {
		if len(params) == 0 {
			return ""
		}

		if params[0].Type.Nullable {
			return "nil"
		}

		return params[0].Type.ZeroValue
	},

	"firstParamZeroReturnValue": func(params []stmt.Param) string {
		if len(params) == 0 {
			return ""
		}

		return params[0].Type.ZeroValue
	},

	"needsPrintf": func(def *stmt.Query) bool {
		return len(def.Params.Spread)+len(def.Params.StructSpread) > 0
	},

	"hasParams": func(def *stmt.Query) bool {
		return !def.Params.None()
	},

	"valueToIdent": func(val string) string {
		normalized := identFix.ReplaceAllString(val, "_")
		normalized = snaker.SnakeToCamel(normalized)
		return normalized
	},
}

func New(templatepath string) (*Generator, error) {
	var err error
	g := &Generator{}

	g.tmpl, err = template.New("").Funcs(tmplfn).ParseGlob(templatepath)
	if err != nil {
		return nil, err
	}

	// Check if the required template exists
	t := g.tmpl.Lookup("query.go.tpl")
	if t == nil {
		return nil, errors.New("template not found: query.go.tpl")
	}

	return g, nil
}

func (g *Generator) Queries(fname string, q *stmt.Queries) error {
	if q == nil || len(q.Queries) == 0 {
		return errors.New("at least one query is required")
	}

	f, err := os.Create(fname)
	if err != nil {
		return err
	}

	if err := g.tmpl.Lookup("query.go.tpl").Execute(f, q); err != nil {
		return err
	}

	f.Close()

	goimports := exec.Command("goimports", "-w", f.Name())
	goimports.Dir = path.Dir(f.Name())
	if output, err := goimports.CombinedOutput(); err != nil {
		return fmt.Errorf("goimports error: %v, %v", err, string(output))
	}

	return nil
}

func (g *Generator) Enums(pkgpath string, enums *stmt.Enums) error {
	fname := path.Join(pkgpath, "enums.sqlty.gen.go")
	if enums == nil || len(enums.Enums) == 0 {
		// Remove the file if there are no custom enums
		os.Remove(fname)
		return nil
	}

	f, err := os.Create(fname)
	if err != nil {
		return err
	}

	if err := g.tmpl.Lookup("enums.go.tpl").Execute(f, enums); err != nil {
		return err
	}

	f.Close()
	goimports := exec.Command("goimports", "-w", f.Name())
	goimports.Dir = path.Dir(f.Name())
	if output, err := goimports.CombinedOutput(); err != nil {
		return fmt.Errorf("goimports error: %v, %v", err, string(output))
	}

	return nil
}

/*
func main() {
	tmpl, err := template.New("").Funcs(tmplfn).ParseGlob("../templates/*.go.tpl")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("/home/ioj/projects/sqlty-gen/generated.go")
	if err != nil {
		log.Fatal(err)
	}

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
			}
	params := &stmt.Query{
		PackageName: "sql",
		Name:        "AddStocks",
		ExecMode:    "exec",
		Statement:   "insert into stocks (symbol, name, market, currency, enabled) values %[1]v",
		Params: stmt.Params{
			AsStruct: false,
			/*
				Basic: []Param{
					{"ID", Type{"pgtype.UUID", "pgtype.UUID{}", false}},
					{"Limit", Type{"int", "0", false}},
				},
			StructSpread: stmt.StructSpread{
				Name: "Stock",
				Params: []stmt.Param{
					{"Symbol", stmt.Type{"string", "\"\"", false}},
					{"Name", stmt.Type{"string", "\"\"", false}},
					{"Market", stmt.Type{"string", "\"\"", false}},
					{"Currency", stmt.Type{"string", "\"\"", false}},
					{"Enabled", stmt.Type{"bool", "false", false}},
				},
			},
		},
		/*
			Returns: []Param{
				{"ID", Type{"pgtype.UUID", "pgtype.UUID{}", false}},
				{"Name", Type{"string", "\"\"", false}},
				{"Active", Type{"bool", "\"\"", false}},
			},*_/
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

		if output, err := exec.Command("gofmt", "-w", f.Name()).CombinedOutput(); err != nil {
			fmt.Println("-- gofmt error --")
			fmt.Println(string(output))
			fmt.Println(err)
			fmt.Printf("-- gofmt error end --\n\n")
		}

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

		if err := os.Remove(f.Name()); err != nil {
			log.Fatal(err)
		}
}
*/
