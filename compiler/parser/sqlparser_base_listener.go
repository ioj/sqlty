// Code generated from SQLParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // SQLParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSQLParserListener is a complete listener for a parse tree produced by SQLParser.
type BaseSQLParserListener struct{}

var _ SQLParserListener = &BaseSQLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSQLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSQLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSQLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSQLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInput is called when production input is entered.
func (s *BaseSQLParserListener) EnterInput(ctx *InputContext) {}

// ExitInput is called when production input is exited.
func (s *BaseSQLParserListener) ExitInput(ctx *InputContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseSQLParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseSQLParserListener) ExitQuery(ctx *QueryContext) {}

// EnterQueryDef is called when production queryDef is entered.
func (s *BaseSQLParserListener) EnterQueryDef(ctx *QueryDefContext) {}

// ExitQueryDef is called when production queryDef is exited.
func (s *BaseSQLParserListener) ExitQueryDef(ctx *QueryDefContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSQLParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSQLParserListener) ExitStatement(ctx *StatementContext) {}

// EnterLineComment is called when production lineComment is entered.
func (s *BaseSQLParserListener) EnterLineComment(ctx *LineCommentContext) {}

// ExitLineComment is called when production lineComment is exited.
func (s *BaseSQLParserListener) ExitLineComment(ctx *LineCommentContext) {}

// EnterStatementBody is called when production statementBody is entered.
func (s *BaseSQLParserListener) EnterStatementBody(ctx *StatementBodyContext) {}

// ExitStatementBody is called when production statementBody is exited.
func (s *BaseSQLParserListener) ExitStatementBody(ctx *StatementBodyContext) {}

// EnterWord is called when production word is entered.
func (s *BaseSQLParserListener) EnterWord(ctx *WordContext) {}

// ExitWord is called when production word is exited.
func (s *BaseSQLParserListener) ExitWord(ctx *WordContext) {}

// EnterParam is called when production param is entered.
func (s *BaseSQLParserListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseSQLParserListener) ExitParam(ctx *ParamContext) {}

// EnterParamId is called when production paramId is entered.
func (s *BaseSQLParserListener) EnterParamId(ctx *ParamIdContext) {}

// ExitParamId is called when production paramId is exited.
func (s *BaseSQLParserListener) ExitParamId(ctx *ParamIdContext) {}

// EnterReturnValueNameId is called when production returnValueNameId is entered.
func (s *BaseSQLParserListener) EnterReturnValueNameId(ctx *ReturnValueNameIdContext) {}

// ExitReturnValueNameId is called when production returnValueNameId is exited.
func (s *BaseSQLParserListener) ExitReturnValueNameId(ctx *ReturnValueNameIdContext) {}

// EnterParamStructNameId is called when production paramStructNameId is entered.
func (s *BaseSQLParserListener) EnterParamStructNameId(ctx *ParamStructNameIdContext) {}

// ExitParamStructNameId is called when production paramStructNameId is exited.
func (s *BaseSQLParserListener) ExitParamStructNameId(ctx *ParamStructNameIdContext) {}

// EnterNameTag is called when production nameTag is entered.
func (s *BaseSQLParserListener) EnterNameTag(ctx *NameTagContext) {}

// ExitNameTag is called when production nameTag is exited.
func (s *BaseSQLParserListener) ExitNameTag(ctx *NameTagContext) {}

// EnterParamTag is called when production paramTag is entered.
func (s *BaseSQLParserListener) EnterParamTag(ctx *ParamTagContext) {}

// ExitParamTag is called when production paramTag is exited.
func (s *BaseSQLParserListener) ExitParamTag(ctx *ParamTagContext) {}

// EnterParamStructNameTag is called when production paramStructNameTag is entered.
func (s *BaseSQLParserListener) EnterParamStructNameTag(ctx *ParamStructNameTagContext) {}

// ExitParamStructNameTag is called when production paramStructNameTag is exited.
func (s *BaseSQLParserListener) ExitParamStructNameTag(ctx *ParamStructNameTagContext) {}

// EnterModeTag is called when production modeTag is entered.
func (s *BaseSQLParserListener) EnterModeTag(ctx *ModeTagContext) {}

// ExitModeTag is called when production modeTag is exited.
func (s *BaseSQLParserListener) ExitModeTag(ctx *ModeTagContext) {}

// EnterNotNullParamsTag is called when production notNullParamsTag is entered.
func (s *BaseSQLParserListener) EnterNotNullParamsTag(ctx *NotNullParamsTagContext) {}

// ExitNotNullParamsTag is called when production notNullParamsTag is exited.
func (s *BaseSQLParserListener) ExitNotNullParamsTag(ctx *NotNullParamsTagContext) {}

// EnterReturnValueName is called when production returnValueName is entered.
func (s *BaseSQLParserListener) EnterReturnValueName(ctx *ReturnValueNameContext) {}

// ExitReturnValueName is called when production returnValueName is exited.
func (s *BaseSQLParserListener) ExitReturnValueName(ctx *ReturnValueNameContext) {}

// EnterAnyTag is called when production anyTag is entered.
func (s *BaseSQLParserListener) EnterAnyTag(ctx *AnyTagContext) {}

// ExitAnyTag is called when production anyTag is exited.
func (s *BaseSQLParserListener) ExitAnyTag(ctx *AnyTagContext) {}

// EnterTransformRule is called when production transformRule is entered.
func (s *BaseSQLParserListener) EnterTransformRule(ctx *TransformRuleContext) {}

// ExitTransformRule is called when production transformRule is exited.
func (s *BaseSQLParserListener) ExitTransformRule(ctx *TransformRuleContext) {}

// EnterSpreadTransform is called when production spreadTransform is entered.
func (s *BaseSQLParserListener) EnterSpreadTransform(ctx *SpreadTransformContext) {}

// ExitSpreadTransform is called when production spreadTransform is exited.
func (s *BaseSQLParserListener) ExitSpreadTransform(ctx *SpreadTransformContext) {}

// EnterStructTransform is called when production structTransform is entered.
func (s *BaseSQLParserListener) EnterStructTransform(ctx *StructTransformContext) {}

// ExitStructTransform is called when production structTransform is exited.
func (s *BaseSQLParserListener) ExitStructTransform(ctx *StructTransformContext) {}

// EnterStructSpreadTransform is called when production structSpreadTransform is entered.
func (s *BaseSQLParserListener) EnterStructSpreadTransform(ctx *StructSpreadTransformContext) {}

// ExitStructSpreadTransform is called when production structSpreadTransform is exited.
func (s *BaseSQLParserListener) ExitStructSpreadTransform(ctx *StructSpreadTransformContext) {}

// EnterNotNullTransform is called when production notNullTransform is entered.
func (s *BaseSQLParserListener) EnterNotNullTransform(ctx *NotNullTransformContext) {}

// ExitNotNullTransform is called when production notNullTransform is exited.
func (s *BaseSQLParserListener) ExitNotNullTransform(ctx *NotNullTransformContext) {}

// EnterKey is called when production key is entered.
func (s *BaseSQLParserListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseSQLParserListener) ExitKey(ctx *KeyContext) {}

// EnterNotNullParam is called when production notNullParam is entered.
func (s *BaseSQLParserListener) EnterNotNullParam(ctx *NotNullParamContext) {}

// ExitNotNullParam is called when production notNullParam is exited.
func (s *BaseSQLParserListener) ExitNotNullParam(ctx *NotNullParamContext) {}

// EnterQueryName is called when production queryName is entered.
func (s *BaseSQLParserListener) EnterQueryName(ctx *QueryNameContext) {}

// ExitQueryName is called when production queryName is exited.
func (s *BaseSQLParserListener) ExitQueryName(ctx *QueryNameContext) {}

// EnterParamName is called when production paramName is entered.
func (s *BaseSQLParserListener) EnterParamName(ctx *ParamNameContext) {}

// ExitParamName is called when production paramName is exited.
func (s *BaseSQLParserListener) ExitParamName(ctx *ParamNameContext) {}
