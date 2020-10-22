// Generated from /home/ioj/projects/sqlty/compiler/parser/SQLLexer.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SQLLexer extends Lexer {
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
		COMMENT=1;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE", "COMMENT"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"QUOT", "ID", "OPEN_COMMENT", "SID", "WORD", "SPECIAL", "EOF_STATEMENT", 
			"WSL", "STRING", "PARAM_MARK", "CAST", "LINE_COMMENT", "CID", "WS", "SPREAD", 
			"NAME_TAG", "TYPE_TAG", "PARAM_STRUCT_NAME_TAG", "ONE_TAG", "MANY_TAG", 
			"EXEC_TAG", "NOT_NULL_PARAMS_TAG", "RETURN_VALUE_NAME_TAG", "OB", "CB", 
			"DOT", "COMMA", "ANY", "CLOSE_COMMENT"
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


	public SQLLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "SQLLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\33\u00f6\b\1\b\1"+
		"\4\2\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t"+
		"\n\4\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4"+
		"\22\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4"+
		"\31\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\3\2\3\2\3\3"+
		"\3\3\7\3C\n\3\f\3\16\3F\13\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\6"+
		"\6R\n\6\r\6\16\6S\3\7\6\7W\n\7\r\7\16\7X\3\7\3\7\3\b\3\b\3\t\6\t`\n\t"+
		"\r\t\16\ta\3\t\3\t\3\n\3\n\3\n\7\ni\n\n\f\n\16\nl\13\n\3\n\3\n\5\np\n"+
		"\n\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\7\r}\n\r\f\r\16\r\u0080"+
		"\13\r\3\r\5\r\u0083\n\r\3\r\3\r\3\16\3\16\3\16\3\16\3\17\6\17\u008c\n"+
		"\17\r\17\16\17\u008d\3\17\3\17\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24"+
		"\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26"+
		"\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\35"+
		"\6\35\u00ee\n\35\r\35\16\35\u00ef\3\36\3\36\3\36\3\36\3\36\4j\u00ef\2"+
		"\37\4\2\6\3\b\4\n\2\f\5\16\2\20\6\22\7\24\b\26\t\30\33\32\n\34\2\36\13"+
		" \f\"\r$\16&\17(\20*\21,\22.\23\60\24\62\25\64\26\66\278\30:\31<\32\4"+
		"\2\3\b\5\2C\\aac|\6\2\62;C\\aac|\t\2#(*\61>B]]_`bb}\u0080\5\2\13\f\17"+
		"\17\"\"\3\2^^\4\2\f\f\17\17\2\u00fc\2\b\3\2\2\2\2\n\3\2\2\2\2\f\3\2\2"+
		"\2\2\16\3\2\2\2\2\20\3\2\2\2\2\22\3\2\2\2\2\24\3\2\2\2\2\26\3\2\2\2\2"+
		"\30\3\2\2\2\2\32\3\2\2\2\3\34\3\2\2\2\3\36\3\2\2\2\3 \3\2\2\2\3\"\3\2"+
		"\2\2\3$\3\2\2\2\3&\3\2\2\2\3(\3\2\2\2\3*\3\2\2\2\3,\3\2\2\2\3.\3\2\2\2"+
		"\3\60\3\2\2\2\3\62\3\2\2\2\3\64\3\2\2\2\3\66\3\2\2\2\38\3\2\2\2\3:\3\2"+
		"\2\2\3<\3\2\2\2\4>\3\2\2\2\6@\3\2\2\2\bG\3\2\2\2\nL\3\2\2\2\fQ\3\2\2\2"+
		"\16V\3\2\2\2\20\\\3\2\2\2\22_\3\2\2\2\24e\3\2\2\2\26q\3\2\2\2\30s\3\2"+
		"\2\2\32x\3\2\2\2\34\u0086\3\2\2\2\36\u008b\3\2\2\2 \u0091\3\2\2\2\"\u0095"+
		"\3\2\2\2$\u009b\3\2\2\2&\u00a2\3\2\2\2(\u00b3\3\2\2\2*\u00b8\3\2\2\2,"+
		"\u00be\3\2\2\2.\u00c4\3\2\2\2\60\u00d3\3\2\2\2\62\u00e4\3\2\2\2\64\u00e6"+
		"\3\2\2\2\66\u00e8\3\2\2\28\u00ea\3\2\2\2:\u00ed\3\2\2\2<\u00f1\3\2\2\2"+
		">?\7)\2\2?\5\3\2\2\2@D\t\2\2\2AC\t\3\2\2BA\3\2\2\2CF\3\2\2\2DB\3\2\2\2"+
		"DE\3\2\2\2E\7\3\2\2\2FD\3\2\2\2GH\7\61\2\2HI\7,\2\2IJ\3\2\2\2JK\b\4\2"+
		"\2K\t\3\2\2\2LM\5\6\3\2MN\3\2\2\2NO\b\5\3\2O\13\3\2\2\2PR\t\3\2\2QP\3"+
		"\2\2\2RS\3\2\2\2SQ\3\2\2\2ST\3\2\2\2T\r\3\2\2\2UW\t\4\2\2VU\3\2\2\2WX"+
		"\3\2\2\2XV\3\2\2\2XY\3\2\2\2YZ\3\2\2\2Z[\b\7\4\2[\17\3\2\2\2\\]\7=\2\2"+
		"]\21\3\2\2\2^`\t\5\2\2_^\3\2\2\2`a\3\2\2\2a_\3\2\2\2ab\3\2\2\2bc\3\2\2"+
		"\2cd\b\t\5\2d\23\3\2\2\2eo\5\4\2\2fp\5\4\2\2gi\13\2\2\2hg\3\2\2\2il\3"+
		"\2\2\2jk\3\2\2\2jh\3\2\2\2km\3\2\2\2lj\3\2\2\2mn\n\6\2\2np\5\4\2\2of\3"+
		"\2\2\2oj\3\2\2\2p\25\3\2\2\2qr\7<\2\2r\27\3\2\2\2st\7<\2\2tu\7<\2\2uv"+
		"\3\2\2\2vw\b\f\4\2w\31\3\2\2\2xy\7/\2\2yz\7/\2\2z~\3\2\2\2{}\n\7\2\2|"+
		"{\3\2\2\2}\u0080\3\2\2\2~|\3\2\2\2~\177\3\2\2\2\177\u0082\3\2\2\2\u0080"+
		"~\3\2\2\2\u0081\u0083\7\17\2\2\u0082\u0081\3\2\2\2\u0082\u0083\3\2\2\2"+
		"\u0083\u0084\3\2\2\2\u0084\u0085\7\f\2\2\u0085\33\3\2\2\2\u0086\u0087"+
		"\5\6\3\2\u0087\u0088\3\2\2\2\u0088\u0089\b\16\3\2\u0089\35\3\2\2\2\u008a"+
		"\u008c\t\5\2\2\u008b\u008a\3\2\2\2\u008c\u008d\3\2\2\2\u008d\u008b\3\2"+
		"\2\2\u008d\u008e\3\2\2\2\u008e\u008f\3\2\2\2\u008f\u0090\b\17\5\2\u0090"+
		"\37\3\2\2\2\u0091\u0092\7\60\2\2\u0092\u0093\7\60\2\2\u0093\u0094\7\60"+
		"\2\2\u0094!\3\2\2\2\u0095\u0096\7B\2\2\u0096\u0097\7p\2\2\u0097\u0098"+
		"\7c\2\2\u0098\u0099\7o\2\2\u0099\u009a\7g\2\2\u009a#\3\2\2\2\u009b\u009c"+
		"\7B\2\2\u009c\u009d\7r\2\2\u009d\u009e\7c\2\2\u009e\u009f\7t\2\2\u009f"+
		"\u00a0\7c\2\2\u00a0\u00a1\7o\2\2\u00a1%\3\2\2\2\u00a2\u00a3\7B\2\2\u00a3"+
		"\u00a4\7r\2\2\u00a4\u00a5\7c\2\2\u00a5\u00a6\7t\2\2\u00a6\u00a7\7c\2\2"+
		"\u00a7\u00a8\7o\2\2\u00a8\u00a9\7U\2\2\u00a9\u00aa\7v\2\2\u00aa\u00ab"+
		"\7t\2\2\u00ab\u00ac\7w\2\2\u00ac\u00ad\7e\2\2\u00ad\u00ae\7v\2\2\u00ae"+
		"\u00af\7P\2\2\u00af\u00b0\7c\2\2\u00b0\u00b1\7o\2\2\u00b1\u00b2\7g\2\2"+
		"\u00b2\'\3\2\2\2\u00b3\u00b4\7B\2\2\u00b4\u00b5\7q\2\2\u00b5\u00b6\7p"+
		"\2\2\u00b6\u00b7\7g\2\2\u00b7)\3\2\2\2\u00b8\u00b9\7B\2\2\u00b9\u00ba"+
		"\7o\2\2\u00ba\u00bb\7c\2\2\u00bb\u00bc\7p\2\2\u00bc\u00bd\7{\2\2\u00bd"+
		"+\3\2\2\2\u00be\u00bf\7B\2\2\u00bf\u00c0\7g\2\2\u00c0\u00c1\7z\2\2\u00c1"+
		"\u00c2\7g\2\2\u00c2\u00c3\7e\2\2\u00c3-\3\2\2\2\u00c4\u00c5\7B\2\2\u00c5"+
		"\u00c6\7p\2\2\u00c6\u00c7\7q\2\2\u00c7\u00c8\7v\2\2\u00c8\u00c9\7P\2\2"+
		"\u00c9\u00ca\7w\2\2\u00ca\u00cb\7n\2\2\u00cb\u00cc\7n\2\2\u00cc\u00cd"+
		"\7R\2\2\u00cd\u00ce\7c\2\2\u00ce\u00cf\7t\2\2\u00cf\u00d0\7c\2\2\u00d0"+
		"\u00d1\7o\2\2\u00d1\u00d2\7u\2\2\u00d2/\3\2\2\2\u00d3\u00d4\7B\2\2\u00d4"+
		"\u00d5\7t\2\2\u00d5\u00d6\7g\2\2\u00d6\u00d7\7v\2\2\u00d7\u00d8\7w\2\2"+
		"\u00d8\u00d9\7t\2\2\u00d9\u00da\7p\2\2\u00da\u00db\7X\2\2\u00db\u00dc"+
		"\7c\2\2\u00dc\u00dd\7n\2\2\u00dd\u00de\7w\2\2\u00de\u00df\7g\2\2\u00df"+
		"\u00e0\7P\2\2\u00e0\u00e1\7c\2\2\u00e1\u00e2\7o\2\2\u00e2\u00e3\7g\2\2"+
		"\u00e3\61\3\2\2\2\u00e4\u00e5\7*\2\2\u00e5\63\3\2\2\2\u00e6\u00e7\7+\2"+
		"\2\u00e7\65\3\2\2\2\u00e8\u00e9\7\60\2\2\u00e9\67\3\2\2\2\u00ea\u00eb"+
		"\7.\2\2\u00eb9\3\2\2\2\u00ec\u00ee\13\2\2\2\u00ed\u00ec\3\2\2\2\u00ee"+
		"\u00ef\3\2\2\2\u00ef\u00f0\3\2\2\2\u00ef\u00ed\3\2\2\2\u00f0;\3\2\2\2"+
		"\u00f1\u00f2\7,\2\2\u00f2\u00f3\7\61\2\2\u00f3\u00f4\3\2\2\2\u00f4\u00f5"+
		"\b\36\6\2\u00f5=\3\2\2\2\16\2\3DSXajo~\u0082\u008d\u00ef\7\4\3\2\t\3\2"+
		"\t\5\2\b\2\2\4\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}