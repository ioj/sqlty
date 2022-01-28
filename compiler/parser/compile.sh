#!/bin/sh
alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.7.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -o . SQLLexer.g4
antlr4 -Dlanguage=Go -o . SQLParser.g4
