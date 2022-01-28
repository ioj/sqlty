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
		TEMPLATE_TAG=19, OB=20, CB=21, DOT=22, COMMA=23, ANY=24, CLOSE_COMMENT=25, 
		CAST=26;
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
			"EXEC_TAG", "NOT_NULL_PARAMS_TAG", "RETURN_VALUE_NAME_TAG", "TEMPLATE_TAG", 
			"OB", "CB", "DOT", "COMMA", "ANY", "CLOSE_COMMENT"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'/*'", null, "';'", null, null, "':'", null, null, "'...'", 
			"'@name'", "'@param'", "'@paramStructName'", "'@one'", "'@many'", "'@exec'", 
			"'@notNullParams'", "'@returnValueName'", "'@template'", "'('", "')'", 
			"'.'", "','", null, "'*/'", "'::'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "ID", "OPEN_COMMENT", "WORD", "EOF_STATEMENT", "WSL", "STRING", 
			"PARAM_MARK", "LINE_COMMENT", "WS", "SPREAD", "NAME_TAG", "TYPE_TAG", 
			"PARAM_STRUCT_NAME_TAG", "ONE_TAG", "MANY_TAG", "EXEC_TAG", "NOT_NULL_PARAMS_TAG", 
			"RETURN_VALUE_NAME_TAG", "TEMPLATE_TAG", "OB", "CB", "DOT", "COMMA", 
			"ANY", "CLOSE_COMMENT", "CAST"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\34\u0102\b\1\b\1"+
		"\4\2\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t"+
		"\n\4\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4"+
		"\22\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4"+
		"\31\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\3"+
		"\2\3\2\3\3\3\3\7\3E\n\3\f\3\16\3H\13\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5"+
		"\3\5\3\6\6\6T\n\6\r\6\16\6U\3\7\6\7Y\n\7\r\7\16\7Z\3\7\3\7\3\b\3\b\3\t"+
		"\6\tb\n\t\r\t\16\tc\3\t\3\t\3\n\3\n\3\n\7\nk\n\n\f\n\16\nn\13\n\3\n\3"+
		"\n\5\nr\n\n\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\7\r\177\n\r"+
		"\f\r\16\r\u0082\13\r\3\r\5\r\u0085\n\r\3\r\3\r\3\16\3\16\3\16\3\16\3\17"+
		"\6\17\u008e\n\17\r\17\16\17\u008f\3\17\3\17\3\20\3\20\3\20\3\20\3\21\3"+
		"\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3"+
		"\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3"+
		"\23\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3"+
		"\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3"+
		"\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3"+
		"\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3"+
		"\31\3\31\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\36\6\36\u00fa"+
		"\n\36\r\36\16\36\u00fb\3\37\3\37\3\37\3\37\3\37\4l\u00fb\2 \4\2\6\3\b"+
		"\4\n\2\f\5\16\2\20\6\22\7\24\b\26\t\30\34\32\n\34\2\36\13 \f\"\r$\16&"+
		"\17(\20*\21,\22.\23\60\24\62\25\64\26\66\278\30:\31<\32>\33\4\2\3\b\5"+
		"\2C\\aac|\6\2\62;C\\aac|\t\2#(*\61>B]]_`bb}\u0080\5\2\13\f\17\17\"\"\3"+
		"\2^^\4\2\f\f\17\17\2\u0108\2\b\3\2\2\2\2\n\3\2\2\2\2\f\3\2\2\2\2\16\3"+
		"\2\2\2\2\20\3\2\2\2\2\22\3\2\2\2\2\24\3\2\2\2\2\26\3\2\2\2\2\30\3\2\2"+
		"\2\2\32\3\2\2\2\3\34\3\2\2\2\3\36\3\2\2\2\3 \3\2\2\2\3\"\3\2\2\2\3$\3"+
		"\2\2\2\3&\3\2\2\2\3(\3\2\2\2\3*\3\2\2\2\3,\3\2\2\2\3.\3\2\2\2\3\60\3\2"+
		"\2\2\3\62\3\2\2\2\3\64\3\2\2\2\3\66\3\2\2\2\38\3\2\2\2\3:\3\2\2\2\3<\3"+
		"\2\2\2\3>\3\2\2\2\4@\3\2\2\2\6B\3\2\2\2\bI\3\2\2\2\nN\3\2\2\2\fS\3\2\2"+
		"\2\16X\3\2\2\2\20^\3\2\2\2\22a\3\2\2\2\24g\3\2\2\2\26s\3\2\2\2\30u\3\2"+
		"\2\2\32z\3\2\2\2\34\u0088\3\2\2\2\36\u008d\3\2\2\2 \u0093\3\2\2\2\"\u0097"+
		"\3\2\2\2$\u009d\3\2\2\2&\u00a4\3\2\2\2(\u00b5\3\2\2\2*\u00ba\3\2\2\2,"+
		"\u00c0\3\2\2\2.\u00c6\3\2\2\2\60\u00d5\3\2\2\2\62\u00e6\3\2\2\2\64\u00f0"+
		"\3\2\2\2\66\u00f2\3\2\2\28\u00f4\3\2\2\2:\u00f6\3\2\2\2<\u00f9\3\2\2\2"+
		">\u00fd\3\2\2\2@A\7)\2\2A\5\3\2\2\2BF\t\2\2\2CE\t\3\2\2DC\3\2\2\2EH\3"+
		"\2\2\2FD\3\2\2\2FG\3\2\2\2G\7\3\2\2\2HF\3\2\2\2IJ\7\61\2\2JK\7,\2\2KL"+
		"\3\2\2\2LM\b\4\2\2M\t\3\2\2\2NO\5\6\3\2OP\3\2\2\2PQ\b\5\3\2Q\13\3\2\2"+
		"\2RT\t\3\2\2SR\3\2\2\2TU\3\2\2\2US\3\2\2\2UV\3\2\2\2V\r\3\2\2\2WY\t\4"+
		"\2\2XW\3\2\2\2YZ\3\2\2\2ZX\3\2\2\2Z[\3\2\2\2[\\\3\2\2\2\\]\b\7\4\2]\17"+
		"\3\2\2\2^_\7=\2\2_\21\3\2\2\2`b\t\5\2\2a`\3\2\2\2bc\3\2\2\2ca\3\2\2\2"+
		"cd\3\2\2\2de\3\2\2\2ef\b\t\5\2f\23\3\2\2\2gq\5\4\2\2hr\5\4\2\2ik\13\2"+
		"\2\2ji\3\2\2\2kn\3\2\2\2lm\3\2\2\2lj\3\2\2\2mo\3\2\2\2nl\3\2\2\2op\n\6"+
		"\2\2pr\5\4\2\2qh\3\2\2\2ql\3\2\2\2r\25\3\2\2\2st\7<\2\2t\27\3\2\2\2uv"+
		"\7<\2\2vw\7<\2\2wx\3\2\2\2xy\b\f\4\2y\31\3\2\2\2z{\7/\2\2{|\7/\2\2|\u0080"+
		"\3\2\2\2}\177\n\7\2\2~}\3\2\2\2\177\u0082\3\2\2\2\u0080~\3\2\2\2\u0080"+
		"\u0081\3\2\2\2\u0081\u0084\3\2\2\2\u0082\u0080\3\2\2\2\u0083\u0085\7\17"+
		"\2\2\u0084\u0083\3\2\2\2\u0084\u0085\3\2\2\2\u0085\u0086\3\2\2\2\u0086"+
		"\u0087\7\f\2\2\u0087\33\3\2\2\2\u0088\u0089\5\6\3\2\u0089\u008a\3\2\2"+
		"\2\u008a\u008b\b\16\3\2\u008b\35\3\2\2\2\u008c\u008e\t\5\2\2\u008d\u008c"+
		"\3\2\2\2\u008e\u008f\3\2\2\2\u008f\u008d\3\2\2\2\u008f\u0090\3\2\2\2\u0090"+
		"\u0091\3\2\2\2\u0091\u0092\b\17\5\2\u0092\37\3\2\2\2\u0093\u0094\7\60"+
		"\2\2\u0094\u0095\7\60\2\2\u0095\u0096\7\60\2\2\u0096!\3\2\2\2\u0097\u0098"+
		"\7B\2\2\u0098\u0099\7p\2\2\u0099\u009a\7c\2\2\u009a\u009b\7o\2\2\u009b"+
		"\u009c\7g\2\2\u009c#\3\2\2\2\u009d\u009e\7B\2\2\u009e\u009f\7r\2\2\u009f"+
		"\u00a0\7c\2\2\u00a0\u00a1\7t\2\2\u00a1\u00a2\7c\2\2\u00a2\u00a3\7o\2\2"+
		"\u00a3%\3\2\2\2\u00a4\u00a5\7B\2\2\u00a5\u00a6\7r\2\2\u00a6\u00a7\7c\2"+
		"\2\u00a7\u00a8\7t\2\2\u00a8\u00a9\7c\2\2\u00a9\u00aa\7o\2\2\u00aa\u00ab"+
		"\7U\2\2\u00ab\u00ac\7v\2\2\u00ac\u00ad\7t\2\2\u00ad\u00ae\7w\2\2\u00ae"+
		"\u00af\7e\2\2\u00af\u00b0\7v\2\2\u00b0\u00b1\7P\2\2\u00b1\u00b2\7c\2\2"+
		"\u00b2\u00b3\7o\2\2\u00b3\u00b4\7g\2\2\u00b4\'\3\2\2\2\u00b5\u00b6\7B"+
		"\2\2\u00b6\u00b7\7q\2\2\u00b7\u00b8\7p\2\2\u00b8\u00b9\7g\2\2\u00b9)\3"+
		"\2\2\2\u00ba\u00bb\7B\2\2\u00bb\u00bc\7o\2\2\u00bc\u00bd\7c\2\2\u00bd"+
		"\u00be\7p\2\2\u00be\u00bf\7{\2\2\u00bf+\3\2\2\2\u00c0\u00c1\7B\2\2\u00c1"+
		"\u00c2\7g\2\2\u00c2\u00c3\7z\2\2\u00c3\u00c4\7g\2\2\u00c4\u00c5\7e\2\2"+
		"\u00c5-\3\2\2\2\u00c6\u00c7\7B\2\2\u00c7\u00c8\7p\2\2\u00c8\u00c9\7q\2"+
		"\2\u00c9\u00ca\7v\2\2\u00ca\u00cb\7P\2\2\u00cb\u00cc\7w\2\2\u00cc\u00cd"+
		"\7n\2\2\u00cd\u00ce\7n\2\2\u00ce\u00cf\7R\2\2\u00cf\u00d0\7c\2\2\u00d0"+
		"\u00d1\7t\2\2\u00d1\u00d2\7c\2\2\u00d2\u00d3\7o\2\2\u00d3\u00d4\7u\2\2"+
		"\u00d4/\3\2\2\2\u00d5\u00d6\7B\2\2\u00d6\u00d7\7t\2\2\u00d7\u00d8\7g\2"+
		"\2\u00d8\u00d9\7v\2\2\u00d9\u00da\7w\2\2\u00da\u00db\7t\2\2\u00db\u00dc"+
		"\7p\2\2\u00dc\u00dd\7X\2\2\u00dd\u00de\7c\2\2\u00de\u00df\7n\2\2\u00df"+
		"\u00e0\7w\2\2\u00e0\u00e1\7g\2\2\u00e1\u00e2\7P\2\2\u00e2\u00e3\7c\2\2"+
		"\u00e3\u00e4\7o\2\2\u00e4\u00e5\7g\2\2\u00e5\61\3\2\2\2\u00e6\u00e7\7"+
		"B\2\2\u00e7\u00e8\7v\2\2\u00e8\u00e9\7g\2\2\u00e9\u00ea\7o\2\2\u00ea\u00eb"+
		"\7r\2\2\u00eb\u00ec\7n\2\2\u00ec\u00ed\7c\2\2\u00ed\u00ee\7v\2\2\u00ee"+
		"\u00ef\7g\2\2\u00ef\63\3\2\2\2\u00f0\u00f1\7*\2\2\u00f1\65\3\2\2\2\u00f2"+
		"\u00f3\7+\2\2\u00f3\67\3\2\2\2\u00f4\u00f5\7\60\2\2\u00f59\3\2\2\2\u00f6"+
		"\u00f7\7.\2\2\u00f7;\3\2\2\2\u00f8\u00fa\13\2\2\2\u00f9\u00f8\3\2\2\2"+
		"\u00fa\u00fb\3\2\2\2\u00fb\u00fc\3\2\2\2\u00fb\u00f9\3\2\2\2\u00fc=\3"+
		"\2\2\2\u00fd\u00fe\7,\2\2\u00fe\u00ff\7\61\2\2\u00ff\u0100\3\2\2\2\u0100"+
		"\u0101\b\37\6\2\u0101?\3\2\2\2\16\2\3FUZclq\u0080\u0084\u008f\u00fb\7"+
		"\4\3\2\t\3\2\t\5\2\b\2\2\4\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}