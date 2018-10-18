package token

import "strconv"

type Token int

const (
	ILLEGAL Token = iota
	EOF
	COMMENT

	// literals
	literals
	IDENT  // main
	INT    // 12345
	FLOAT  // 123.45
	IMAG   // 123.45i
	CHAR   // 'a'
	STRING // "abc"
	literals_end

	// operators
	operators
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	EXP // **
	REM // %

	AND     // &
	OR      // |
	XOR     // ^
	SHL     // <<
	SHR     // >>
	AND_NOT // &^

	ADD_ASS // +=
	SUB_ASS // -=
	MUL_ASS // *=
	QUO_ASS // /=
	EXP_ASS // **=
	REM_ASS // %=

	AND_ASS     // &=
	OR_ASS      // |=
	XOR_ASS     // ^=
	SHL_ASS     // <<=
	SHR_ASS     // >>=
	AND_NOT_ASS // &^=

	LOG_AND // &&
	LOG_OR  // ||

	INC // ++
	DEC // --

	EQL    // ==
	LT     // <
	GT     // >
	ASSIGN // =
	NOT    // !

	NEQ      // !=
	LTE      // <=
	GTE      // >=
	DEF      // :=
	ELLIPSIS // ...ADD

	LPAREN // (
	RPAREN // )

	LBRACK // [
	RBRACK // ]

	LBRACE // {
	RBRACE // }

	COMMA // ,
	DOT   // .
	SEMI  // ;
	COLON // :
	operators_end

	// keywords
	keywords_init
	// control
	IF
	ELSE
	SWITCH
	CASE
	BREAK
	CONTINUE

	// func
	FUNCTION
	RETURN

	// types
	TYPE
	STRUCT
	CLASS
	INTERFACE
	ACTOR

	// modules
	PACKAGE
	IMPORT

	// builtin
	SELECT
	RANGE
	MAP

	// constraints
	WHERE
	MATCH
	IS

	// declarations
	CONST
	VAR
	// perhaps: public, private, protected, abstract, sealed
	keywords_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",

	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT:  "IDENT",
	INT:    "INT",
	FLOAT:  "FLOAT",
	IMAG:   "IMAG",
	CHAR:   "CHAR",
	STRING: "STRING",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	EXP: "**",
	REM: "%",

	AND:     "&",
	OR:      "|",
	XOR:     "^",
	SHL:     "<<",
	SHR:     ">>",
	AND_NOT: "&^",

	ADD_ASS: "+=",
	SUB_ASS: "-=",
	MUL_ASS: "*=",
	QUO_ASS: "/=",
	EXP_ASS: "**=",
	REM_ASS: "%=",

	AND_ASS:     "&=",
	OR_ASS:      "|=",
	XOR_ASS:     "^=",
	SHL_ASS:     "<<=",
	SHR_ASS:     ">>=",
	AND_NOT_ASS: "&^=",

	LOG_AND: "&&",
	LOG_OR:  "||",

	INC: "++",
	DEC: "--",

	EQL:    "==",
	LT:     "<",
	GT:     ">",
	ASSIGN: "=",
	NOT:    "!",

	NEQ:      "!=",
	LTE:      "<=",
	GTE:      ">=",
	DEF:      ":=",
	ELLIPSIS: "...",

	LPAREN: "(",
	RPAREN: ")",

	LBRACK: "[",
	RBRACK: "]",

	LBRACE: "{",
	RBRACE: "}",

	COMMA: ",",
	DOT:   ".",
	SEMI:  ";",
	COLON: ":",

	BREAK:    "break",
	CONTINUE: "continue",

	IF:     "if",
	ELSE:   "else",
	SWITCH: "switch",
	CASE:   "case",

	// probably need a while and for

	FUNCTION: "func",
	RETURN:   "return",

	TYPE:      "type",
	STRUCT:    "struct",
	CLASS:     "class",
	ACTOR:     "actor",
	INTERFACE: "interface",

	PACKAGE: "package",
	IMPORT:  "import",

	CONST: "const",
	VAR:   "var",

	SELECT: "select",
	WHERE:  "where",
	MATCH:  "match",
	IS:     "is",

	RANGE: "range",
	MAP:   "map",
}

func (token Token) String() string {
	s := ""
	if 0 <= token && token < Token(len(tokens)) {
		s = tokens[token]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(token)) + ")"
	}
	return s
}

const (
	LOW   = 0
	UNARY = 6
	HIGH  = 7
)

func (op Token) Precedence() int {
	switch op {
	case LOG_OR:
		return 1
	case LOG_AND:
		return 2
	case EQL, NEQ, LT, LTE, GT, GTE:
		return 3
	case ADD, SUB, OR, XOR:
		return 4
	case MUL, QUO, REM, EXP, SHL, SHR, AND, AND_NOT:
		return 5
	}
	return LOW
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token)
	for i := keywords_init + 1; i < keywords_end; i++ {
		keywords[tokens[i]] = i
	}
}

func Lookup(ident string) Token {
	if token, is_keyword := keywords[ident]; is_keyword {
		return token
	}
	return IDENT
}

func (token Token) IsLiteral() bool  { return literals < token && token < literals_end }
func (token Token) IsOperator() bool { return operators < token && token < operators_end }
func (token Token) IsKeyword() bool  { return keywords_init < token && token < keywords_end }
