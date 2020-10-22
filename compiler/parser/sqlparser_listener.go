// Code generated from SQLParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // SQLParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SQLParserListener is a complete listener for a parse tree produced by SQLParser.
type SQLParserListener interface {
	antlr.ParseTreeListener

	// EnterInput is called when entering the input production.
	EnterInput(c *InputContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterQueryDef is called when entering the queryDef production.
	EnterQueryDef(c *QueryDefContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterLineComment is called when entering the lineComment production.
	EnterLineComment(c *LineCommentContext)

	// EnterStatementBody is called when entering the statementBody production.
	EnterStatementBody(c *StatementBodyContext)

	// EnterWord is called when entering the word production.
	EnterWord(c *WordContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterParamId is called when entering the paramId production.
	EnterParamId(c *ParamIdContext)

	// EnterReturnValueNameId is called when entering the returnValueNameId production.
	EnterReturnValueNameId(c *ReturnValueNameIdContext)

	// EnterParamStructNameId is called when entering the paramStructNameId production.
	EnterParamStructNameId(c *ParamStructNameIdContext)

	// EnterNameTag is called when entering the nameTag production.
	EnterNameTag(c *NameTagContext)

	// EnterParamTag is called when entering the paramTag production.
	EnterParamTag(c *ParamTagContext)

	// EnterParamStructNameTag is called when entering the paramStructNameTag production.
	EnterParamStructNameTag(c *ParamStructNameTagContext)

	// EnterModeTag is called when entering the modeTag production.
	EnterModeTag(c *ModeTagContext)

	// EnterNotNullParamsTag is called when entering the notNullParamsTag production.
	EnterNotNullParamsTag(c *NotNullParamsTagContext)

	// EnterReturnValueName is called when entering the returnValueName production.
	EnterReturnValueName(c *ReturnValueNameContext)

	// EnterAnyTag is called when entering the anyTag production.
	EnterAnyTag(c *AnyTagContext)

	// EnterTransformRule is called when entering the transformRule production.
	EnterTransformRule(c *TransformRuleContext)

	// EnterSpreadTransform is called when entering the spreadTransform production.
	EnterSpreadTransform(c *SpreadTransformContext)

	// EnterStructTransform is called when entering the structTransform production.
	EnterStructTransform(c *StructTransformContext)

	// EnterStructSpreadTransform is called when entering the structSpreadTransform production.
	EnterStructSpreadTransform(c *StructSpreadTransformContext)

	// EnterNotNullTransform is called when entering the notNullTransform production.
	EnterNotNullTransform(c *NotNullTransformContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterNotNullParam is called when entering the notNullParam production.
	EnterNotNullParam(c *NotNullParamContext)

	// EnterQueryName is called when entering the queryName production.
	EnterQueryName(c *QueryNameContext)

	// EnterParamName is called when entering the paramName production.
	EnterParamName(c *ParamNameContext)

	// ExitInput is called when exiting the input production.
	ExitInput(c *InputContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitQueryDef is called when exiting the queryDef production.
	ExitQueryDef(c *QueryDefContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitLineComment is called when exiting the lineComment production.
	ExitLineComment(c *LineCommentContext)

	// ExitStatementBody is called when exiting the statementBody production.
	ExitStatementBody(c *StatementBodyContext)

	// ExitWord is called when exiting the word production.
	ExitWord(c *WordContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitParamId is called when exiting the paramId production.
	ExitParamId(c *ParamIdContext)

	// ExitReturnValueNameId is called when exiting the returnValueNameId production.
	ExitReturnValueNameId(c *ReturnValueNameIdContext)

	// ExitParamStructNameId is called when exiting the paramStructNameId production.
	ExitParamStructNameId(c *ParamStructNameIdContext)

	// ExitNameTag is called when exiting the nameTag production.
	ExitNameTag(c *NameTagContext)

	// ExitParamTag is called when exiting the paramTag production.
	ExitParamTag(c *ParamTagContext)

	// ExitParamStructNameTag is called when exiting the paramStructNameTag production.
	ExitParamStructNameTag(c *ParamStructNameTagContext)

	// ExitModeTag is called when exiting the modeTag production.
	ExitModeTag(c *ModeTagContext)

	// ExitNotNullParamsTag is called when exiting the notNullParamsTag production.
	ExitNotNullParamsTag(c *NotNullParamsTagContext)

	// ExitReturnValueName is called when exiting the returnValueName production.
	ExitReturnValueName(c *ReturnValueNameContext)

	// ExitAnyTag is called when exiting the anyTag production.
	ExitAnyTag(c *AnyTagContext)

	// ExitTransformRule is called when exiting the transformRule production.
	ExitTransformRule(c *TransformRuleContext)

	// ExitSpreadTransform is called when exiting the spreadTransform production.
	ExitSpreadTransform(c *SpreadTransformContext)

	// ExitStructTransform is called when exiting the structTransform production.
	ExitStructTransform(c *StructTransformContext)

	// ExitStructSpreadTransform is called when exiting the structSpreadTransform production.
	ExitStructSpreadTransform(c *StructSpreadTransformContext)

	// ExitNotNullTransform is called when exiting the notNullTransform production.
	ExitNotNullTransform(c *NotNullTransformContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitNotNullParam is called when exiting the notNullParam production.
	ExitNotNullParam(c *NotNullParamContext)

	// ExitQueryName is called when exiting the queryName production.
	ExitQueryName(c *QueryNameContext)

	// ExitParamName is called when exiting the paramName production.
	ExitParamName(c *ParamNameContext)
}
