package compiler

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ioj/sqlty/compiler/parser"
)

func compile(fname string, input antlr.CharStream) (*Query, error) {
	pl := &parserListener{
		filename: fname,
		errors:   newErrorListener(fname),
	}

	lexer := parser.NewSQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSQLParser(stream)
	p.RemoveErrorListeners()
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.AddErrorListener(pl.errors)
	p.BuildParseTrees = true

	antlr.ParseTreeWalkerDefault.Walk(pl, p.Input())

	if pl.Query == nil {
		return nil, nil
	}

	pl.CheckUnusedParameters()
	pl.PopulateNotNullParams()
	pl.VerifyExecMode()

	fmt.Println(pl.Query.DebugString())

	if err := pl.errors.Error(); err != nil {
		return nil, err
	}

	return pl.Query, nil
}

// CompileFile returns the list of compiled queries based on contents of a file.
func CompileFile(fname string) (*Query, error) {
	input, err := antlr.NewFileStream(fname)
	if err != nil {
		return nil, err
	}

	return compile(fname, input)
}

// CompileString returns the list of compiled queries based on the input provided
// as a string.
func CompileString(defaultQueryName string, data string) (*Query, error) {
	input := antlr.NewInputStream(data)
	return compile(defaultQueryName, input)
}
