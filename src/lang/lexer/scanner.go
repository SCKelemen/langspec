package lexer

type Scanner struct {
	input        string
	position     int
	readPosition int
	ch           byte
}
