package compiler

import (
	"strings"

	"github.com/ioj/sqlty/compiler/parser"
)

type parserListener struct {
	*parser.BaseSQLParserListener

	filename string
	errors   *errorListener

	Query *Query

	currentParam           *Param
	currentStructTransform []*StructKey
	currentNotNull         []*token
}

// EnterQuery is called when a new query (annotations + the query itself) is encountered.
func (s *parserListener) EnterQuery(ctx *parser.QueryContext) {
	s.Query = &Query{
		Filename: s.filename,
		paramIdx: make(map[ParamType]int),
		params:   make(map[string]*Param),
	}
}

// EnterParamTag is called the parser encounters the @param tag.
func (s *parserListener) EnterParamTag(ctx *parser.ParamTagContext) {
	t := newToken(ctx.BaseParserRuleContext)
	s.currentParam = &Param{definition: t}
}

// ExitParamTag is called when the parser leaves the '@param name -> ...' statement.
func (s *parserListener) ExitParamTag(ctx *parser.ParamTagContext) {
	// The current param is nil when it was already defined before.
	if s.currentParam != nil {
		lname := strings.ToLower(s.currentParam.definition.Value)
		s.Query.params[lname] = s.currentParam
		s.currentParam = nil
	}
}

// EnterParamName is called when the parser processes @param name. It also checks
// whether the name was already declared.
func (s *parserListener) EnterParamName(ctx *parser.ParamNameContext) {
	t := newToken(ctx.BaseParserRuleContext)
	s.currentParam.definition.Value = t.Value
	lname := strings.ToLower(s.currentParam.definition.Value)
	if alreadyDefined, ok := s.Query.params[lname]; ok {
		s.currentParam = nil
		s.errors.AlreadyDeclaredError("parameter", alreadyDefined.definition, t)
	}
}

// ExitSpreadTransform is called when production spreadTransform is exited.
func (s *parserListener) ExitSpreadTransform(ctx *parser.SpreadTransformContext) {
	if s.currentParam == nil {
		return
	}

	s.currentParam.Type = Spread
	s.currentParam.Idx = s.Query.paramIdx[Spread]
	s.Query.paramIdx[Spread]++
}

// ExitStructTransform is called when the parser has finished parsing the struct field
// list. It happens in @param defined as a struct spread.
func (s *parserListener) ExitStructTransform(ctx *parser.StructTransformContext) {
	if s.currentParam == nil {
		return
	}

	s.currentParam.keys = make(map[string]*StructKey)

	for _, sk := range s.currentStructTransform {
		lname := strings.ToLower(sk.token.Value)
		_, ok := s.currentParam.keys[lname]
		if ok {
			s.errors.DuplicateStructKeyError(s.currentParam.definition.Value, sk.token)
			continue
		}

		s.currentParam.keys[lname] = sk
	}

	s.currentStructTransform = nil
}

// ExitStructSpreadTransform is called when the parser finishes parsing of @param defined
// as a struct spread.
func (s *parserListener) ExitStructSpreadTransform(ctx *parser.StructSpreadTransformContext) {
	if s.currentParam == nil {
		return
	}

	s.currentParam.Type = StructSpread
	s.currentParam.Idx = s.Query.paramIdx[StructSpread]
	s.Query.paramIdx[StructSpread]++
}

// EnterKey is called when production key is entered.
func (s *parserListener) EnterKey(ctx *parser.KeyContext) {
	structkey := &StructKey{
		token: newToken(ctx.BaseParserRuleContext),
		idx:   len(s.currentStructTransform),
	}
	s.currentStructTransform = append(s.currentStructTransform, structkey)
}

// EnterParamStructNameId is called when production paramStructNameId is entered.
func (s *parserListener) EnterParamStructNameId(ctx *parser.ParamStructNameIdContext) {
	t := newToken(ctx.BaseParserRuleContext)
	if s.Query.paramStructName == nil {
		s.Query.paramStructName = t
	} else {
		s.errors.AlreadyDeclaredError("param struct name", s.Query.paramStructName, t)
	}
}

// EnterNotNullParam is called when production notNullParam is entered.
func (s *parserListener) EnterNotNullParam(ctx *parser.NotNullParamContext) {
	s.currentNotNull = append(s.currentNotNull, newToken(ctx.BaseParserRuleContext))
}

// ExitNotNullParamsTag is called when production notNullParamsTag is exited.
func (s *parserListener) ExitNotNullParamsTag(ctx *parser.NotNullParamsTagContext) {
	s.Query.notNullParams = s.currentNotNull
	s.currentNotNull = nil
}

// EnterModeTag is called when production modeTag is entered.
func (s *parserListener) EnterModeTag(ctx *parser.ModeTagContext) {
	t := newToken(ctx.BaseParserRuleContext)
	if s.Query.execMode == nil {
		s.Query.execMode = t
	} else {
		s.errors.AlreadyDeclaredError("exec mode", s.Query.execMode, t)
	}
}

// EnterQueryName is called when production queryName is entered.
func (s *parserListener) EnterQueryName(ctx *parser.QueryNameContext) {
	t := newToken(ctx.BaseParserRuleContext)
	if s.Query.name == nil {
		s.Query.name = t
	} else {
		s.errors.AlreadyDeclaredError("query name", s.Query.name, t)
	}
}

// EnterReturnValueNameId is called when production returnValueNameId is entered.
func (s *parserListener) EnterReturnValueNameId(ctx *parser.ReturnValueNameIdContext) {
	t := newToken(ctx.BaseParserRuleContext)
	if s.Query.returnValueName == nil {
		s.Query.returnValueName = t
	} else {
		s.errors.AlreadyDeclaredError("return value name", s.Query.returnValueName, t)
	}
}

// EnterLineComment is called when production lineComment is entered.
func (s *parserListener) EnterLineComment(ctx *parser.LineCommentContext) {
	if s.Query.statement != nil {
		// Ignore this comment, as it's in the statement body. Only line comments
		// before the query are treated as docstrings.
		return
	}

	for _, c := range strings.Split(ctx.GetText(), "\n") {
		if len(c) < 2 {
			continue
		}

		line := c[2:]
		if len(line) > 0 && line[0] == ' ' {
			line = line[1:]
		}
		s.Query.Comments = append(s.Query.Comments, line)
	}
}

// EnterWord is called when production word is entered.
func (s *parserListener) EnterWord(ctx *parser.WordContext) {
	t := ctx.GetText()
	start := ctx.GetStart().GetStart()
	offset := 0
	for {
		idx := strings.Index(t[offset:], "%")
		if idx == -1 {
			break
		}
		s.Query.percents = append(s.Query.percents, &token{
			Start:  start + offset + idx,
			Stop:   start + offset + idx,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		})
		offset += idx + 1
		if offset >= len(t) {
			break
		}
	}
}

// EnterParamId is called when :<name> is encountered inside the query.
func (s *parserListener) EnterParamId(ctx *parser.ParamIdContext) {
	t := newToken(ctx.BaseParserRuleContext)

	input := ctx.GetStart()
	name := input.GetInputStream().GetText(t.Start, t.Stop)
	lname := strings.ToLower(name)

	param, ok := s.Query.params[lname]
	if ok {
		param.uses = append(param.uses, t)
	} else {
		s.Query.params[lname] = &Param{
			definition: &token{Value: name},
			Idx:        s.Query.paramIdx[Scalar],
			Type:       Scalar,
			uses:       []*token{t},
		}
		s.Query.paramIdx[Scalar]++
	}
}

// EnterStatementBody is called when production statementBody is entered.
func (s *parserListener) EnterStatementBody(ctx *parser.StatementBodyContext) {
	input := ctx.GetStart()
	t := newToken(ctx.BaseParserRuleContext)
	t.Value = input.GetInputStream().GetText(t.Start, t.Stop)
	s.Query.statement = t
}

func (s *parserListener) CheckUnusedParameters() error {
	errors := false
	for _, p := range s.Query.params {
		if len(p.uses) == 0 {
			errors = true
			s.errors.UnusedParamError(p)
		}
	}

	if errors {
		return s.errors.Error()
	}

	return nil
}

func (s *parserListener) PopulateNotNullParams() {
	for _, p := range s.Query.notNullParams {
		segments := strings.SplitN(strings.ToLower(p.Value), ".", 2)
		param, ok := s.Query.params[segments[0]]
		if !ok {
			s.errors.MissingParamError(p)
			continue
		}

		switch len(segments) {
		case 1:
			// This is 'plain' parameter
			param.NotNull = true
		case 2:
			// This is struct field in the struct spread param
			structkey, ok := param.keys[segments[1]]
			if !ok {
				s.errors.MissingParamError(p)
			} else {
				structkey.NotNull = true
			}
		default:
			panic("parser error")
		}
	}
}

func (s *parserListener) VerifyExecMode() {
	if s.Query.execMode == nil {
		s.errors.MissingExecModeError(s.Query.statement)
	}
}
