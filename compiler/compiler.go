package compiler

import (
	"os"
)

// CompileFile returns the compiled query based on contents of a file.
func CompileFile(fname string) (*Query, error) {
	data, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	return CompileString(fname, string(data))
}

// CompileString returns the compiled query based on the input provided as a string.
func CompileString(defaultQueryName string, data string) (*Query, error) {
	p := NewParser(defaultQueryName, data)
	return p.Parse()
}
