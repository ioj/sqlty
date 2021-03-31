package stmt

type Param struct {
	Name string
	Type Type
}

type Params struct {
	// If name is not empty, all params are going to be enclosed in a struct instead
	// of passed directly to the query function.
	Name string

	Scalar       []Param
	Spread       []Param
	StructSpread []Struct
}

func (p *Params) None() bool {
	return len(p.Scalar)+len(p.Spread)+len(p.StructSpread) == 0
}
