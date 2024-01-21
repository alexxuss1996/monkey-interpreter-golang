package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	char         byte // current char under examination
}

// New returns a new Lexer with the given input.
//
// input string
// *Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar()
	return l
}

// ReadChar reads the next character of input and advances the position in the lexer.
//
// No parameters.
// No return value.
func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// NewToken creates a new token of the given TokenType with the provided character.
//
// TokenType: type of the token.
// char: character for the token.
// token: the newly created token.
func NewToken(TokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: TokenType, Literal: string(char)}
}

// NextToken returns the next token in the lexer.
//
// It does not take any parameters and returns a token.Token.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.char {
	case '=':
		tok = NewToken(token.ASSIGN, l.char)
	case '+':
		tok = NewToken(token.PLUS, l.char)
	case '-':
		tok = NewToken(token.MINUS, l.char)
	case '!':
		tok = NewToken(token.BANG, l.char)
	case ',':
		tok = NewToken(token.COMMA, l.char)
	case ';':
		tok = NewToken(token.SEMICOLON, l.char)
	case '(':
		tok = NewToken(token.LPAREN, l.char)
	case ')':
		tok = NewToken(token.RPAREN, l.char)
	case '{':
		tok = NewToken(token.LBRACE, l.char)
	case '}':
		tok = NewToken(token.RBRACE, l.char)
	case '#':
		tok = NewToken(token.COMMENT, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.ReadChar()
	return tok
}
