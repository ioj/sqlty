package stmt

// Enum represents a custom enum type. Name is a valid Go identifier.
type Enum struct {
	Name   string
	Values []string
}

type Enums struct {
	PackageName string
	Enums       []*Enum
}
