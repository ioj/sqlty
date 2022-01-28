// Code generated from SQLParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // SQLParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 191,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	7, 4, 69, 10, 4, 12, 4, 14, 4, 72, 11, 4, 3, 4, 3, 4, 3, 5, 5, 5, 77, 10,
	5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6, 84, 10, 6, 12, 6, 14, 6, 87, 11,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 93, 10, 7, 12, 7, 14, 7, 96, 11, 7, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 137, 10, 20, 3, 21,
	3, 21, 5, 21, 141, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3,
	23, 3, 23, 7, 23, 151, 10, 23, 12, 23, 14, 23, 154, 11, 23, 3, 23, 5, 23,
	157, 10, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 25, 7, 25, 169, 10, 25, 12, 25, 14, 25, 172, 11, 25, 3, 25, 5, 25,
	175, 10, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 183, 10,
	27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 2, 2, 31, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 2, 4, 5, 2, 3, 3, 5, 5, 8, 8, 3, 2, 16,
	18, 2, 179, 2, 60, 3, 2, 2, 2, 4, 63, 3, 2, 2, 2, 6, 66, 3, 2, 2, 2, 8,
	76, 3, 2, 2, 2, 10, 81, 3, 2, 2, 2, 12, 88, 3, 2, 2, 2, 14, 97, 3, 2, 2,
	2, 16, 99, 3, 2, 2, 2, 18, 102, 3, 2, 2, 2, 20, 104, 3, 2, 2, 2, 22, 106,
	3, 2, 2, 2, 24, 108, 3, 2, 2, 2, 26, 111, 3, 2, 2, 2, 28, 115, 3, 2, 2,
	2, 30, 118, 3, 2, 2, 2, 32, 120, 3, 2, 2, 2, 34, 123, 3, 2, 2, 2, 36, 126,
	3, 2, 2, 2, 38, 136, 3, 2, 2, 2, 40, 140, 3, 2, 2, 2, 42, 142, 3, 2, 2,
	2, 44, 146, 3, 2, 2, 2, 46, 160, 3, 2, 2, 2, 48, 165, 3, 2, 2, 2, 50, 176,
	3, 2, 2, 2, 52, 182, 3, 2, 2, 2, 54, 184, 3, 2, 2, 2, 56, 186, 3, 2, 2,
	2, 58, 188, 3, 2, 2, 2, 60, 61, 5, 4, 3, 2, 61, 62, 7, 2, 2, 3, 62, 3,
	3, 2, 2, 2, 63, 64, 5, 6, 4, 2, 64, 65, 5, 8, 5, 2, 65, 5, 3, 2, 2, 2,
	66, 70, 7, 4, 2, 2, 67, 69, 5, 38, 20, 2, 68, 67, 3, 2, 2, 2, 69, 72, 3,
	2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 73, 3, 2, 2, 2, 72,
	70, 3, 2, 2, 2, 73, 74, 7, 27, 2, 2, 74, 7, 3, 2, 2, 2, 75, 77, 5, 10,
	6, 2, 76, 75, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79,
	5, 12, 7, 2, 79, 80, 7, 6, 2, 2, 80, 9, 3, 2, 2, 2, 81, 85, 7, 10, 2, 2,
	82, 84, 7, 10, 2, 2, 83, 82, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 3,
	2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 11, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 88,
	94, 5, 14, 8, 2, 89, 93, 5, 16, 9, 2, 90, 93, 5, 14, 8, 2, 91, 93, 5, 10,
	6, 2, 92, 89, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 91, 3, 2, 2, 2, 93, 96,
	3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 13, 3, 2, 2, 2,
	96, 94, 3, 2, 2, 2, 97, 98, 9, 2, 2, 2, 98, 15, 3, 2, 2, 2, 99, 100, 7,
	9, 2, 2, 100, 101, 5, 18, 10, 2, 101, 17, 3, 2, 2, 2, 102, 103, 7, 3, 2,
	2, 103, 19, 3, 2, 2, 2, 104, 105, 7, 3, 2, 2, 105, 21, 3, 2, 2, 2, 106,
	107, 7, 3, 2, 2, 107, 23, 3, 2, 2, 2, 108, 109, 7, 13, 2, 2, 109, 110,
	5, 54, 28, 2, 110, 25, 3, 2, 2, 2, 111, 112, 7, 14, 2, 2, 112, 113, 5,
	56, 29, 2, 113, 114, 5, 40, 21, 2, 114, 27, 3, 2, 2, 2, 115, 116, 7, 15,
	2, 2, 116, 117, 5, 22, 12, 2, 117, 29, 3, 2, 2, 2, 118, 119, 9, 3, 2, 2,
	119, 31, 3, 2, 2, 2, 120, 121, 7, 19, 2, 2, 121, 122, 5, 48, 25, 2, 122,
	33, 3, 2, 2, 2, 123, 124, 7, 20, 2, 2, 124, 125, 5, 20, 11, 2, 125, 35,
	3, 2, 2, 2, 126, 127, 7, 21, 2, 2, 127, 128, 5, 58, 30, 2, 128, 37, 3,
	2, 2, 2, 129, 137, 5, 24, 13, 2, 130, 137, 5, 26, 14, 2, 131, 137, 5, 28,
	15, 2, 132, 137, 5, 30, 16, 2, 133, 137, 5, 32, 17, 2, 134, 137, 5, 34,
	18, 2, 135, 137, 5, 36, 19, 2, 136, 129, 3, 2, 2, 2, 136, 130, 3, 2, 2,
	2, 136, 131, 3, 2, 2, 2, 136, 132, 3, 2, 2, 2, 136, 133, 3, 2, 2, 2, 136,
	134, 3, 2, 2, 2, 136, 135, 3, 2, 2, 2, 137, 39, 3, 2, 2, 2, 138, 141, 5,
	42, 22, 2, 139, 141, 5, 46, 24, 2, 140, 138, 3, 2, 2, 2, 140, 139, 3, 2,
	2, 2, 141, 41, 3, 2, 2, 2, 142, 143, 7, 22, 2, 2, 143, 144, 7, 12, 2, 2,
	144, 145, 7, 23, 2, 2, 145, 43, 3, 2, 2, 2, 146, 147, 7, 22, 2, 2, 147,
	152, 5, 50, 26, 2, 148, 149, 7, 25, 2, 2, 149, 151, 5, 50, 26, 2, 150,
	148, 3, 2, 2, 2, 151, 154, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 152, 153,
	3, 2, 2, 2, 153, 156, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 155, 157, 7, 25,
	2, 2, 156, 155, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2,
	158, 159, 7, 23, 2, 2, 159, 45, 3, 2, 2, 2, 160, 161, 7, 22, 2, 2, 161,
	162, 5, 44, 23, 2, 162, 163, 7, 12, 2, 2, 163, 164, 7, 23, 2, 2, 164, 47,
	3, 2, 2, 2, 165, 170, 5, 52, 27, 2, 166, 167, 7, 25, 2, 2, 167, 169, 5,
	52, 27, 2, 168, 166, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168, 3, 2,
	2, 2, 170, 171, 3, 2, 2, 2, 171, 174, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2,
	173, 175, 7, 25, 2, 2, 174, 173, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175,
	49, 3, 2, 2, 2, 176, 177, 7, 3, 2, 2, 177, 51, 3, 2, 2, 2, 178, 183, 7,
	3, 2, 2, 179, 180, 7, 3, 2, 2, 180, 181, 7, 24, 2, 2, 181, 183, 7, 3, 2,
	2, 182, 178, 3, 2, 2, 2, 182, 179, 3, 2, 2, 2, 183, 53, 3, 2, 2, 2, 184,
	185, 7, 3, 2, 2, 185, 55, 3, 2, 2, 2, 186, 187, 7, 3, 2, 2, 187, 57, 3,
	2, 2, 2, 188, 189, 7, 3, 2, 2, 189, 59, 3, 2, 2, 2, 14, 70, 76, 85, 92,
	94, 136, 140, 152, 156, 170, 174, 182,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'/*'", "", "';'", "", "", "':'", "", "", "'...'", "'@name'", "'@param'",
	"'@paramStructName'", "'@one'", "'@many'", "'@exec'", "'@notNullParams'",
	"'@returnValueName'", "'@template'", "'('", "')'", "'.'", "','", "", "'*/'",
	"'::'",
}
var symbolicNames = []string{
	"", "ID", "OPEN_COMMENT", "WORD", "EOF_STATEMENT", "WSL", "STRING", "PARAM_MARK",
	"LINE_COMMENT", "WS", "SPREAD", "NAME_TAG", "TYPE_TAG", "PARAM_STRUCT_NAME_TAG",
	"ONE_TAG", "MANY_TAG", "EXEC_TAG", "NOT_NULL_PARAMS_TAG", "RETURN_VALUE_NAME_TAG",
	"TEMPLATE_TAG", "OB", "CB", "DOT", "COMMA", "ANY", "CLOSE_COMMENT", "CAST",
}

var ruleNames = []string{
	"input", "query", "queryDef", "statement", "lineComment", "statementBody",
	"word", "param", "paramId", "returnValueNameId", "paramStructNameId", "nameTag",
	"paramTag", "paramStructNameTag", "modeTag", "notNullParamsTag", "returnValueName",
	"templateTag", "anyTag", "transformRule", "spreadTransform", "structTransform",
	"structSpreadTransform", "notNullTransform", "key", "notNullParam", "queryName",
	"paramName", "templateName",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SQLParser struct {
	*antlr.BaseParser
}

func NewSQLParser(input antlr.TokenStream) *SQLParser {
	this := new(SQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SQLParser.g4"

	return this
}

// SQLParser tokens.
const (
	SQLParserEOF                   = antlr.TokenEOF
	SQLParserID                    = 1
	SQLParserOPEN_COMMENT          = 2
	SQLParserWORD                  = 3
	SQLParserEOF_STATEMENT         = 4
	SQLParserWSL                   = 5
	SQLParserSTRING                = 6
	SQLParserPARAM_MARK            = 7
	SQLParserLINE_COMMENT          = 8
	SQLParserWS                    = 9
	SQLParserSPREAD                = 10
	SQLParserNAME_TAG              = 11
	SQLParserTYPE_TAG              = 12
	SQLParserPARAM_STRUCT_NAME_TAG = 13
	SQLParserONE_TAG               = 14
	SQLParserMANY_TAG              = 15
	SQLParserEXEC_TAG              = 16
	SQLParserNOT_NULL_PARAMS_TAG   = 17
	SQLParserRETURN_VALUE_NAME_TAG = 18
	SQLParserTEMPLATE_TAG          = 19
	SQLParserOB                    = 20
	SQLParserCB                    = 21
	SQLParserDOT                   = 22
	SQLParserCOMMA                 = 23
	SQLParserANY                   = 24
	SQLParserCLOSE_COMMENT         = 25
	SQLParserCAST                  = 26
)

// SQLParser rules.
const (
	SQLParserRULE_input                 = 0
	SQLParserRULE_query                 = 1
	SQLParserRULE_queryDef              = 2
	SQLParserRULE_statement             = 3
	SQLParserRULE_lineComment           = 4
	SQLParserRULE_statementBody         = 5
	SQLParserRULE_word                  = 6
	SQLParserRULE_param                 = 7
	SQLParserRULE_paramId               = 8
	SQLParserRULE_returnValueNameId     = 9
	SQLParserRULE_paramStructNameId     = 10
	SQLParserRULE_nameTag               = 11
	SQLParserRULE_paramTag              = 12
	SQLParserRULE_paramStructNameTag    = 13
	SQLParserRULE_modeTag               = 14
	SQLParserRULE_notNullParamsTag      = 15
	SQLParserRULE_returnValueName       = 16
	SQLParserRULE_templateTag           = 17
	SQLParserRULE_anyTag                = 18
	SQLParserRULE_transformRule         = 19
	SQLParserRULE_spreadTransform       = 20
	SQLParserRULE_structTransform       = 21
	SQLParserRULE_structSpreadTransform = 22
	SQLParserRULE_notNullTransform      = 23
	SQLParserRULE_key                   = 24
	SQLParserRULE_notNullParam          = 25
	SQLParserRULE_queryName             = 26
	SQLParserRULE_paramName             = 27
	SQLParserRULE_templateName          = 28
)

// IInputContext is an interface to support dynamic dispatch.
type IInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputContext differentiates from other interfaces.
	IsInputContext()
}

type InputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputContext() *InputContext {
	var p = new(InputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_input
	return p
}

func (*InputContext) IsInputContext() {}

func NewInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputContext {
	var p = new(InputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_input

	return p
}

func (s *InputContext) GetParser() antlr.Parser { return s.parser }

func (s *InputContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *InputContext) EOF() antlr.TerminalNode {
	return s.GetToken(SQLParserEOF, 0)
}

func (s *InputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterInput(s)
	}
}

func (s *InputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitInput(s)
	}
}

func (p *SQLParser) Input() (localctx IInputContext) {
	localctx = NewInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SQLParserRULE_input)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Query()
	}
	{
		p.SetState(59)
		p.Match(SQLParserEOF)
	}

	return localctx
}

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) QueryDef() IQueryDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryDefContext)
}

func (s *QueryContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (p *SQLParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SQLParserRULE_query)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.QueryDef()
	}
	{
		p.SetState(62)
		p.Statement()
	}

	return localctx
}

// IQueryDefContext is an interface to support dynamic dispatch.
type IQueryDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryDefContext differentiates from other interfaces.
	IsQueryDefContext()
}

type QueryDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryDefContext() *QueryDefContext {
	var p = new(QueryDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_queryDef
	return p
}

func (*QueryDefContext) IsQueryDefContext() {}

func NewQueryDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryDefContext {
	var p = new(QueryDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_queryDef

	return p
}

func (s *QueryDefContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryDefContext) OPEN_COMMENT() antlr.TerminalNode {
	return s.GetToken(SQLParserOPEN_COMMENT, 0)
}

func (s *QueryDefContext) CLOSE_COMMENT() antlr.TerminalNode {
	return s.GetToken(SQLParserCLOSE_COMMENT, 0)
}

func (s *QueryDefContext) AllAnyTag() []IAnyTagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnyTagContext)(nil)).Elem())
	var tst = make([]IAnyTagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnyTagContext)
		}
	}

	return tst
}

func (s *QueryDefContext) AnyTag(i int) IAnyTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnyTagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnyTagContext)
}

func (s *QueryDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterQueryDef(s)
	}
}

func (s *QueryDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitQueryDef(s)
	}
}

func (p *SQLParser) QueryDef() (localctx IQueryDefContext) {
	localctx = NewQueryDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SQLParserRULE_queryDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(SQLParserOPEN_COMMENT)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SQLParserNAME_TAG)|(1<<SQLParserTYPE_TAG)|(1<<SQLParserPARAM_STRUCT_NAME_TAG)|(1<<SQLParserONE_TAG)|(1<<SQLParserMANY_TAG)|(1<<SQLParserEXEC_TAG)|(1<<SQLParserNOT_NULL_PARAMS_TAG)|(1<<SQLParserRETURN_VALUE_NAME_TAG)|(1<<SQLParserTEMPLATE_TAG))) != 0 {
		{
			p.SetState(65)
			p.AnyTag()
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(71)
		p.Match(SQLParserCLOSE_COMMENT)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) StatementBody() IStatementBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementBodyContext)
}

func (s *StatementContext) EOF_STATEMENT() antlr.TerminalNode {
	return s.GetToken(SQLParserEOF_STATEMENT, 0)
}

func (s *StatementContext) LineComment() ILineCommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineCommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineCommentContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *SQLParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SQLParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserLINE_COMMENT {
		{
			p.SetState(73)
			p.LineComment()
		}

	}
	{
		p.SetState(76)
		p.StatementBody()
	}
	{
		p.SetState(77)
		p.Match(SQLParserEOF_STATEMENT)
	}

	return localctx
}

// ILineCommentContext is an interface to support dynamic dispatch.
type ILineCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineCommentContext differentiates from other interfaces.
	IsLineCommentContext()
}

type LineCommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineCommentContext() *LineCommentContext {
	var p = new(LineCommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_lineComment
	return p
}

func (*LineCommentContext) IsLineCommentContext() {}

func NewLineCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineCommentContext {
	var p = new(LineCommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_lineComment

	return p
}

func (s *LineCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *LineCommentContext) AllLINE_COMMENT() []antlr.TerminalNode {
	return s.GetTokens(SQLParserLINE_COMMENT)
}

func (s *LineCommentContext) LINE_COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserLINE_COMMENT, i)
}

func (s *LineCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterLineComment(s)
	}
}

func (s *LineCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitLineComment(s)
	}
}

func (p *SQLParser) LineComment() (localctx ILineCommentContext) {
	localctx = NewLineCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SQLParserRULE_lineComment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(SQLParserLINE_COMMENT)
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(80)
				p.Match(SQLParserLINE_COMMENT)
			}

		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IStatementBodyContext is an interface to support dynamic dispatch.
type IStatementBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementBodyContext differentiates from other interfaces.
	IsStatementBodyContext()
}

type StatementBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementBodyContext() *StatementBodyContext {
	var p = new(StatementBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_statementBody
	return p
}

func (*StatementBodyContext) IsStatementBodyContext() {}

func NewStatementBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementBodyContext {
	var p = new(StatementBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_statementBody

	return p
}

func (s *StatementBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementBodyContext) AllWord() []IWordContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWordContext)(nil)).Elem())
	var tst = make([]IWordContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWordContext)
		}
	}

	return tst
}

func (s *StatementBodyContext) Word(i int) IWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWordContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *StatementBodyContext) AllParam() []IParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamContext)(nil)).Elem())
	var tst = make([]IParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamContext)
		}
	}

	return tst
}

func (s *StatementBodyContext) Param(i int) IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *StatementBodyContext) AllLineComment() []ILineCommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineCommentContext)(nil)).Elem())
	var tst = make([]ILineCommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineCommentContext)
		}
	}

	return tst
}

func (s *StatementBodyContext) LineComment(i int) ILineCommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineCommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineCommentContext)
}

func (s *StatementBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterStatementBody(s)
	}
}

func (s *StatementBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitStatementBody(s)
	}
}

func (p *SQLParser) StatementBody() (localctx IStatementBodyContext) {
	localctx = NewStatementBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SQLParserRULE_statementBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Word()
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SQLParserID)|(1<<SQLParserWORD)|(1<<SQLParserSTRING)|(1<<SQLParserPARAM_MARK)|(1<<SQLParserLINE_COMMENT))) != 0 {
		p.SetState(90)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SQLParserPARAM_MARK:
			{
				p.SetState(87)
				p.Param()
			}

		case SQLParserID, SQLParserWORD, SQLParserSTRING:
			{
				p.SetState(88)
				p.Word()
			}

		case SQLParserLINE_COMMENT:
			{
				p.SetState(89)
				p.LineComment()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IWordContext is an interface to support dynamic dispatch.
type IWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWordContext differentiates from other interfaces.
	IsWordContext()
}

type WordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordContext() *WordContext {
	var p = new(WordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_word
	return p
}

func (*WordContext) IsWordContext() {}

func NewWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordContext {
	var p = new(WordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_word

	return p
}

func (s *WordContext) GetParser() antlr.Parser { return s.parser }

func (s *WordContext) WORD() antlr.TerminalNode {
	return s.GetToken(SQLParserWORD, 0)
}

func (s *WordContext) ID() antlr.TerminalNode {
	return s.GetToken(SQLParserID, 0)
}

func (s *WordContext) STRING() antlr.TerminalNode {
	return s.GetToken(SQLParserSTRING, 0)
}

func (s *WordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterWord(s)
	}
}

func (s *WordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitWord(s)
	}
}

func (p *SQLParser) Word() (localctx IWordContext) {
	localctx = NewWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SQLParserRULE_word)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SQLParserID)|(1<<SQLParserWORD)|(1<<SQLParserSTRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) PARAM_MARK() antlr.TerminalNode {
	return s.GetToken(SQLParserPARAM_MARK, 0)
}

func (s *ParamContext) ParamId() IParamIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamIdContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *SQLParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SQLParserRULE_param)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(SQLParserPARAM_MARK)
	}
	{
		p.SetState(98)
		p.ParamId()
	}

	return localctx
}

// IParamIdContext is an interface to support dynamic dispatch.
type IParamIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamIdContext differentiates from other interfaces.
	IsParamIdContext()
}

type ParamIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamIdContext() *ParamIdContext {
	var p = new(ParamIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_paramId
	return p
}

func (*ParamIdContext) IsParamIdContext() {}

func NewParamIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamIdContext {
	var p = new(ParamIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_paramId

	return p
}

func (s *ParamIdContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamIdContext) ID() antlr.TerminalNode {
	return s.GetToken(SQLParserID, 0)
}

func (s *ParamIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterParamId(s)
	}
}

func (s *ParamIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitParamId(s)
	}
}

func (p *SQLParser) ParamId() (localctx IParamIdContext) {
	localctx = NewParamIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SQLParserRULE_paramId)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(SQLParserID)
	}

	return localctx
}

// IReturnValueNameIdContext is an interface to support dynamic dispatch.
type IReturnValueNameIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnValueNameIdContext differentiates from other interfaces.
	IsReturnValueNameIdContext()
}

type ReturnValueNameIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnValueNameIdContext() *ReturnValueNameIdContext {
	var p = new(ReturnValueNameIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_returnValueNameId
	return p
}

func (*ReturnValueNameIdContext) IsReturnValueNameIdContext() {}

func NewReturnValueNameIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnValueNameIdContext {
	var p = new(ReturnValueNameIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_returnValueNameId

	return p
}

func (s *ReturnValueNameIdContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnValueNameIdContext) ID() antlr.TerminalNode {
	return s.GetToken(SQLParserID, 0)
}

func (s *ReturnValueNameIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnValueNameIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnValueNameIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterReturnValueNameId(s)
	}
}

func (s *ReturnValueNameIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitReturnValueNameId(s)
	}
}

func (p *SQLParser) ReturnValueNameId() (localctx IReturnValueNameIdContext) {
	localctx = NewReturnValueNameIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SQLParserRULE_returnValueNameId)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(SQLParserID)
	}

	return localctx
}

// IParamStructNameIdContext is an interface to support dynamic dispatch.
type IParamStructNameIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamStructNameIdContext differentiates from other interfaces.
	IsParamStructNameIdContext()
}

type ParamStructNameIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamStructNameIdContext() *ParamStructNameIdContext {
	var p = new(ParamStructNameIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_paramStructNameId
	return p
}

func (*ParamStructNameIdContext) IsParamStructNameIdContext() {}

func NewParamStructNameIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamStructNameIdContext {
	var p = new(ParamStructNameIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_paramStructNameId

	return p
}

func (s *ParamStructNameIdContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamStructNameIdContext) ID() antlr.TerminalNode {
	return s.GetToken(SQLParserID, 0)
}

func (s *ParamStructNameIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamStructNameIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamStructNameIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterParamStructNameId(s)
	}
}

func (s *ParamStructNameIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitParamStructNameId(s)
	}
}

func (p *SQLParser) ParamStructNameId() (localctx IParamStructNameIdContext) {
	localctx = NewParamStructNameIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SQLParserRULE_paramStructNameId)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(SQLParserID)
	}

	return localctx
}

// INameTagContext is an interface to support dynamic dispatch.
type INameTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameTagContext differentiates from other interfaces.
	IsNameTagContext()
}

type NameTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameTagContext() *NameTagContext {
	var p = new(NameTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_nameTag
	return p
}

func (*NameTagContext) IsNameTagContext() {}

func NewNameTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameTagContext {
	var p = new(NameTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_nameTag

	return p
}

func (s *NameTagContext) GetParser() antlr.Parser { return s.parser }

func (s *NameTagContext) NAME_TAG() antlr.TerminalNode {
	return s.GetToken(SQLParserNAME_TAG, 0)
}

func (s *NameTagContext) QueryName() IQueryNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryNameContext)
}

func (s *NameTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterNameTag(s)
	}
}

func (s *NameTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitNameTag(s)
	}
}

func (p *SQLParser) NameTag() (localctx INameTagContext) {
	localctx = NewNameTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SQLParserRULE_nameTag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(SQLParserNAME_TAG)
	}
	{
		p.SetState(107)
		p.QueryName()
	}

	return localctx
}

// IParamTagContext is an interface to support dynamic dispatch.
type IParamTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamTagContext differentiates from other interfaces.
	IsParamTagContext()
}

type ParamTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamTagContext() *ParamTagContext {
	var p = new(ParamTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_paramTag
	return p
}

func (*ParamTagContext) IsParamTagContext() {}

func NewParamTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamTagContext {
	var p = new(ParamTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_paramTag

	return p
}

func (s *ParamTagContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamTagContext) TYPE_TAG() antlr.TerminalNode {
	return s.GetToken(SQLParserTYPE_TAG, 0)
}

func (s *ParamTagContext) ParamName() IParamNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamNameContext)
}

func (s *ParamTagContext) TransformRule() ITransformRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransformRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransformRuleContext)
}

func (s *ParamTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterParamTag(s)
	}
}

func (s *ParamTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitParamTag(s)
	}
}

func (p *SQLParser) ParamTag() (localctx IParamTagContext) {
	localctx = NewParamTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SQLParserRULE_paramTag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(SQLParserTYPE_TAG)
	}
	{
		p.SetState(110)
		p.ParamName()
	}
	{
		p.SetState(111)
		p.TransformRule()
	}

	return localctx
}

// IParamStructNameTagContext is an interface to support dynamic dispatch.
type IParamStructNameTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamStructNameTagContext differentiates from other interfaces.
	IsParamStructNameTagContext()
}

type ParamStructNameTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamStructNameTagContext() *ParamStructNameTagContext {
	var p = new(ParamStructNameTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_paramStructNameTag
	return p
}

func (*ParamStructNameTagContext) IsParamStructNameTagContext() {}

func NewParamStructNameTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamStructNameTagContext {
	var p = new(ParamStructNameTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_paramStructNameTag

	return p
}

func (s *ParamStructNameTagContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamStructNameTagContext) PARAM_STRUCT_NAME_TAG() antlr.TerminalNode {
	return s.GetToken(SQLParserPARAM_STRUCT_NAME_TAG, 0)
}

func (s *ParamStructNameTagContext) ParamStructNameId() IParamStructNameIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamStructNameIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamStructNameIdContext)
}

func (s *ParamStructNameTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamStructNameTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamStructNameTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterParamStructNameTag(s)
	}
}

func (s *ParamStructNameTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitParamStructNameTag(s)
	}
}

func (p *SQLParser) ParamStructNameTag() (localctx IParamStructNameTagContext) {
	localctx = NewParamStructNameTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SQLParserRULE_paramStructNameTag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(SQLParserPARAM_STRUCT_NAME_TAG)
	}
	{
		p.SetState(114)
		p.ParamStructNameId()
	}

	return localctx
}

// IModeTagContext is an interface to support dynamic dispatch.
type IModeTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModeTagContext differentiates from other interfaces.
	IsModeTagContext()
}

type ModeTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModeTagContext() *ModeTagContext {
	var p = new(ModeTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_modeTag
	return p
}

func (*ModeTagContext) IsModeTagContext() {}

func NewModeTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModeTagContext {
	var p = new(ModeTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_modeTag

	return p
}

func (s *ModeTagContext) GetParser() antlr.Parser { return s.parser }

func (s *ModeTagContext) ONE_TAG() antlr.TerminalNode {
	return s.GetToken(SQLParserONE_TAG, 0)
}

func (s *ModeTagContext) MANY_TAG() antlr.TerminalNode {
	return s.GetToken(SQLParserMANY_TAG, 0)
}

func (s *ModeTagContext) EXEC_TAG() antlr.TerminalNode {
	return s.GetToken(SQLParserEXEC_TAG, 0)
}

func (s *ModeTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModeTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModeTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterModeTag(s)
	}
}

func (s *ModeTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitModeTag(s)
	}
}

func (p *SQLParser) ModeTag() (localctx IModeTagContext) {
	localctx = NewModeTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SQLParserRULE_modeTag)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SQLParserONE_TAG)|(1<<SQLParserMANY_TAG)|(1<<SQLParserEXEC_TAG))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INotNullParamsTagContext is an interface to support dynamic dispatch.
type INotNullParamsTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotNullParamsTagContext differentiates from other interfaces.
	IsNotNullParamsTagContext()
}

type NotNullParamsTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotNullParamsTagContext() *NotNullParamsTagContext {
	var p = new(NotNullParamsTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_notNullParamsTag
	return p
}

func (*NotNullParamsTagContext) IsNotNullParamsTagContext() {}

func NewNotNullParamsTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotNullParamsTagContext {
	var p = new(NotNullParamsTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_notNullParamsTag

	return p
}

func (s *NotNullParamsTagContext) GetParser() antlr.Parser { return s.parser }

func (s *NotNullParamsTagContext) NOT_NULL_PARAMS_TAG() antlr.TerminalNode {
	return s.GetToken(SQLParserNOT_NULL_PARAMS_TAG, 0)
}

func (s *NotNullParamsTagContext) NotNullTransform() INotNullTransformContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotNullTransformContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotNullTransformContext)
}

func (s *NotNullParamsTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotNullParamsTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotNullParamsTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterNotNullParamsTag(s)
	}
}

func (s *NotNullParamsTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitNotNullParamsTag(s)
	}
}

func (p *SQLParser) NotNullParamsTag() (localctx INotNullParamsTagContext) {
	localctx = NewNotNullParamsTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SQLParserRULE_notNullParamsTag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(SQLParserNOT_NULL_PARAMS_TAG)
	}
	{
		p.SetState(119)
		p.NotNullTransform()
	}

	return localctx
}

// IReturnValueNameContext is an interface to support dynamic dispatch.
type IReturnValueNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnValueNameContext differentiates from other interfaces.
	IsReturnValueNameContext()
}

type ReturnValueNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnValueNameContext() *ReturnValueNameContext {
	var p = new(ReturnValueNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_returnValueName
	return p
}

func (*ReturnValueNameContext) IsReturnValueNameContext() {}

func NewReturnValueNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnValueNameContext {
	var p = new(ReturnValueNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_returnValueName

	return p
}

func (s *ReturnValueNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnValueNameContext) RETURN_VALUE_NAME_TAG() antlr.TerminalNode {
	return s.GetToken(SQLParserRETURN_VALUE_NAME_TAG, 0)
}

func (s *ReturnValueNameContext) ReturnValueNameId() IReturnValueNameIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnValueNameIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnValueNameIdContext)
}

func (s *ReturnValueNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnValueNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnValueNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterReturnValueName(s)
	}
}

func (s *ReturnValueNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitReturnValueName(s)
	}
}

func (p *SQLParser) ReturnValueName() (localctx IReturnValueNameContext) {
	localctx = NewReturnValueNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SQLParserRULE_returnValueName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(SQLParserRETURN_VALUE_NAME_TAG)
	}
	{
		p.SetState(122)
		p.ReturnValueNameId()
	}

	return localctx
}

// ITemplateTagContext is an interface to support dynamic dispatch.
type ITemplateTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateTagContext differentiates from other interfaces.
	IsTemplateTagContext()
}

type TemplateTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateTagContext() *TemplateTagContext {
	var p = new(TemplateTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_templateTag
	return p
}

func (*TemplateTagContext) IsTemplateTagContext() {}

func NewTemplateTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateTagContext {
	var p = new(TemplateTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_templateTag

	return p
}

func (s *TemplateTagContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateTagContext) TEMPLATE_TAG() antlr.TerminalNode {
	return s.GetToken(SQLParserTEMPLATE_TAG, 0)
}

func (s *TemplateTagContext) TemplateName() ITemplateNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateNameContext)
}

func (s *TemplateTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterTemplateTag(s)
	}
}

func (s *TemplateTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitTemplateTag(s)
	}
}

func (p *SQLParser) TemplateTag() (localctx ITemplateTagContext) {
	localctx = NewTemplateTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SQLParserRULE_templateTag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(SQLParserTEMPLATE_TAG)
	}
	{
		p.SetState(125)
		p.TemplateName()
	}

	return localctx
}

// IAnyTagContext is an interface to support dynamic dispatch.
type IAnyTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnyTagContext differentiates from other interfaces.
	IsAnyTagContext()
}

type AnyTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnyTagContext() *AnyTagContext {
	var p = new(AnyTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_anyTag
	return p
}

func (*AnyTagContext) IsAnyTagContext() {}

func NewAnyTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyTagContext {
	var p = new(AnyTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_anyTag

	return p
}

func (s *AnyTagContext) GetParser() antlr.Parser { return s.parser }

func (s *AnyTagContext) NameTag() INameTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameTagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameTagContext)
}

func (s *AnyTagContext) ParamTag() IParamTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamTagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamTagContext)
}

func (s *AnyTagContext) ParamStructNameTag() IParamStructNameTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamStructNameTagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamStructNameTagContext)
}

func (s *AnyTagContext) ModeTag() IModeTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModeTagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModeTagContext)
}

func (s *AnyTagContext) NotNullParamsTag() INotNullParamsTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotNullParamsTagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotNullParamsTagContext)
}

func (s *AnyTagContext) ReturnValueName() IReturnValueNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnValueNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnValueNameContext)
}

func (s *AnyTagContext) TemplateTag() ITemplateTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateTagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateTagContext)
}

func (s *AnyTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnyTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterAnyTag(s)
	}
}

func (s *AnyTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitAnyTag(s)
	}
}

func (p *SQLParser) AnyTag() (localctx IAnyTagContext) {
	localctx = NewAnyTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SQLParserRULE_anyTag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserNAME_TAG:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.NameTag()
		}

	case SQLParserTYPE_TAG:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.ParamTag()
		}

	case SQLParserPARAM_STRUCT_NAME_TAG:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			p.ParamStructNameTag()
		}

	case SQLParserONE_TAG, SQLParserMANY_TAG, SQLParserEXEC_TAG:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(130)
			p.ModeTag()
		}

	case SQLParserNOT_NULL_PARAMS_TAG:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(131)
			p.NotNullParamsTag()
		}

	case SQLParserRETURN_VALUE_NAME_TAG:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(132)
			p.ReturnValueName()
		}

	case SQLParserTEMPLATE_TAG:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(133)
			p.TemplateTag()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITransformRuleContext is an interface to support dynamic dispatch.
type ITransformRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransformRuleContext differentiates from other interfaces.
	IsTransformRuleContext()
}

type TransformRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransformRuleContext() *TransformRuleContext {
	var p = new(TransformRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_transformRule
	return p
}

func (*TransformRuleContext) IsTransformRuleContext() {}

func NewTransformRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransformRuleContext {
	var p = new(TransformRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_transformRule

	return p
}

func (s *TransformRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *TransformRuleContext) SpreadTransform() ISpreadTransformContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpreadTransformContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpreadTransformContext)
}

func (s *TransformRuleContext) StructSpreadTransform() IStructSpreadTransformContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructSpreadTransformContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructSpreadTransformContext)
}

func (s *TransformRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransformRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransformRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterTransformRule(s)
	}
}

func (s *TransformRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitTransformRule(s)
	}
}

func (p *SQLParser) TransformRule() (localctx ITransformRuleContext) {
	localctx = NewTransformRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SQLParserRULE_transformRule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.SpreadTransform()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.StructSpreadTransform()
		}

	}

	return localctx
}

// ISpreadTransformContext is an interface to support dynamic dispatch.
type ISpreadTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpreadTransformContext differentiates from other interfaces.
	IsSpreadTransformContext()
}

type SpreadTransformContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpreadTransformContext() *SpreadTransformContext {
	var p = new(SpreadTransformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_spreadTransform
	return p
}

func (*SpreadTransformContext) IsSpreadTransformContext() {}

func NewSpreadTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpreadTransformContext {
	var p = new(SpreadTransformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_spreadTransform

	return p
}

func (s *SpreadTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *SpreadTransformContext) OB() antlr.TerminalNode {
	return s.GetToken(SQLParserOB, 0)
}

func (s *SpreadTransformContext) SPREAD() antlr.TerminalNode {
	return s.GetToken(SQLParserSPREAD, 0)
}

func (s *SpreadTransformContext) CB() antlr.TerminalNode {
	return s.GetToken(SQLParserCB, 0)
}

func (s *SpreadTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpreadTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpreadTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterSpreadTransform(s)
	}
}

func (s *SpreadTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitSpreadTransform(s)
	}
}

func (p *SQLParser) SpreadTransform() (localctx ISpreadTransformContext) {
	localctx = NewSpreadTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SQLParserRULE_spreadTransform)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(SQLParserOB)
	}
	{
		p.SetState(141)
		p.Match(SQLParserSPREAD)
	}
	{
		p.SetState(142)
		p.Match(SQLParserCB)
	}

	return localctx
}

// IStructTransformContext is an interface to support dynamic dispatch.
type IStructTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructTransformContext differentiates from other interfaces.
	IsStructTransformContext()
}

type StructTransformContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructTransformContext() *StructTransformContext {
	var p = new(StructTransformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_structTransform
	return p
}

func (*StructTransformContext) IsStructTransformContext() {}

func NewStructTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructTransformContext {
	var p = new(StructTransformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_structTransform

	return p
}

func (s *StructTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *StructTransformContext) OB() antlr.TerminalNode {
	return s.GetToken(SQLParserOB, 0)
}

func (s *StructTransformContext) AllKey() []IKeyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyContext)(nil)).Elem())
	var tst = make([]IKeyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyContext)
		}
	}

	return tst
}

func (s *StructTransformContext) Key(i int) IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *StructTransformContext) CB() antlr.TerminalNode {
	return s.GetToken(SQLParserCB, 0)
}

func (s *StructTransformContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *StructTransformContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *StructTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterStructTransform(s)
	}
}

func (s *StructTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitStructTransform(s)
	}
}

func (p *SQLParser) StructTransform() (localctx IStructTransformContext) {
	localctx = NewStructTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SQLParserRULE_structTransform)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(SQLParserOB)
	}
	{
		p.SetState(145)
		p.Key()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(146)
				p.Match(SQLParserCOMMA)
			}
			{
				p.SetState(147)
				p.Key()
			}

		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserCOMMA {
		{
			p.SetState(153)
			p.Match(SQLParserCOMMA)
		}

	}
	{
		p.SetState(156)
		p.Match(SQLParserCB)
	}

	return localctx
}

// IStructSpreadTransformContext is an interface to support dynamic dispatch.
type IStructSpreadTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructSpreadTransformContext differentiates from other interfaces.
	IsStructSpreadTransformContext()
}

type StructSpreadTransformContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructSpreadTransformContext() *StructSpreadTransformContext {
	var p = new(StructSpreadTransformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_structSpreadTransform
	return p
}

func (*StructSpreadTransformContext) IsStructSpreadTransformContext() {}

func NewStructSpreadTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructSpreadTransformContext {
	var p = new(StructSpreadTransformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_structSpreadTransform

	return p
}

func (s *StructSpreadTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *StructSpreadTransformContext) OB() antlr.TerminalNode {
	return s.GetToken(SQLParserOB, 0)
}

func (s *StructSpreadTransformContext) StructTransform() IStructTransformContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructTransformContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructTransformContext)
}

func (s *StructSpreadTransformContext) SPREAD() antlr.TerminalNode {
	return s.GetToken(SQLParserSPREAD, 0)
}

func (s *StructSpreadTransformContext) CB() antlr.TerminalNode {
	return s.GetToken(SQLParserCB, 0)
}

func (s *StructSpreadTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructSpreadTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructSpreadTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterStructSpreadTransform(s)
	}
}

func (s *StructSpreadTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitStructSpreadTransform(s)
	}
}

func (p *SQLParser) StructSpreadTransform() (localctx IStructSpreadTransformContext) {
	localctx = NewStructSpreadTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SQLParserRULE_structSpreadTransform)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(SQLParserOB)
	}
	{
		p.SetState(159)
		p.StructTransform()
	}
	{
		p.SetState(160)
		p.Match(SQLParserSPREAD)
	}
	{
		p.SetState(161)
		p.Match(SQLParserCB)
	}

	return localctx
}

// INotNullTransformContext is an interface to support dynamic dispatch.
type INotNullTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotNullTransformContext differentiates from other interfaces.
	IsNotNullTransformContext()
}

type NotNullTransformContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotNullTransformContext() *NotNullTransformContext {
	var p = new(NotNullTransformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_notNullTransform
	return p
}

func (*NotNullTransformContext) IsNotNullTransformContext() {}

func NewNotNullTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotNullTransformContext {
	var p = new(NotNullTransformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_notNullTransform

	return p
}

func (s *NotNullTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *NotNullTransformContext) AllNotNullParam() []INotNullParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INotNullParamContext)(nil)).Elem())
	var tst = make([]INotNullParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INotNullParamContext)
		}
	}

	return tst
}

func (s *NotNullTransformContext) NotNullParam(i int) INotNullParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotNullParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INotNullParamContext)
}

func (s *NotNullTransformContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *NotNullTransformContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *NotNullTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotNullTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotNullTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterNotNullTransform(s)
	}
}

func (s *NotNullTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitNotNullTransform(s)
	}
}

func (p *SQLParser) NotNullTransform() (localctx INotNullTransformContext) {
	localctx = NewNotNullTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SQLParserRULE_notNullTransform)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.NotNullParam()
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(164)
				p.Match(SQLParserCOMMA)
			}
			{
				p.SetState(165)
				p.NotNullParam()
			}

		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserCOMMA {
		{
			p.SetState(171)
			p.Match(SQLParserCOMMA)
		}

	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) ID() antlr.TerminalNode {
	return s.GetToken(SQLParserID, 0)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *SQLParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SQLParserRULE_key)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(SQLParserID)
	}

	return localctx
}

// INotNullParamContext is an interface to support dynamic dispatch.
type INotNullParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotNullParamContext differentiates from other interfaces.
	IsNotNullParamContext()
}

type NotNullParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotNullParamContext() *NotNullParamContext {
	var p = new(NotNullParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_notNullParam
	return p
}

func (*NotNullParamContext) IsNotNullParamContext() {}

func NewNotNullParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotNullParamContext {
	var p = new(NotNullParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_notNullParam

	return p
}

func (s *NotNullParamContext) GetParser() antlr.Parser { return s.parser }

func (s *NotNullParamContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SQLParserID)
}

func (s *NotNullParamContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserID, i)
}

func (s *NotNullParamContext) DOT() antlr.TerminalNode {
	return s.GetToken(SQLParserDOT, 0)
}

func (s *NotNullParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotNullParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotNullParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterNotNullParam(s)
	}
}

func (s *NotNullParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitNotNullParam(s)
	}
}

func (p *SQLParser) NotNullParam() (localctx INotNullParamContext) {
	localctx = NewNotNullParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SQLParserRULE_notNullParam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.Match(SQLParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.Match(SQLParserID)
		}
		{
			p.SetState(178)
			p.Match(SQLParserDOT)
		}
		{
			p.SetState(179)
			p.Match(SQLParserID)
		}

	}

	return localctx
}

// IQueryNameContext is an interface to support dynamic dispatch.
type IQueryNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryNameContext differentiates from other interfaces.
	IsQueryNameContext()
}

type QueryNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryNameContext() *QueryNameContext {
	var p = new(QueryNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_queryName
	return p
}

func (*QueryNameContext) IsQueryNameContext() {}

func NewQueryNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryNameContext {
	var p = new(QueryNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_queryName

	return p
}

func (s *QueryNameContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryNameContext) ID() antlr.TerminalNode {
	return s.GetToken(SQLParserID, 0)
}

func (s *QueryNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterQueryName(s)
	}
}

func (s *QueryNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitQueryName(s)
	}
}

func (p *SQLParser) QueryName() (localctx IQueryNameContext) {
	localctx = NewQueryNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SQLParserRULE_queryName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(SQLParserID)
	}

	return localctx
}

// IParamNameContext is an interface to support dynamic dispatch.
type IParamNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamNameContext differentiates from other interfaces.
	IsParamNameContext()
}

type ParamNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamNameContext() *ParamNameContext {
	var p = new(ParamNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_paramName
	return p
}

func (*ParamNameContext) IsParamNameContext() {}

func NewParamNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamNameContext {
	var p = new(ParamNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_paramName

	return p
}

func (s *ParamNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamNameContext) ID() antlr.TerminalNode {
	return s.GetToken(SQLParserID, 0)
}

func (s *ParamNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterParamName(s)
	}
}

func (s *ParamNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitParamName(s)
	}
}

func (p *SQLParser) ParamName() (localctx IParamNameContext) {
	localctx = NewParamNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SQLParserRULE_paramName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(SQLParserID)
	}

	return localctx
}

// ITemplateNameContext is an interface to support dynamic dispatch.
type ITemplateNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateNameContext differentiates from other interfaces.
	IsTemplateNameContext()
}

type TemplateNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateNameContext() *TemplateNameContext {
	var p = new(TemplateNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_templateName
	return p
}

func (*TemplateNameContext) IsTemplateNameContext() {}

func NewTemplateNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateNameContext {
	var p = new(TemplateNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_templateName

	return p
}

func (s *TemplateNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateNameContext) ID() antlr.TerminalNode {
	return s.GetToken(SQLParserID, 0)
}

func (s *TemplateNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.EnterTemplateName(s)
	}
}

func (s *TemplateNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLParserListener); ok {
		listenerT.ExitTemplateName(s)
	}
}

func (p *SQLParser) TemplateName() (localctx ITemplateNameContext) {
	localctx = NewTemplateNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SQLParserRULE_templateName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(SQLParserID)
	}

	return localctx
}
