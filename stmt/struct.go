package stmt

import "github.com/ioj/sqlty/helpers"

type Struct struct {
	Name   string
	Params []Param

	// Struct may be rendered in other files if they're used in more than one place
	// or are composite type.
	IsCompositeType bool
}

func (s *Struct) GolangizeParamNames() error {
	sn := helpers.NewStructFieldNormalizer()

	for n, param := range s.Params {
		normalized, err := sn.Add(param.Name, false)
		if err != nil {
			return err
		}

		param.Name = normalized
		s.Params[n] = param
	}

	return nil
}
