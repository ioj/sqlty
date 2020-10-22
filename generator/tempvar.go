package main

import "errors"

type ExecMode string

const (
	ExecModeExec = "exec"
	ExecModeOne  = "one"
	ExecModeMany = "many"
)

type Type struct {
	Name      string
	ZeroValue string
	Nullable  bool
}

type Param struct {
	Name string
	Type Type
}

type StructPick struct {
	Name   string
	Params []Param
}

type Params struct {
	// If true, params are going to be packed into the struct of name Name+"Params"
	AsStruct   bool
	Basic      []Param
	Spread     []Param
	StructPick StructPick
}

type StatementDef struct {
	PackageName    string
	Name           string
	Statement      string
	ExecMode       ExecMode
	ReturnTypeName string
	Comments       []string
	Params         Params
	Returns        []Param
}

func (p *Params) None() bool {
	return len(p.Basic) == 0 && len(p.Spread) == 0 && (p.StructPick.Name == "")
}

func (sd *StatementDef) Validate() error {
	if sd.PackageName == "" {
		return errors.New("missing package name")
	}

	if sd.Name == "" {
		return errors.New("missing function name")
	}

	if sd.ExecMode == "exec" && len(sd.Returns) > 0 {
		return errors.New("statement in exec mode can't have return variables")
	}

	if (sd.ExecMode == "one" || sd.ExecMode == "many") && len(sd.Returns) == 0 {
		return errors.New("statement in one/many mode must have return variables")
	}

	return nil
}
