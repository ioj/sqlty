package compiler

import (
	"strings"
)

// Parser parses SQL files with annotations.
type Parser struct {
	lexer    *Lexer
	filename string
	errors   *errorListener
	input    string

	current Token
	peeked  bool
	peekTok Token

	query *Query

	// Temporary state during parsing
	currentParam           *Param
	currentStructTransform []*StructKey
	currentNotNull         []*token
}

// NewParser creates a new parser for the given input.
func NewParser(filename string, input string) *Parser {
	return &Parser{
		lexer:    NewLexer(input),
		filename: filename,
		errors:   newErrorListener(filename),
		input:    input,
	}
}

// next returns the next token, consuming any peeked token first.
func (p *Parser) next() Token {
	if p.peeked {
		p.peeked = false
		p.current = p.peekTok
		return p.current
	}
	p.current = p.lexer.NextToken()
	return p.current
}

// makeToken converts a lexer Token to the legacy token type.
func (p *Parser) makeToken(t Token) *token {
	return &token{
		Value:  t.Value,
		Start:  t.Start,
		Stop:   t.Stop,
		Line:   t.Line,
		Column: t.Column,
	}
}

// Parse parses the input and returns a Query.
func (p *Parser) Parse() (*Query, error) {
	p.query = &Query{
		Filename: p.filename,
		paramIdx: make(map[ParamType]int),
		params:   make(map[string]*Param),
	}

	// Get first token
	p.next()

	// Check for annotation block
	if p.current.Type != TokenOpenComment {
		// No annotation block - return nil (not an annotated query)
		return nil, nil
	}

	// Parse annotation block
	if err := p.parseAnnotationBlock(); err != nil {
		return nil, err
	}

	// Parse optional line comments (docstrings)
	p.parseLineComments()

	// Parse statement body
	if err := p.parseStatementBody(); err != nil {
		return nil, err
	}

	// Run validation
	p.checkUnusedParameters()
	p.populateNotNullParams()
	p.verifyExecMode()

	if p.query == nil {
		return nil, p.errors.Error()
	}

	// Check for empty file error
	for _, err := range p.errors.errors {
		if err.errtype == "empty" {
			return nil, ErrEmptyFile
		}
	}

	if err := p.errors.Error(); err != nil {
		return nil, err
	}

	return p.query, nil
}

// parseAnnotationBlock parses the content between /* and */.
func (p *Parser) parseAnnotationBlock() error {
	// We've already consumed /*
	p.next() // move past TokenOpenComment

	// Parse tags until we hit */
	for p.current.Type != TokenCloseComment && p.current.Type != TokenEOF {
		if err := p.parseAnnotation(); err != nil {
			return err
		}
	}

	if p.current.Type == TokenCloseComment {
		p.next() // consume */
	}

	return nil
}

// parseAnnotation parses a single annotation tag.
func (p *Parser) parseAnnotation() error {
	switch p.current.Type {
	case TokenAtName:
		return p.parseNameTag()
	case TokenAtParam:
		return p.parseParamTag()
	case TokenAtParamStructName:
		return p.parseParamStructNameTag()
	case TokenAtOne, TokenAtMany, TokenAtExec:
		return p.parseModeTag()
	case TokenAtNotNullParams:
		return p.parseNotNullParamsTag()
	case TokenAtReturnValueName:
		return p.parseReturnValueNameTag()
	case TokenAtTemplate:
		return p.parseTemplateTag()
	default:
		// Skip unknown tokens (like @paramAsStruct which is ignored)
		p.next()
		return nil
	}
}

// parseSimpleTag parses tags that expect a single identifier value.
// It handles duplicate detection and error reporting.
func (p *Parser) parseSimpleTag(dest **token, fieldName string) error {
	p.next() // consume @tag

	if p.current.Type != TokenIdentifier {
		return nil
	}

	t := p.makeToken(p.current)
	if *dest == nil {
		*dest = t
	} else {
		p.errors.AlreadyDeclaredError(fieldName, *dest, t)
	}

	p.next()
	return nil
}

// parseNameTag parses @name queryName.
func (p *Parser) parseNameTag() error {
	return p.parseSimpleTag(&p.query.name, "query name")
}

// parseParamTag parses @param name -> transform.
func (p *Parser) parseParamTag() error {
	paramStart := p.current
	p.next() // consume @param

	if p.current.Type != TokenIdentifier {
		p.next()
		return nil
	}

	// Create the param with definition token pointing to @param position
	p.currentParam = &Param{
		definition: p.makeToken(paramStart),
	}

	// Get the param name
	nameToken := p.current
	p.currentParam.definition.Value = nameToken.Value
	p.currentParam.definition.Start = nameToken.Start
	p.currentParam.definition.Stop = nameToken.Stop
	p.currentParam.definition.Line = nameToken.Line
	p.currentParam.definition.Column = nameToken.Column

	// Check for duplicate
	lname := strings.ToLower(nameToken.Value)
	if alreadyDefined, ok := p.query.params[lname]; ok {
		p.errors.AlreadyDeclaredError("parameter", alreadyDefined.definition, p.makeToken(nameToken))
		p.currentParam = nil
		p.next()
		return nil
	}

	p.next() // consume name

	// Expect ->
	if p.current.Type != TokenArrow {
		// No transform specified, might be an old/different syntax
		// Add as scalar param
		if p.currentParam != nil {
			p.query.params[lname] = p.currentParam
		}
		p.currentParam = nil
		return nil
	}
	p.next() // consume ->

	// Parse transform
	if err := p.parseTransform(); err != nil {
		return err
	}

	// Add param to query
	if p.currentParam != nil {
		lname := strings.ToLower(p.currentParam.definition.Value)
		p.query.params[lname] = p.currentParam
		p.currentParam = nil
	}

	return nil
}

// parseTransform parses (...) or ((fields)...).
func (p *Parser) parseTransform() error {
	if p.current.Type != TokenOpenParen {
		return nil
	}
	p.next() // consume (

	// Check for struct spread: ((field1, field2)...)
	if p.current.Type == TokenOpenParen {
		return p.parseStructSpreadTransform()
	}

	// Simple spread: (...)
	return p.parseSpreadTransform()
}

// parseSpreadTransform parses (...).
func (p *Parser) parseSpreadTransform() error {
	if p.current.Type != TokenSpread {
		return nil
	}
	p.next() // consume ...

	if p.current.Type == TokenCloseParen {
		p.next() // consume )
	}

	if p.currentParam != nil {
		p.currentParam.Type = Spread
		p.currentParam.Idx = p.query.paramIdx[Spread]
		p.query.paramIdx[Spread]++
	}

	return nil
}

// parseStructSpreadTransform parses ((field1, field2)...).
func (p *Parser) parseStructSpreadTransform() error {
	p.next() // consume inner (

	// Parse field list
	p.currentStructTransform = nil
	for p.current.Type == TokenIdentifier {
		structkey := &StructKey{
			token: p.makeToken(p.current),
			idx:   len(p.currentStructTransform),
		}
		p.currentStructTransform = append(p.currentStructTransform, structkey)
		p.next() // consume identifier

		if p.current.Type == TokenComma {
			p.next() // consume ,
		} else {
			break
		}
	}

	if p.current.Type == TokenCloseParen {
		p.next() // consume inner )
	}

	if p.current.Type == TokenSpread {
		p.next() // consume ...
	}

	if p.current.Type == TokenCloseParen {
		p.next() // consume outer )
	}

	// Finalize struct transform
	if p.currentParam != nil {
		p.currentParam.keys = make(map[string]*StructKey)

		for _, sk := range p.currentStructTransform {
			lname := strings.ToLower(sk.token.Value)
			if _, ok := p.currentParam.keys[lname]; ok {
				p.errors.DuplicateStructKeyError(p.currentParam.definition.Value, sk.token)
				continue
			}
			p.currentParam.keys[lname] = sk
		}

		p.currentParam.Type = StructSpread
		p.currentParam.Idx = p.query.paramIdx[StructSpread]
		p.query.paramIdx[StructSpread]++
	}

	p.currentStructTransform = nil
	return nil
}

// parseParamStructNameTag parses @paramStructName name.
func (p *Parser) parseParamStructNameTag() error {
	return p.parseSimpleTag(&p.query.paramStructName, "param struct name")
}

// parseModeTag parses @one, @many, or @exec.
func (p *Parser) parseModeTag() error {
	t := p.makeToken(p.current)

	if p.query.execMode == nil {
		p.query.execMode = t
	} else {
		p.errors.AlreadyDeclaredError("exec mode", p.query.execMode, t)
	}

	p.next()
	return nil
}

// parseNotNullParamsTag parses @notNullParams (param1, param2.field).
func (p *Parser) parseNotNullParamsTag() error {
	p.next() // consume @notNullParams

	// Old syntax: @notNullParams (a, b, c)
	// or just: @notNullParams a, b, c
	if p.current.Type == TokenOpenParen {
		p.next() // consume (
	}

	p.currentNotNull = nil
	for p.current.Type == TokenIdentifier {
		// Build the full name (could be param or param.field)
		var fullName strings.Builder
		start := p.current.Start
		line := p.current.Line
		col := p.current.Column

		fullName.WriteString(p.current.Value)
		p.next()

		// Check for .field
		if p.current.Type == TokenDot {
			fullName.WriteString(".")
			p.next() // consume .

			if p.current.Type == TokenIdentifier {
				fullName.WriteString(p.current.Value)
				p.next()
			}
		}

		p.currentNotNull = append(p.currentNotNull, &token{
			Value:  fullName.String(),
			Start:  start,
			Stop:   start + len(fullName.String()) - 1,
			Line:   line,
			Column: col,
		})

		if p.current.Type == TokenComma {
			p.next() // consume ,
		} else {
			break
		}
	}

	if p.current.Type == TokenCloseParen {
		p.next() // consume )
	}

	p.query.notNullParams = p.currentNotNull
	p.currentNotNull = nil
	return nil
}

// parseReturnValueNameTag parses @returnValueName name.
func (p *Parser) parseReturnValueNameTag() error {
	return p.parseSimpleTag(&p.query.returnValueName, "return value name")
}

// parseTemplateTag parses @template name.
func (p *Parser) parseTemplateTag() error {
	return p.parseSimpleTag(&p.query.template, "template name")
}

// parseLineComments collects line comments as docstrings.
func (p *Parser) parseLineComments() {
	for p.current.Type == TokenLineComment {
		text := p.current.Value

		// Process line comments, stripping -- prefix
		for _, c := range strings.Split(text, "\n") {
			if len(c) < 2 {
				continue
			}

			line := c[2:] // strip --
			if len(line) > 0 && line[0] == ' ' {
				line = line[1:]
			}
			// Remove trailing \r if present
			line = strings.TrimSuffix(line, "\r")
			if line != "" || len(p.query.Comments) > 0 {
				p.query.Comments = append(p.query.Comments, line)
			}
		}

		p.next()
	}
}

// parseStatementBody parses the SQL statement.
func (p *Parser) parseStatementBody() error {
	if p.current.Type == TokenEOF {
		p.errors.EmptyFileError()
		return nil
	}

	// Record start position
	startPos := p.current.Start
	startLine := p.current.Line
	startCol := p.current.Column

	// Collect all tokens until semicolon
	var endPos int

	for p.current.Type != TokenSemicolon && p.current.Type != TokenEOF {
		// Track percent signs for sprintf escaping
		if p.current.Type == TokenPercent {
			p.query.percents = append(p.query.percents, &token{
				Start:  p.current.Start,
				Stop:   p.current.Stop,
				Line:   p.current.Line,
				Column: p.current.Column,
			})
		}

		// Track parameters
		if p.current.Type == TokenParamMark {
			p.next()

			if p.current.Type == TokenIdentifier {
				p.trackParameter(p.current)
			}
			continue
		}

		// For words/strings, also scan for percent signs within the value
		if p.current.Type == TokenWord || p.current.Type == TokenString || p.current.Type == TokenIdentifier {
			p.scanForPercents(p.current)
		}

		endPos = p.current.Stop
		p.next()
	}

	if p.current.Type == TokenSemicolon {
		endPos = p.current.Start - 1
	}

	// Build statement token
	if startPos <= endPos && endPos < len(p.input) {
		statementText := p.input[startPos : endPos+1]
		p.query.statement = &token{
			Value:  statementText,
			Start:  startPos,
			Stop:   endPos,
			Line:   startLine,
			Column: startCol,
		}
	} else {
		p.errors.EmptyFileError()
	}

	return nil
}

// scanForPercents scans a token's value for percent signs.
func (p *Parser) scanForPercents(t Token) {
	text := t.Value
	offset := 0
	for {
		idx := strings.Index(text[offset:], "%")
		if idx == -1 {
			break
		}
		p.query.percents = append(p.query.percents, &token{
			Start:  t.Start + offset + idx,
			Stop:   t.Start + offset + idx,
			Line:   t.Line,
			Column: t.Column,
		})
		offset += idx + 1
		if offset >= len(text) {
			break
		}
	}
}

// trackParameter tracks a parameter usage in the SQL statement.
func (p *Parser) trackParameter(ident Token) {
	name := ident.Value
	lname := strings.ToLower(name)

	// Create token for the param identifier (not including the :)
	t := &token{
		Value:  name,
		Start:  ident.Start,
		Stop:   ident.Stop,
		Line:   ident.Line,
		Column: ident.Column,
	}

	param, ok := p.query.params[lname]
	if ok {
		param.uses = append(param.uses, t)
	} else {
		// Auto-create as scalar param
		p.query.params[lname] = &Param{
			definition: &token{Value: name},
			Idx:        p.query.paramIdx[Scalar],
			Type:       Scalar,
			uses:       []*token{t},
		}
		p.query.paramIdx[Scalar]++
	}
}

// checkUnusedParameters checks for declared but unused parameters.
func (p *Parser) checkUnusedParameters() {
	for _, param := range p.query.params {
		if len(param.uses) == 0 {
			p.errors.UnusedParamError(param)
		}
	}
}

// populateNotNullParams links not-null params to actual params.
func (p *Parser) populateNotNullParams() {
	for _, nnp := range p.query.notNullParams {
		segments := strings.SplitN(strings.ToLower(nnp.Value), ".", 2)
		param, ok := p.query.params[segments[0]]
		if !ok {
			p.errors.MissingParamError(nnp)
			continue
		}

		switch len(segments) {
		case 1:
			param.NotNull = true
		case 2:
			structkey, ok := param.keys[segments[1]]
			if !ok {
				p.errors.MissingParamError(nnp)
			} else {
				structkey.NotNull = true
			}
		}
	}
}

// verifyExecMode checks that exec mode was specified.
func (p *Parser) verifyExecMode() {
	if p.query.execMode == nil && p.query.statement != nil {
		p.errors.MissingExecModeError(p.query.statement)
	}
}
