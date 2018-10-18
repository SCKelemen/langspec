package lexer

import (
	"LANGSPEC/token"
	"unicode"
	"unicode/utf8"
)

type Scanner struct {
	input        string
	position     int
	readPosition int
	ch           rune
	src          []byte
}

func New(input string) *Scanner {
	s := &Scanner{input: input}
	s.readChar()
	return l
}

func (s *Scanner) NextToken() token.Token {

}

func (s *Scanner) scanIdentifier() string {
	offset := s.position
	for isLetter(s.ch) || isDigit(s.ch) {
		s.next()
	}
	return string(s.src[offset:s.position])
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9' || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}

func newLexeme(token token.Token, ch byte) Lexeme {
	return Lexeme{Type: token, Literal: string(ch)}
}

type Lexeme struct {
	Type    Token
	Literal string
}
