package generator

import (
	"errors"
	"os"
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

func (g *Generator) Query(fname string, q *stmt.Query) error {
	if q == nil {
		return errors.New("query is required")
	}

	f, err := os.Create(fname)
	if err != nil {
		return err
	}

	if err := g.tmpl.Lookup("query.go.tpl").Execute(f, q); err != nil {
		return err
	}

	return f.Close()
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

	return f.Close()
}
