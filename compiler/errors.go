package compiler

import (
	"errors"
	"fmt"
	"strings"
)

var ErrEmptyFile = errors.New("empty file")

// ParserError is a syntax error reported by the parser or lexer,
// or an annotation error.
type ParserError struct {
	fname   string
	errtype string
	msg     string
	line    int
	column  int
}

// ErrCompilationFailed is returned when the compiler has encountered
// one of more compilation errors.
type ErrCompilationFailed struct {
	Errors []*ParserError
}

type errorListener struct {
	fname  string
	errors []*ParserError
}

// newErrorListener creates a new error listener.
func newErrorListener(fname string) *errorListener {
	return &errorListener{fname: fname}
}

func (e *ParserError) Error() string {
	return fmt.Sprintf("line %v:%v: [%v] %v", e.line, e.column, e.errtype, e.msg)
}

// Sprintf returns a formatted error message.
func (e *ParserError) Sprintf() string {
	return fmt.Sprintf("[error][%v] %v:%v:%v: %v", e.errtype, e.fname, e.line, e.column, e.msg)
}

func (e *ErrCompilationFailed) Error() string {
	msg := []string{"compilation error(s)"}
	if len(e.Errors) > 0 {
		msg = append(msg, ": ", e.Errors[0].Error())
	}
	if len(e.Errors) > 1 {
		msg = append(msg, fmt.Sprintf(", and %v more", len(e.Errors)-1))
	}

	return strings.Join(msg, "")
}

// SyntaxError handles syntax errors.
func (el *errorListener) SyntaxError(line, column int, msg string) {
	el.errors = append(el.errors, &ParserError{
		fname:   el.fname,
		errtype: "syntax",
		msg:     msg,
		line:    line,
		column:  column,
	})
}

func (el *errorListener) EmptyFileError() {
	el.errors = append(el.errors, &ParserError{
		fname:   el.fname,
		errtype: "empty",
		msg:     "empty file",
	})
}

func (el *errorListener) AlreadyDeclaredError(what string, existing, current *token) {
	el.errors = append(el.errors, &ParserError{
		fname:   el.fname,
		errtype: "annotation",
		line:    current.Line,
		column:  current.Column,
		msg: fmt.Sprintf("%v is already declared (`%v`) at line %v:%v", what,
			existing.Value, existing.Line, existing.Column),
	})
}

func (el *errorListener) UnusedParamError(p *Param) {
	el.errors = append(el.errors, &ParserError{
		fname:   el.fname,
		errtype: "annotation",
		line:    p.definition.Line,
		column:  p.definition.Column,
		msg:     fmt.Sprintf("parameter `%v` is declared, but not used in the query", p.definition.Value),
	})
}

func (el *errorListener) MissingParamError(t *token) {
	el.errors = append(el.errors, &ParserError{
		fname:   el.fname,
		errtype: "annotation",
		line:    t.Line,
		column:  t.Column,
		msg:     fmt.Sprintf("missing parameter `%v` marked as not null", t.Value),
	})
}

func (el *errorListener) DuplicateStructKeyError(param string, t *token) {
	el.errors = append(el.errors, &ParserError{
		fname:   el.fname,
		errtype: "annotation",
		line:    t.Line,
		column:  t.Column,
		msg:     fmt.Sprintf("duplicate property `%v` for struct parameter `%v`", t.Value, param),
	})
}

func (el *errorListener) MissingExecModeError(stmt *token) {
	q := strings.SplitN(stmt.Value, "\n", 2)
	firstline := q[0]
	if len(q) > 1 {
		firstline += "..."
	}

	el.errors = append(el.errors, &ParserError{
		fname:   el.fname,
		errtype: "annotation",
		line:    stmt.Line,
		column:  stmt.Column,
		msg:     fmt.Sprintf("exec mode is missing [one/many/exec] for query `%v`", firstline),
	})
}

func (el *errorListener) DeprecatedArrowError(t *token) {
	el.errors = append(el.errors, &ParserError{
		fname:   el.fname,
		errtype: "annotation",
		line:    t.Line,
		column:  t.Column,
		msg:     "the '->' arrow is no longer required in @param syntax (use '@param name (...)' instead of '@param name -> (...)')",
	})
}

func (el *errorListener) InconsistentNotNullError(paramName string, firstToken *token, markedCount, totalCount int) {
	el.errors = append(el.errors, &ParserError{
		fname:   el.fname,
		errtype: "annotation",
		line:    firstToken.Line,
		column:  firstToken.Column,
		msg: fmt.Sprintf("parameter `%v` has inconsistent not-null markers: %d of %d uses have '!' (all uses must be consistent)",
			paramName, markedCount, totalCount),
	})
}

func (el *errorListener) Error() error {
	if len(el.errors) == 0 {
		return nil
	}

	return &ErrCompilationFailed{Errors: el.errors}
}
