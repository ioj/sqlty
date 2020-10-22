parser grammar SQLParser;

options {
	tokenVocab = SQLLexer;
}

input: query+ EOF;

query: queryDef statement;

queryDef: OPEN_COMMENT anyTag* CLOSE_COMMENT;

statement: lineComment? statementBody EOF_STATEMENT;

lineComment: LINE_COMMENT LINE_COMMENT*;

statementBody: word (param | word)*;

word: WORD | ID | STRING;

param: PARAM_MARK paramId;

paramId: ID;

returnValueNameId: ID;

paramStructNameId: ID;

nameTag: NAME_TAG queryName;

paramTag: TYPE_TAG paramName transformRule;

paramStructNameTag: PARAM_STRUCT_NAME_TAG paramStructNameId;

modeTag: ONE_TAG | MANY_TAG | EXEC_TAG;

notNullParamsTag: NOT_NULL_PARAMS_TAG notNullTransform;

returnValueName: RETURN_VALUE_NAME_TAG returnValueNameId;

anyTag:
	nameTag
	| paramTag
	| paramStructNameTag
	| modeTag
	| notNullParamsTag
	| returnValueName;

transformRule: spreadTransform | structSpreadTransform;

spreadTransform: OB SPREAD CB;

structTransform: OB key (COMMA key)* COMMA? CB;

structSpreadTransform: OB structTransform SPREAD CB;

notNullTransform: notNullParam (COMMA notNullParam)* COMMA?;

key: ID;

notNullParam: ID | (ID DOT ID);

queryName: ID;
paramName: ID;

