lexer grammar SQLLexer;

tokens {
	ID
}

fragment QUOT: '\'';
fragment ID: [a-zA-Z_][a-zA-Z_0-9]*;

OPEN_COMMENT: '/*' -> mode(COMMENT);
SID: ID -> type(ID);
WORD: [a-zA-Z_0-9]+;
SPECIAL: [\-+*/<>=~!@#%^&|`?$(){},.[\]"]+ -> type(WORD);
EOF_STATEMENT: ';';
WSL: [ \t\r\n]+ -> skip;
// parse strings and recognize escaped quotes
STRING: QUOT (QUOT | .*? ~([\\]) QUOT);
PARAM_MARK: ':';
CAST: '::' -> type(WORD);
LINE_COMMENT: '--' ~[\r\n]* '\r'? '\n';
mode COMMENT;
CID: ID -> type(ID);
WS: [ \t\r\n]+ -> skip;
SPREAD: '...';
NAME_TAG: '@name';
TYPE_TAG: '@param';
PARAM_STRUCT_NAME_TAG: '@paramStructName';
ONE_TAG: '@one';
MANY_TAG: '@many';
EXEC_TAG: '@exec';
NOT_NULL_PARAMS_TAG: '@notNullParams';
RETURN_VALUE_NAME_TAG: '@returnValueName';
TEMPLATE_TAG: '@template';
OB: '(';
CB: ')';
DOT: '.';
COMMA: ',';
ANY: .+?;
CLOSE_COMMENT: '*/' -> mode(DEFAULT_MODE);