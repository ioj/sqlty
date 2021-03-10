package stmt

import "errors"

type ExecMode string

const (
	ExecModeExec ExecMode = "exec"
	ExecModeOne  ExecMode = "one"
	ExecModeMany ExecMode = "many"
)

type Type struct {
	Name      string `yaml:"name"`
	ZeroValue string `yaml:"zeroValue"`
	Nullable  bool   `yaml:"nullable"`
}

type Param struct {
	Name string
	Type Type
}

type Struct struct {
	Name   string
	Params []Param

	// Struct may be rendered in other files if they're used in more than one place.
	ShouldRender bool
}

type Params struct {
	// If name is not empty, all params are going to be enclosed in a struct instead
	// of passed directly to the query function.
	Name string

	Scalar       []Param
	Spread       []Param
	StructSpread []Struct
}

type Query struct {
	PackageName string

	Name      string
	Statement string
	ExecMode  ExecMode
	Comments  []string
	Params    Params
	Returns   Struct
}

func (p *Params) None() bool {
	return len(p.Scalar)+len(p.Spread)+len(p.StructSpread) == 0
}

func (sd *Query) Validate() error {
	if sd.Name == "" {
		return errors.New("missing function name")
	}

	if sd.ExecMode == "exec" && len(sd.Returns.Params) > 0 {
		return errors.New("statement in exec mode can't have return variables")
	}

	if (sd.ExecMode == "one" || sd.ExecMode == "many") && len(sd.Returns.Params) == 0 {
		return errors.New("statement in one/many mode must have return variables")
	}

	return nil
}
