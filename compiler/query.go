package compiler

import (
	"fmt"
	"path"
	"regexp"
	"sort"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ioj/sqlty/helpers"
	"github.com/ioj/sqlty/stmt"
	"github.com/serenize/snaker"
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
	paramIdx map[ParamType]int

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

func (q *Param) Keys() []*StructKey {
	keys := make([]*StructKey, len(q.keys))
	for _, k := range q.keys {
		keys[k.idx] = k
	}

	return keys
}

func (q *Query) Name() string {
	return q.name.Value
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

// Params returns the query parameters of a given type, sorted
// by their index.
func (q *Query) Params(ptype ParamType) []*Param {
	var params []*Param

	for _, p := range q.params {
		if p.Type == ptype {
			params = append(params, p)
		}
	}

	sort.Slice(params, func(i, j int) bool {
		return params[i].Idx < params[j].Idx
	})

	return params
}

// NotNullArray returns a list of not-null properties for sorted
// params. Params are sorted so that scalars go first, then spreads,
// then a struct spreads.
func (q *Query) NotNullArray() []bool {
	var retval []bool

	for _, scalar := range q.Params(Scalar) {
		retval = append(retval, scalar.NotNull)
	}

	for _, spread := range q.Params(Spread) {
		retval = append(retval, spread.NotNull)
	}

	for _, ss := range q.Params(StructSpread) {
		for _, k := range ss.Keys() {
			retval = append(retval, k.NotNull)
		}
	}

	return retval
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

	structSpreadIdx := len(q.Params(Spread)) + 1

	for _, p := range q.params {
		var repstr string

		switch p.Type {
		case Scalar:
			repstr = fmt.Sprintf("$%d", p.Idx+1)
		case Spread:
			repstr = fmt.Sprintf("%%[%d]v", p.Idx+1)
		case StructSpread:
			repstr = fmt.Sprintf("%%[%d]v", p.Idx+structSpreadIdx)
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

func (q *Query) PreparedQuery() string {
	if !q.NeedsSprintf() {
		return q.Statement()
	}

	stmt := q.Statement()
	params := []interface{}{}
	idx := len(q.Params(Scalar)) + 1

	for range q.Params(Spread) {
		params = append(params, fmt.Sprintf("$%[1]d,$%[1]d", idx))
		idx++
	}

	for _, ss := range q.Params(StructSpread) {
		keyslen := len(ss.keys)
		keys := make([]string, keyslen)
		for i := 0; i < keyslen; i++ {
			keys[i] = fmt.Sprintf("$%d", idx)
			idx++
		}

		params = append(params, fmt.Sprintf("(%[1]v),(%[1]v)", strings.Join(keys, ",")))
	}

	return fmt.Sprintf(stmt, params...)
}

// StmtQuery returns a stmt.Query based on internal values and type resolutions for arguments
// and return values provided in parameters.
func (q *Query) StmtQuery(packageName string, ptypes []stmt.Type, returns *stmt.Struct) (*stmt.Query, error) {
	stmtq := &stmt.Query{
		PackageName: packageName,
		Statement:   q.Statement(),
		Comments:    q.Comments,
	}

	// Set the name. If it's empty, then use the name derived from the filename.
	// If there are multiple queries in the filename, they'll be fixed later on during
	// per-filename checks.
	var name string

	if q.name != nil {
		name = q.name.Value
	}

	if name == "" {
		// Get filename without extension
		name = strings.TrimSuffix(path.Base(q.Filename), path.Ext(q.Filename))
		name = regexp.MustCompile(`[^a-zA-Z0-9_.]`).ReplaceAllString(name, "_")
	}

	stmtq.Name = snaker.SnakeToCamel(name)

	// Set the exec mode
	switch q.execMode.Value {
	case "@one":
		stmtq.ExecMode = stmt.ExecModeOne
	case "@many":
		stmtq.ExecMode = stmt.ExecModeMany
	case "@exec":
		stmtq.ExecMode = stmt.ExecModeExec
	default:
		return nil, fmt.Errorf("unknown exec mode: %v", q.execMode.Value)
	}

	// Handle params. Param types passed as ptypes are expected to be sorted the same way
	// as the q.PreparedStatement does, i.e. scalars, spreads, struct spreads.
	// This way, we can apply the types to the parameters in the correct order.
	//
	// Also, make sure that there are no duplicate param names. This was checked by the
	// compiler before, but new duplicates may appear after snake->camel normalization.
	private := true

	if q.paramStructName != nil {
		stmtq.Params.Name = snaker.SnakeToCamel(q.paramStructName.Value)

		// If there's a param struct name, all parameters are going to be placed in a struct,
		// so their names should be public. If the param struct name is empty, all parameters
		// are generated as function parameters.
		private = false
	}

	ptypeidx := 0
	pnames := helpers.NewStructFieldNormalizer()
	for _, p := range q.Params(Scalar) {
		normalized, err := pnames.Add(p.definition.Value, private)
		if err != nil {
			return nil, err
		}

		stmtq.Params.Scalar = append(stmtq.Params.Scalar, stmt.Param{
			Name: normalized,
			Type: ptypes[ptypeidx],
		})
		ptypeidx++
	}

	for _, p := range q.Params(Spread) {
		normalized, err := pnames.Add(p.definition.Value, private)
		if err != nil {
			return nil, err
		}

		stmtq.Params.Spread = append(stmtq.Params.Spread, stmt.Param{
			Name: normalized,
			Type: ptypes[ptypeidx],
		})
		ptypeidx++
	}

	for _, p := range q.Params(StructSpread) {
		normalized, err := pnames.Add(p.definition.Value, false)
		if err != nil {
			return nil, err
		}

		s := stmt.Struct{Name: normalized, DontRender: true}
		structkeynames := helpers.NewStructFieldNormalizer()
		for _, sp := range p.Keys() {
			normalizedkey, err := structkeynames.Add(sp.Name(), false)
			if err != nil {
				return nil, err
			}
			s.Params = append(s.Params, stmt.Param{
				Name: normalizedkey,
				Type: ptypes[ptypeidx],
			})
			ptypeidx++
		}

		stmtq.Params.StructSpread = append(stmtq.Params.StructSpread, s)
	}

	// Handle return value: normalize return value struct name, struct fields,
	// make sure that there are no duplicate struct field names.
	if returns != nil {
		stmtq.Returns = *returns
	}

	if q.returnValueName != nil {
		// Override return value name
		stmtq.Returns.Name = q.returnValueName.Value
	}

	if stmtq.Returns.Name == "" {
		stmtq.Returns.Name = stmtq.Name + "Row"
	}

	retparams := helpers.NewStructFieldNormalizer()
	for n, param := range stmtq.Returns.Params {
		normalized, err := retparams.Add(param.Name, false)
		if err != nil {
			return nil, err
		}

		param.Name = normalized
		stmtq.Returns.Params[n] = param
	}

	if len(stmtq.Returns.Params) < 2 {
		stmtq.Returns.DontRender = true
	}

	return stmtq, nil
}
