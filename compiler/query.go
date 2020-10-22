package compiler

import (
	"fmt"
	"sort"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// token is a string with a location.
type token struct {
	Value string

	Start  int
	Stop   int
	Line   int
	Column int
}

type replacement struct {
	t *token
	r string
}

type StructKey struct {
	token *token
	idx   int

	NotNull bool
}

type ParamType string

var (
	Scalar       ParamType = "scalar"
	Spread       ParamType = "spread"
	StructSpread ParamType = "structspread"
)

type Param struct {
	definition *token
	uses       []*token
	keys       map[string]*StructKey

	Type    ParamType
	Idx     int
	NotNull bool
}

type Query struct {
	scalarIdx int
	spreadIdx int

	name            *token
	execMode        *token
	params          map[string]*Param
	paramStructName *token
	returnValueName *token
	notNullParams   []*token
	percents        []*token

	Filename  string
	Comments  []string
	statement *token
}

// newToken creates a new string with a location marker from the parser context.
func newToken(ctx *antlr.BaseParserRuleContext) *token {
	return &token{
		Value:  ctx.GetText(),
		Start:  ctx.GetStart().GetStart(),
		Stop:   ctx.GetStop().GetStop(),
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (t *token) String() string {
	if t == nil {
		return "nil"
	}

	return fmt.Sprintf("'%v' @ (%v, %v | %v, %v)", t.Value, t.Line, t.Column, t.Start, t.Stop)
}

func (sk *StructKey) Name() string {
	return sk.token.Value
}

func (p *Param) DebugString() string {
	lines := []string{
		fmt.Sprintf("  Definition         : %v", p.definition.String()),
		fmt.Sprintf("  Type, idx, not null: %v, %v, %v", p.Type, p.Idx, p.NotNull),
	}

	var keys []string
	for _, k := range p.keys {
		keys = append(keys, fmt.Sprintf("(%v, %v, %v)", k.token.Value, k.idx, k.NotNull))
	}

	if len(keys) > 0 {
		lines = append(lines,
			"  Keys (key / idx / not null):",
			"    "+strings.Join(keys, ", "))
	}

	var uses []string
	for _, u := range p.uses {
		uses = append(uses, u.String())
	}

	lines = append(lines,
		"  Uses:",
		"    "+strings.Join(uses, ", "))

	return strings.Join(lines, "\n")
}

func (q *Query) DebugString() string {
	lines := []string{
		fmt.Sprintf("Filename         : %v", q.Filename),
		fmt.Sprintf("Name             : %v", q.name.String()),
		fmt.Sprintf("Exec mode        : %v", q.execMode.String()),
		fmt.Sprintf("Param struct name: %v", q.paramStructName.String()),
		fmt.Sprintf("Return value name: %v", q.returnValueName.String()),
		"Params:",
	}

	for _, p := range q.params {
		lines = append(lines, p.DebugString(), "  ---")
	}

	lines = append(lines, "Statement:", q.statement.String())

	return strings.Join(lines, "\n")
}

// NeedsSprintf returns true when the query contains spreads or
// struct spreads which need to be Sprintf'ed when executing the
// query.
func (q *Query) NeedsSprintf() bool {
	for _, p := range q.params {
		if p.Type != "Scalar" {
			return true
		}
	}

	return false
}

func (q *Query) Statement() string {
	if len(q.params) == 0 {
		return q.statement.Value
	}

	var reps []replacement

	if q.NeedsSprintf() {
		for _, p := range q.percents {
			reps = append(reps, replacement{t: p, r: "%%"})
		}
	}

	var structSpreadIdx int
	for _, p := range q.params {
		if p.Type == "Spread" {
			structSpreadIdx++
		}
	}

	for _, p := range q.params {
		var repstr string

		switch p.Type {
		case "Scalar":
			repstr = fmt.Sprintf("$%d", p.Idx+1)
		case "Spread":
			repstr = fmt.Sprintf("%%[%d]v", p.Idx)
		case "StructSpread":
			repstr = fmt.Sprintf("%%[%d]v", structSpreadIdx)
		}

		for _, u := range p.uses {
			reps = append(reps, replacement{t: u, r: repstr})
		}
	}

	sort.Slice(reps, func(i, j int) bool {
		return reps[i].t.Start < reps[j].t.Start
	})

	var stmt strings.Builder

	// Add the fragment before the first replacement
	origstart := q.statement.Start + 1
	orig := q.statement.Value

	offset := reps[0].t.Start - origstart
	stmt.WriteString(orig[0:offset])

	// insert all replacements
	for n, r := range reps {
		stmt.WriteString(r.r)

		offset = r.t.Stop - origstart + 2

		var next int
		if n == len(reps)-1 {
			// if this is the last replacement, add the rest
			// of the original statement
			next = len(orig)
		} else {
			// otherwise, write the statement up to the next
			// replacement and do next iteration
			next = reps[n+1].t.Start - origstart
		}

		stmt.WriteString(orig[offset:next])
	}

	return stmt.String()
}
