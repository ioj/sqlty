package compiler

import (
	"unicode"
)

// TokenType represents the type of a lexer token.
type TokenType int

const (
	TokenEOF TokenType = iota
	TokenError

	// DEFAULT mode tokens
	TokenOpenComment // /*
	TokenLineComment // -- ... \n
	TokenParamMark   // :
	TokenIdentifier  // identifier (after : or in comment)
	TokenWord        // any word or operator sequence
	TokenString      // '...'
	TokenSemicolon   // ;
	TokenPercent     // %

	// COMMENT mode tokens
	TokenCloseComment      // */
	TokenAtName            // @name
	TokenAtParam           // @param
	TokenAtParamStructName // @paramStructName
	TokenAtOne             // @one
	TokenAtMany            // @many
	TokenAtExec            // @exec
	TokenAtNotNullParams   // @notNullParams
	TokenAtReturnValueName // @returnValueName
	TokenAtTemplate        // @template
	TokenOpenParen         // (
	TokenCloseParen        // )
	TokenComma             // ,
	TokenDot               // .
	TokenSpread            // ...
	TokenArrow             // ->
)

// LexerMode represents the current lexing context.
type LexerMode int

const (
	ModeDefault LexerMode = iota
	ModeComment
)

// Token represents a lexer token with position information.
type Token struct {
	Type   TokenType
	Value  string
	Start  int // byte offset in input
	Stop   int // byte offset in input (inclusive)
	Line   int // 1-based line number
	Column int // 0-based column number
}

// Lexer tokenizes SQL input with support for mode switching.
type Lexer struct {
	input  []rune
	pos    int       // current position in input
	line   int       // current line (1-based)
	column int       // current column (0-based)
	mode   LexerMode // current lexer mode

	// Token start position
	tokenStart  int
	tokenLine   int
	tokenColumn int
}

// NewLexer creates a new lexer for the given input string.
func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  []rune(input),
		pos:    0,
		line:   1,
		column: 0,
		mode:   ModeDefault,
	}
}

// peek returns the current rune without advancing.
func (l *Lexer) peek() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

// peekNext returns the next rune without advancing.
func (l *Lexer) peekNext() rune {
	if l.pos+1 >= len(l.input) {
		return 0
	}
	return l.input[l.pos+1]
}

// peekAt returns the rune at offset n from current position.
func (l *Lexer) peekAt(n int) rune {
	if n < 0 || l.pos+n >= len(l.input) {
		return 0
	}
	return l.input[l.pos+n]
}

// advance consumes and returns the current rune.
func (l *Lexer) advance() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	r := l.input[l.pos]
	l.pos++
	if r == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
	return r
}

// skipWhitespace consumes whitespace characters.
func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.input) && unicode.IsSpace(l.peek()) {
		l.advance()
	}
}

// markTokenStart saves the current position as the token start.
func (l *Lexer) markTokenStart() {
	l.tokenStart = l.pos
	l.tokenLine = l.line
	l.tokenColumn = l.column
}

// emitToken creates a token from the marked start to current position.
func (l *Lexer) emitToken(typ TokenType) Token {
	return Token{
		Type:   typ,
		Value:  string(l.input[l.tokenStart:l.pos]),
		Start:  l.tokenStart,
		Stop:   l.pos - 1,
		Line:   l.tokenLine,
		Column: l.tokenColumn,
	}
}

// emitError creates an error token with a message.
func (l *Lexer) emitError(msg string) Token {
	return Token{
		Type:   TokenError,
		Value:  msg,
		Start:  l.tokenStart,
		Stop:   l.pos - 1,
		Line:   l.tokenLine,
		Column: l.tokenColumn,
	}
}

// isIdentStart returns true if r can start an identifier.
func isIdentStart(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}

// isIdentChar returns true if r can be part of an identifier.
func isIdentChar(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

// NextToken returns the next token from the input.
func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	if l.pos >= len(l.input) {
		l.markTokenStart()
		return l.emitToken(TokenEOF)
	}

	if l.mode == ModeComment {
		return l.scanCommentMode()
	}
	return l.scanDefaultMode()
}

// scanDefaultMode handles tokenization in default (SQL) mode.
func (l *Lexer) scanDefaultMode() Token {
	l.markTokenStart()
	r := l.peek()

	// Check for block comment start
	if r == '/' && l.peekNext() == '*' {
		l.advance() // /
		l.advance() // *
		l.mode = ModeComment
		return l.emitToken(TokenOpenComment)
	}

	// Check for line comment
	if r == '-' && l.peekNext() == '-' {
		return l.scanLineComment()
	}

	// Check for parameter marker
	if r == ':' {
		// Check it's not :: (PostgreSQL cast)
		if l.peekNext() == ':' {
			l.advance() // :
			l.advance() // :
			return l.emitToken(TokenWord)
		}
		l.advance() // :
		return l.emitToken(TokenParamMark)
	}

	// Check for string literal
	if r == '\'' {
		return l.scanString()
	}

	// Check for semicolon
	if r == ';' {
		l.advance()
		return l.emitToken(TokenSemicolon)
	}

	// Check for percent
	if r == '%' {
		l.advance()
		return l.emitToken(TokenPercent)
	}

	// Scan identifier or word
	if isIdentStart(r) {
		return l.scanIdentifier()
	}

	// Scan any other word (operators, punctuation, etc.)
	return l.scanWord()
}

// scanCommentMode handles tokenization in comment (annotation) mode.
func (l *Lexer) scanCommentMode() Token {
	l.markTokenStart()
	r := l.peek()

	// Check for close comment
	if r == '*' && l.peekNext() == '/' {
		l.advance() // *
		l.advance() // /
		l.mode = ModeDefault
		return l.emitToken(TokenCloseComment)
	}

	// Check for @ tag
	if r == '@' {
		return l.scanAtTag()
	}

	// Check for spread (...)
	if r == '.' && l.peekNext() == '.' && l.peekAt(2) == '.' {
		l.advance() // .
		l.advance() // .
		l.advance() // .
		return l.emitToken(TokenSpread)
	}

	// Check for arrow (->)
	if r == '-' && l.peekNext() == '>' {
		l.advance() // -
		l.advance() // >
		return l.emitToken(TokenArrow)
	}

	// Check for single characters
	switch r {
	case '(':
		l.advance()
		return l.emitToken(TokenOpenParen)
	case ')':
		l.advance()
		return l.emitToken(TokenCloseParen)
	case ',':
		l.advance()
		return l.emitToken(TokenComma)
	case '.':
		l.advance()
		return l.emitToken(TokenDot)
	}

	// Scan identifier
	if isIdentStart(r) {
		return l.scanIdentifier()
	}

	// Skip unknown characters in comments (like numbers)
	l.advance()
	return l.NextToken()
}

// scanLineComment scans a line comment (-- ...).
func (l *Lexer) scanLineComment() Token {
	l.markTokenStart()
	// Consume --
	l.advance()
	l.advance()

	// Consume until end of line or EOF
	for l.pos < len(l.input) && l.peek() != '\n' {
		l.advance()
	}
	// Consume the newline if present
	if l.pos < len(l.input) && l.peek() == '\n' {
		l.advance()
	}

	return l.emitToken(TokenLineComment)
}

// scanString scans a single-quoted string literal.
func (l *Lexer) scanString() Token {
	l.markTokenStart()
	l.advance() // opening quote

	for l.pos < len(l.input) {
		r := l.peek()
		if r == '\'' {
			l.advance()
			// Check for escaped quote ('')
			if l.peek() == '\'' {
				l.advance()
				continue
			}
			return l.emitToken(TokenString)
		}
		l.advance()
	}

	return l.emitError("unterminated string literal")
}

// scanIdentifier scans an identifier.
func (l *Lexer) scanIdentifier() Token {
	l.markTokenStart()

	for l.pos < len(l.input) && isIdentChar(l.peek()) {
		l.advance()
	}

	return l.emitToken(TokenIdentifier)
}

// scanWord scans a word (any non-whitespace, non-special sequence).
func (l *Lexer) scanWord() Token {
	l.markTokenStart()

	// Consume until we hit whitespace or a special character
	for l.pos < len(l.input) {
		r := l.peek()
		if unicode.IsSpace(r) {
			break
		}
		// Stop at special characters
		if r == ';' || r == ':' || r == '\'' || r == '%' || r == '/' || r == '-' {
			// Check for comment starts
			if r == '/' && l.peekNext() == '*' {
				break
			}
			if r == '-' && l.peekNext() == '-' {
				break
			}
			// Check for :: (cast)
			if r == ':' && l.peekNext() != ':' {
				break
			}
		}
		l.advance()
	}

	if l.pos == l.tokenStart {
		// Consume at least one character
		l.advance()
	}

	return l.emitToken(TokenWord)
}

// scanAtTag scans an @ tag in comment mode.
func (l *Lexer) scanAtTag() Token {
	l.markTokenStart()
	l.advance() // @

	// Scan the tag name
	start := l.pos
	for l.pos < len(l.input) && isIdentChar(l.peek()) {
		l.advance()
	}

	tagName := string(l.input[start:l.pos])

	switch tagName {
	case "name":
		return l.emitToken(TokenAtName)
	case "param":
		return l.emitToken(TokenAtParam)
	case "paramStructName":
		return l.emitToken(TokenAtParamStructName)
	case "one":
		return l.emitToken(TokenAtOne)
	case "many":
		return l.emitToken(TokenAtMany)
	case "exec":
		return l.emitToken(TokenAtExec)
	case "notNullParams":
		return l.emitToken(TokenAtNotNullParams)
	case "returnValueName":
		return l.emitToken(TokenAtReturnValueName)
	case "template":
		return l.emitToken(TokenAtTemplate)
	default:
		// Unknown @ tag, just return as identifier
		return l.emitToken(TokenIdentifier)
	}
}

// Mode returns the current lexer mode.
func (l *Lexer) Mode() LexerMode {
	return l.mode
}

// Pos returns the current position.
func (l *Lexer) Pos() int {
	return l.pos
}

// Line returns the current line number.
func (l *Lexer) Line() int {
	return l.line
}

// Column returns the current column number.
func (l *Lexer) Column() int {
	return l.column
}
