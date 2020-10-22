// Generated from /home/ioj/projects/sqlty/compiler/parser/SQLParser.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SQLParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		ID=1, OPEN_COMMENT=2, WORD=3, EOF_STATEMENT=4, WSL=5, STRING=6, PARAM_MARK=7, 
		LINE_COMMENT=8, WS=9, SPREAD=10, NAME_TAG=11, TYPE_TAG=12, PARAM_STRUCT_NAME_TAG=13, 
		ONE_TAG=14, MANY_TAG=15, EXEC_TAG=16, NOT_NULL_PARAMS_TAG=17, RETURN_VALUE_NAME_TAG=18, 
		OB=19, CB=20, DOT=21, COMMA=22, ANY=23, CLOSE_COMMENT=24, CAST=25;
	public static final int
		RULE_input = 0, RULE_query = 1, RULE_queryDef = 2, RULE_statement = 3, 
		RULE_lineComment = 4, RULE_statementBody = 5, RULE_word = 6, RULE_param = 7, 
		RULE_paramId = 8, RULE_returnValueNameId = 9, RULE_paramStructNameId = 10, 
		RULE_nameTag = 11, RULE_paramTag = 12, RULE_paramStructNameTag = 13, RULE_modeTag = 14, 
		RULE_notNullParamsTag = 15, RULE_returnValueName = 16, RULE_anyTag = 17, 
		RULE_transformRule = 18, RULE_spreadTransform = 19, RULE_structTransform = 20, 
		RULE_structSpreadTransform = 21, RULE_notNullTransform = 22, RULE_key = 23, 
		RULE_notNullParam = 24, RULE_queryName = 25, RULE_paramName = 26;
	private static String[] makeRuleNames() {
		return new String[] {
			"input", "query", "queryDef", "statement", "lineComment", "statementBody", 
			"word", "param", "paramId", "returnValueNameId", "paramStructNameId", 
			"nameTag", "paramTag", "paramStructNameTag", "modeTag", "notNullParamsTag", 
			"returnValueName", "anyTag", "transformRule", "spreadTransform", "structTransform", 
			"structSpreadTransform", "notNullTransform", "key", "notNullParam", "queryName", 
			"paramName"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'/*'", null, "';'", null, null, "':'", null, null, "'...'", 
			"'@name'", "'@param'", "'@paramStructName'", "'@one'", "'@many'", "'@exec'", 
			"'@notNullParams'", "'@returnValueName'", "'('", "')'", "'.'", "','", 
			null, "'*/'", "'::'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "ID", "OPEN_COMMENT", "WORD", "EOF_STATEMENT", "WSL", "STRING", 
			"PARAM_MARK", "LINE_COMMENT", "WS", "SPREAD", "NAME_TAG", "TYPE_TAG", 
			"PARAM_STRUCT_NAME_TAG", "ONE_TAG", "MANY_TAG", "EXEC_TAG", "NOT_NULL_PARAMS_TAG", 
			"RETURN_VALUE_NAME_TAG", "OB", "CB", "DOT", "COMMA", "ANY", "CLOSE_COMMENT", 
			"CAST"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "SQLParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SQLParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class InputContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(SQLParser.EOF, 0); }
		public List<QueryContext> query() {
			return getRuleContexts(QueryContext.class);
		}
		public QueryContext query(int i) {
			return getRuleContext(QueryContext.class,i);
		}
		public InputContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_input; }
	}

	public final InputContext input() throws RecognitionException {
		InputContext _localctx = new InputContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_input);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(55); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(54);
				query();
				}
				}
				setState(57); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==OPEN_COMMENT );
			setState(59);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class QueryContext extends ParserRuleContext {
		public QueryDefContext queryDef() {
			return getRuleContext(QueryDefContext.class,0);
		}
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public QueryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_query; }
	}

	public final QueryContext query() throws RecognitionException {
		QueryContext _localctx = new QueryContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_query);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(61);
			queryDef();
			setState(62);
			statement();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class QueryDefContext extends ParserRuleContext {
		public TerminalNode OPEN_COMMENT() { return getToken(SQLParser.OPEN_COMMENT, 0); }
		public TerminalNode CLOSE_COMMENT() { return getToken(SQLParser.CLOSE_COMMENT, 0); }
		public List<AnyTagContext> anyTag() {
			return getRuleContexts(AnyTagContext.class);
		}
		public AnyTagContext anyTag(int i) {
			return getRuleContext(AnyTagContext.class,i);
		}
		public QueryDefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_queryDef; }
	}

	public final QueryDefContext queryDef() throws RecognitionException {
		QueryDefContext _localctx = new QueryDefContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_queryDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(64);
			match(OPEN_COMMENT);
			setState(68);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NAME_TAG) | (1L << TYPE_TAG) | (1L << PARAM_STRUCT_NAME_TAG) | (1L << ONE_TAG) | (1L << MANY_TAG) | (1L << EXEC_TAG) | (1L << NOT_NULL_PARAMS_TAG) | (1L << RETURN_VALUE_NAME_TAG))) != 0)) {
				{
				{
				setState(65);
				anyTag();
				}
				}
				setState(70);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(71);
			match(CLOSE_COMMENT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public StatementBodyContext statementBody() {
			return getRuleContext(StatementBodyContext.class,0);
		}
		public TerminalNode EOF_STATEMENT() { return getToken(SQLParser.EOF_STATEMENT, 0); }
		public LineCommentContext lineComment() {
			return getRuleContext(LineCommentContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(74);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LINE_COMMENT) {
				{
				setState(73);
				lineComment();
				}
			}

			setState(76);
			statementBody();
			setState(77);
			match(EOF_STATEMENT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LineCommentContext extends ParserRuleContext {
		public List<TerminalNode> LINE_COMMENT() { return getTokens(SQLParser.LINE_COMMENT); }
		public TerminalNode LINE_COMMENT(int i) {
			return getToken(SQLParser.LINE_COMMENT, i);
		}
		public LineCommentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lineComment; }
	}

	public final LineCommentContext lineComment() throws RecognitionException {
		LineCommentContext _localctx = new LineCommentContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_lineComment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(79);
			match(LINE_COMMENT);
			setState(83);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==LINE_COMMENT) {
				{
				{
				setState(80);
				match(LINE_COMMENT);
				}
				}
				setState(85);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementBodyContext extends ParserRuleContext {
		public List<WordContext> word() {
			return getRuleContexts(WordContext.class);
		}
		public WordContext word(int i) {
			return getRuleContext(WordContext.class,i);
		}
		public List<ParamContext> param() {
			return getRuleContexts(ParamContext.class);
		}
		public ParamContext param(int i) {
			return getRuleContext(ParamContext.class,i);
		}
		public StatementBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statementBody; }
	}

	public final StatementBodyContext statementBody() throws RecognitionException {
		StatementBodyContext _localctx = new StatementBodyContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_statementBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(86);
			word();
			setState(91);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ID) | (1L << WORD) | (1L << STRING) | (1L << PARAM_MARK))) != 0)) {
				{
				setState(89);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case PARAM_MARK:
					{
					setState(87);
					param();
					}
					break;
				case ID:
				case WORD:
				case STRING:
					{
					setState(88);
					word();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(93);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class WordContext extends ParserRuleContext {
		public TerminalNode WORD() { return getToken(SQLParser.WORD, 0); }
		public TerminalNode ID() { return getToken(SQLParser.ID, 0); }
		public TerminalNode STRING() { return getToken(SQLParser.STRING, 0); }
		public WordContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_word; }
	}

	public final WordContext word() throws RecognitionException {
		WordContext _localctx = new WordContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_word);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ID) | (1L << WORD) | (1L << STRING))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParamContext extends ParserRuleContext {
		public TerminalNode PARAM_MARK() { return getToken(SQLParser.PARAM_MARK, 0); }
		public ParamIdContext paramId() {
			return getRuleContext(ParamIdContext.class,0);
		}
		public ParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_param; }
	}

	public final ParamContext param() throws RecognitionException {
		ParamContext _localctx = new ParamContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(96);
			match(PARAM_MARK);
			setState(97);
			paramId();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParamIdContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SQLParser.ID, 0); }
		public ParamIdContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramId; }
	}

	public final ParamIdContext paramId() throws RecognitionException {
		ParamIdContext _localctx = new ParamIdContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_paramId);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(99);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ReturnValueNameIdContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SQLParser.ID, 0); }
		public ReturnValueNameIdContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnValueNameId; }
	}

	public final ReturnValueNameIdContext returnValueNameId() throws RecognitionException {
		ReturnValueNameIdContext _localctx = new ReturnValueNameIdContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_returnValueNameId);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(101);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParamStructNameIdContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SQLParser.ID, 0); }
		public ParamStructNameIdContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramStructNameId; }
	}

	public final ParamStructNameIdContext paramStructNameId() throws RecognitionException {
		ParamStructNameIdContext _localctx = new ParamStructNameIdContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_paramStructNameId);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(103);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NameTagContext extends ParserRuleContext {
		public TerminalNode NAME_TAG() { return getToken(SQLParser.NAME_TAG, 0); }
		public QueryNameContext queryName() {
			return getRuleContext(QueryNameContext.class,0);
		}
		public NameTagContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nameTag; }
	}

	public final NameTagContext nameTag() throws RecognitionException {
		NameTagContext _localctx = new NameTagContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_nameTag);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(105);
			match(NAME_TAG);
			setState(106);
			queryName();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParamTagContext extends ParserRuleContext {
		public TerminalNode TYPE_TAG() { return getToken(SQLParser.TYPE_TAG, 0); }
		public ParamNameContext paramName() {
			return getRuleContext(ParamNameContext.class,0);
		}
		public TransformRuleContext transformRule() {
			return getRuleContext(TransformRuleContext.class,0);
		}
		public ParamTagContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramTag; }
	}

	public final ParamTagContext paramTag() throws RecognitionException {
		ParamTagContext _localctx = new ParamTagContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_paramTag);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			match(TYPE_TAG);
			setState(109);
			paramName();
			setState(110);
			transformRule();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParamStructNameTagContext extends ParserRuleContext {
		public TerminalNode PARAM_STRUCT_NAME_TAG() { return getToken(SQLParser.PARAM_STRUCT_NAME_TAG, 0); }
		public ParamStructNameIdContext paramStructNameId() {
			return getRuleContext(ParamStructNameIdContext.class,0);
		}
		public ParamStructNameTagContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramStructNameTag; }
	}

	public final ParamStructNameTagContext paramStructNameTag() throws RecognitionException {
		ParamStructNameTagContext _localctx = new ParamStructNameTagContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_paramStructNameTag);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(112);
			match(PARAM_STRUCT_NAME_TAG);
			setState(113);
			paramStructNameId();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ModeTagContext extends ParserRuleContext {
		public TerminalNode ONE_TAG() { return getToken(SQLParser.ONE_TAG, 0); }
		public TerminalNode MANY_TAG() { return getToken(SQLParser.MANY_TAG, 0); }
		public TerminalNode EXEC_TAG() { return getToken(SQLParser.EXEC_TAG, 0); }
		public ModeTagContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_modeTag; }
	}

	public final ModeTagContext modeTag() throws RecognitionException {
		ModeTagContext _localctx = new ModeTagContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_modeTag);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(115);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ONE_TAG) | (1L << MANY_TAG) | (1L << EXEC_TAG))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NotNullParamsTagContext extends ParserRuleContext {
		public TerminalNode NOT_NULL_PARAMS_TAG() { return getToken(SQLParser.NOT_NULL_PARAMS_TAG, 0); }
		public NotNullTransformContext notNullTransform() {
			return getRuleContext(NotNullTransformContext.class,0);
		}
		public NotNullParamsTagContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_notNullParamsTag; }
	}

	public final NotNullParamsTagContext notNullParamsTag() throws RecognitionException {
		NotNullParamsTagContext _localctx = new NotNullParamsTagContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_notNullParamsTag);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			match(NOT_NULL_PARAMS_TAG);
			setState(118);
			notNullTransform();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ReturnValueNameContext extends ParserRuleContext {
		public TerminalNode RETURN_VALUE_NAME_TAG() { return getToken(SQLParser.RETURN_VALUE_NAME_TAG, 0); }
		public ReturnValueNameIdContext returnValueNameId() {
			return getRuleContext(ReturnValueNameIdContext.class,0);
		}
		public ReturnValueNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnValueName; }
	}

	public final ReturnValueNameContext returnValueName() throws RecognitionException {
		ReturnValueNameContext _localctx = new ReturnValueNameContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_returnValueName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(120);
			match(RETURN_VALUE_NAME_TAG);
			setState(121);
			returnValueNameId();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AnyTagContext extends ParserRuleContext {
		public NameTagContext nameTag() {
			return getRuleContext(NameTagContext.class,0);
		}
		public ParamTagContext paramTag() {
			return getRuleContext(ParamTagContext.class,0);
		}
		public ParamStructNameTagContext paramStructNameTag() {
			return getRuleContext(ParamStructNameTagContext.class,0);
		}
		public ModeTagContext modeTag() {
			return getRuleContext(ModeTagContext.class,0);
		}
		public NotNullParamsTagContext notNullParamsTag() {
			return getRuleContext(NotNullParamsTagContext.class,0);
		}
		public ReturnValueNameContext returnValueName() {
			return getRuleContext(ReturnValueNameContext.class,0);
		}
		public AnyTagContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_anyTag; }
	}

	public final AnyTagContext anyTag() throws RecognitionException {
		AnyTagContext _localctx = new AnyTagContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_anyTag);
		try {
			setState(129);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NAME_TAG:
				enterOuterAlt(_localctx, 1);
				{
				setState(123);
				nameTag();
				}
				break;
			case TYPE_TAG:
				enterOuterAlt(_localctx, 2);
				{
				setState(124);
				paramTag();
				}
				break;
			case PARAM_STRUCT_NAME_TAG:
				enterOuterAlt(_localctx, 3);
				{
				setState(125);
				paramStructNameTag();
				}
				break;
			case ONE_TAG:
			case MANY_TAG:
			case EXEC_TAG:
				enterOuterAlt(_localctx, 4);
				{
				setState(126);
				modeTag();
				}
				break;
			case NOT_NULL_PARAMS_TAG:
				enterOuterAlt(_localctx, 5);
				{
				setState(127);
				notNullParamsTag();
				}
				break;
			case RETURN_VALUE_NAME_TAG:
				enterOuterAlt(_localctx, 6);
				{
				setState(128);
				returnValueName();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TransformRuleContext extends ParserRuleContext {
		public SpreadTransformContext spreadTransform() {
			return getRuleContext(SpreadTransformContext.class,0);
		}
		public StructSpreadTransformContext structSpreadTransform() {
			return getRuleContext(StructSpreadTransformContext.class,0);
		}
		public TransformRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_transformRule; }
	}

	public final TransformRuleContext transformRule() throws RecognitionException {
		TransformRuleContext _localctx = new TransformRuleContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_transformRule);
		try {
			setState(133);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(131);
				spreadTransform();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(132);
				structSpreadTransform();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SpreadTransformContext extends ParserRuleContext {
		public TerminalNode OB() { return getToken(SQLParser.OB, 0); }
		public TerminalNode SPREAD() { return getToken(SQLParser.SPREAD, 0); }
		public TerminalNode CB() { return getToken(SQLParser.CB, 0); }
		public SpreadTransformContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_spreadTransform; }
	}

	public final SpreadTransformContext spreadTransform() throws RecognitionException {
		SpreadTransformContext _localctx = new SpreadTransformContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_spreadTransform);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(135);
			match(OB);
			setState(136);
			match(SPREAD);
			setState(137);
			match(CB);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StructTransformContext extends ParserRuleContext {
		public TerminalNode OB() { return getToken(SQLParser.OB, 0); }
		public List<KeyContext> key() {
			return getRuleContexts(KeyContext.class);
		}
		public KeyContext key(int i) {
			return getRuleContext(KeyContext.class,i);
		}
		public TerminalNode CB() { return getToken(SQLParser.CB, 0); }
		public List<TerminalNode> COMMA() { return getTokens(SQLParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(SQLParser.COMMA, i);
		}
		public StructTransformContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structTransform; }
	}

	public final StructTransformContext structTransform() throws RecognitionException {
		StructTransformContext _localctx = new StructTransformContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_structTransform);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(139);
			match(OB);
			setState(140);
			key();
			setState(145);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(141);
					match(COMMA);
					setState(142);
					key();
					}
					} 
				}
				setState(147);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			}
			setState(149);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(148);
				match(COMMA);
				}
			}

			setState(151);
			match(CB);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StructSpreadTransformContext extends ParserRuleContext {
		public TerminalNode OB() { return getToken(SQLParser.OB, 0); }
		public StructTransformContext structTransform() {
			return getRuleContext(StructTransformContext.class,0);
		}
		public TerminalNode SPREAD() { return getToken(SQLParser.SPREAD, 0); }
		public TerminalNode CB() { return getToken(SQLParser.CB, 0); }
		public StructSpreadTransformContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structSpreadTransform; }
	}

	public final StructSpreadTransformContext structSpreadTransform() throws RecognitionException {
		StructSpreadTransformContext _localctx = new StructSpreadTransformContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_structSpreadTransform);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(153);
			match(OB);
			setState(154);
			structTransform();
			setState(155);
			match(SPREAD);
			setState(156);
			match(CB);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NotNullTransformContext extends ParserRuleContext {
		public List<NotNullParamContext> notNullParam() {
			return getRuleContexts(NotNullParamContext.class);
		}
		public NotNullParamContext notNullParam(int i) {
			return getRuleContext(NotNullParamContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(SQLParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(SQLParser.COMMA, i);
		}
		public NotNullTransformContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_notNullTransform; }
	}

	public final NotNullTransformContext notNullTransform() throws RecognitionException {
		NotNullTransformContext _localctx = new NotNullTransformContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_notNullTransform);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(158);
			notNullParam();
			setState(163);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(159);
					match(COMMA);
					setState(160);
					notNullParam();
					}
					} 
				}
				setState(165);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			}
			setState(167);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(166);
				match(COMMA);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class KeyContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SQLParser.ID, 0); }
		public KeyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_key; }
	}

	public final KeyContext key() throws RecognitionException {
		KeyContext _localctx = new KeyContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_key);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(169);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NotNullParamContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(SQLParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SQLParser.ID, i);
		}
		public TerminalNode DOT() { return getToken(SQLParser.DOT, 0); }
		public NotNullParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_notNullParam; }
	}

	public final NotNullParamContext notNullParam() throws RecognitionException {
		NotNullParamContext _localctx = new NotNullParamContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_notNullParam);
		try {
			setState(175);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(171);
				match(ID);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				{
				setState(172);
				match(ID);
				setState(173);
				match(DOT);
				setState(174);
				match(ID);
				}
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class QueryNameContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SQLParser.ID, 0); }
		public QueryNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_queryName; }
	}

	public final QueryNameContext queryName() throws RecognitionException {
		QueryNameContext _localctx = new QueryNameContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_queryName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(177);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParamNameContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SQLParser.ID, 0); }
		public ParamNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramName; }
	}

	public final ParamNameContext paramName() throws RecognitionException {
		ParamNameContext _localctx = new ParamNameContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_paramName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(179);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\33\u00b8\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\3\2\6\2:\n\2\r\2\16\2;\3\2\3\2\3\3\3\3"+
		"\3\3\3\4\3\4\7\4E\n\4\f\4\16\4H\13\4\3\4\3\4\3\5\5\5M\n\5\3\5\3\5\3\5"+
		"\3\6\3\6\7\6T\n\6\f\6\16\6W\13\6\3\7\3\7\3\7\7\7\\\n\7\f\7\16\7_\13\7"+
		"\3\b\3\b\3\t\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\r\3\16\3\16\3"+
		"\16\3\16\3\17\3\17\3\17\3\20\3\20\3\21\3\21\3\21\3\22\3\22\3\22\3\23\3"+
		"\23\3\23\3\23\3\23\3\23\5\23\u0084\n\23\3\24\3\24\5\24\u0088\n\24\3\25"+
		"\3\25\3\25\3\25\3\26\3\26\3\26\3\26\7\26\u0092\n\26\f\26\16\26\u0095\13"+
		"\26\3\26\5\26\u0098\n\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\30\3\30"+
		"\3\30\7\30\u00a4\n\30\f\30\16\30\u00a7\13\30\3\30\5\30\u00aa\n\30\3\31"+
		"\3\31\3\32\3\32\3\32\3\32\5\32\u00b2\n\32\3\33\3\33\3\34\3\34\3\34\2\2"+
		"\35\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\66\2\4\5"+
		"\2\3\3\5\5\b\b\3\2\20\22\2\u00ad\29\3\2\2\2\4?\3\2\2\2\6B\3\2\2\2\bL\3"+
		"\2\2\2\nQ\3\2\2\2\fX\3\2\2\2\16`\3\2\2\2\20b\3\2\2\2\22e\3\2\2\2\24g\3"+
		"\2\2\2\26i\3\2\2\2\30k\3\2\2\2\32n\3\2\2\2\34r\3\2\2\2\36u\3\2\2\2 w\3"+
		"\2\2\2\"z\3\2\2\2$\u0083\3\2\2\2&\u0087\3\2\2\2(\u0089\3\2\2\2*\u008d"+
		"\3\2\2\2,\u009b\3\2\2\2.\u00a0\3\2\2\2\60\u00ab\3\2\2\2\62\u00b1\3\2\2"+
		"\2\64\u00b3\3\2\2\2\66\u00b5\3\2\2\28:\5\4\3\298\3\2\2\2:;\3\2\2\2;9\3"+
		"\2\2\2;<\3\2\2\2<=\3\2\2\2=>\7\2\2\3>\3\3\2\2\2?@\5\6\4\2@A\5\b\5\2A\5"+
		"\3\2\2\2BF\7\4\2\2CE\5$\23\2DC\3\2\2\2EH\3\2\2\2FD\3\2\2\2FG\3\2\2\2G"+
		"I\3\2\2\2HF\3\2\2\2IJ\7\32\2\2J\7\3\2\2\2KM\5\n\6\2LK\3\2\2\2LM\3\2\2"+
		"\2MN\3\2\2\2NO\5\f\7\2OP\7\6\2\2P\t\3\2\2\2QU\7\n\2\2RT\7\n\2\2SR\3\2"+
		"\2\2TW\3\2\2\2US\3\2\2\2UV\3\2\2\2V\13\3\2\2\2WU\3\2\2\2X]\5\16\b\2Y\\"+
		"\5\20\t\2Z\\\5\16\b\2[Y\3\2\2\2[Z\3\2\2\2\\_\3\2\2\2][\3\2\2\2]^\3\2\2"+
		"\2^\r\3\2\2\2_]\3\2\2\2`a\t\2\2\2a\17\3\2\2\2bc\7\t\2\2cd\5\22\n\2d\21"+
		"\3\2\2\2ef\7\3\2\2f\23\3\2\2\2gh\7\3\2\2h\25\3\2\2\2ij\7\3\2\2j\27\3\2"+
		"\2\2kl\7\r\2\2lm\5\64\33\2m\31\3\2\2\2no\7\16\2\2op\5\66\34\2pq\5&\24"+
		"\2q\33\3\2\2\2rs\7\17\2\2st\5\26\f\2t\35\3\2\2\2uv\t\3\2\2v\37\3\2\2\2"+
		"wx\7\23\2\2xy\5.\30\2y!\3\2\2\2z{\7\24\2\2{|\5\24\13\2|#\3\2\2\2}\u0084"+
		"\5\30\r\2~\u0084\5\32\16\2\177\u0084\5\34\17\2\u0080\u0084\5\36\20\2\u0081"+
		"\u0084\5 \21\2\u0082\u0084\5\"\22\2\u0083}\3\2\2\2\u0083~\3\2\2\2\u0083"+
		"\177\3\2\2\2\u0083\u0080\3\2\2\2\u0083\u0081\3\2\2\2\u0083\u0082\3\2\2"+
		"\2\u0084%\3\2\2\2\u0085\u0088\5(\25\2\u0086\u0088\5,\27\2\u0087\u0085"+
		"\3\2\2\2\u0087\u0086\3\2\2\2\u0088\'\3\2\2\2\u0089\u008a\7\25\2\2\u008a"+
		"\u008b\7\f\2\2\u008b\u008c\7\26\2\2\u008c)\3\2\2\2\u008d\u008e\7\25\2"+
		"\2\u008e\u0093\5\60\31\2\u008f\u0090\7\30\2\2\u0090\u0092\5\60\31\2\u0091"+
		"\u008f\3\2\2\2\u0092\u0095\3\2\2\2\u0093\u0091\3\2\2\2\u0093\u0094\3\2"+
		"\2\2\u0094\u0097\3\2\2\2\u0095\u0093\3\2\2\2\u0096\u0098\7\30\2\2\u0097"+
		"\u0096\3\2\2\2\u0097\u0098\3\2\2\2\u0098\u0099\3\2\2\2\u0099\u009a\7\26"+
		"\2\2\u009a+\3\2\2\2\u009b\u009c\7\25\2\2\u009c\u009d\5*\26\2\u009d\u009e"+
		"\7\f\2\2\u009e\u009f\7\26\2\2\u009f-\3\2\2\2\u00a0\u00a5\5\62\32\2\u00a1"+
		"\u00a2\7\30\2\2\u00a2\u00a4\5\62\32\2\u00a3\u00a1\3\2\2\2\u00a4\u00a7"+
		"\3\2\2\2\u00a5\u00a3\3\2\2\2\u00a5\u00a6\3\2\2\2\u00a6\u00a9\3\2\2\2\u00a7"+
		"\u00a5\3\2\2\2\u00a8\u00aa\7\30\2\2\u00a9\u00a8\3\2\2\2\u00a9\u00aa\3"+
		"\2\2\2\u00aa/\3\2\2\2\u00ab\u00ac\7\3\2\2\u00ac\61\3\2\2\2\u00ad\u00b2"+
		"\7\3\2\2\u00ae\u00af\7\3\2\2\u00af\u00b0\7\27\2\2\u00b0\u00b2\7\3\2\2"+
		"\u00b1\u00ad\3\2\2\2\u00b1\u00ae\3\2\2\2\u00b2\63\3\2\2\2\u00b3\u00b4"+
		"\7\3\2\2\u00b4\65\3\2\2\2\u00b5\u00b6\7\3\2\2\u00b6\67\3\2\2\2\17;FLU"+
		"[]\u0083\u0087\u0093\u0097\u00a5\u00a9\u00b1";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}