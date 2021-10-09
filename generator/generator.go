package generator

import (
	"embed"
	"errors"
	"fmt"
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

//go:embed templates/*
var defaultTemplates embed.FS

type Generator struct {
	tmpl *template.Template

	cache    *cache
	cachedir string
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

	"returnsStruct": func(val *stmt.Query) bool {
		return (val.Returns.IsCompositeType || len(val.Returns.Params) > 1) && val.ExecMode != stmt.ExecModeExec
	},

	"returnsInlineStruct": func(val *stmt.Query) bool {
		return !val.Returns.IsCompositeType && len(val.Returns.Params) > 1 && val.ExecMode != stmt.ExecModeExec
	},
}

func New(templatedir string, cachedir string) (*Generator, error) {
	var err error
	g := &Generator{}

	g.tmpl, err = template.New("").Funcs(tmplfn).ParseFS(defaultTemplates, "templates/*.go.tpl")

	if templatedir != "" {
		glob := path.Join(templatedir, "*.go.tpl")
		g.tmpl, err = g.tmpl.ParseGlob(glob)
	}

	if err != nil {
		return nil, err
	}

	// Check if the required template exists
	t := g.tmpl.Lookup("query.go.tpl")
	if t == nil {
		return nil, errors.New("template not found: query.go.tpl")
	}

	g.cachedir = cachedir
	g.cache, err = newCacheFromFile(cachedir)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Generator) generate(fname string, template string, params interface{}) error {
	updated, err := g.cache.update(fname, params)
	if err != nil {
		return err
	}

	if !updated {
		return nil
	}

	f, err := os.Create(fname)
	if err != nil {
		return err
	}

	if err := g.tmpl.Lookup(template).Execute(f, params); err != nil {
		return err
	}

	return f.Close()
}

func (g *Generator) Query(fname string, q *stmt.Query) error {
	if q == nil {
		return errors.New("query is required")
	}

	return g.generate(fname, "query.go.tpl", q)
}

func (g *Generator) Enums(pkgpath string, enums *stmt.Enums) error {
	fname := path.Join(pkgpath, "enums.sqlty.gen.go")
	if enums == nil || len(enums.Enums) == 0 {
		// Remove the file if there are no custom enums
		os.Remove(fname)
		return nil
	}

	return g.generate(fname, "enums.go.tpl", enums)
}

func (g *Generator) CompositeTypes(pkgpath string, types *stmt.CompositeTypes) error {
	fname := path.Join(pkgpath, "composite_types.sqlty.gen.go")
	if types == nil || len(types.Types) == 0 {
		// Remove the file if there are no custom enums
		os.Remove(fname)
		return nil
	}

	return g.generate(fname, "enums.go.tpl", types)
}

func (g *Generator) DB(pkgpath string, db *stmt.DB) error {
	if db == nil {
		return errors.New("db is required")
	}

	fname := path.Join(pkgpath, "db.sqlty.gen.go")
	return g.generate(fname, "db.sqlty.go.tpl", db)
}

func (g *Generator) Close() error {
	err := g.cache.save(g.cachedir)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
