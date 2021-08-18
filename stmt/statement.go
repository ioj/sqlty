package stmt

import (
	"errors"
)

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

type Query struct {
	PackageName string

	Name      string
	Statement string
	ExecMode  ExecMode
	Comments  []string
	Params    Params
	Returns   Struct
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
