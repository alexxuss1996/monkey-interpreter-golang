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

// isLetter checks if the given byte is a letter.
//
// char byte
// bool
func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

// NewToken creates a new token of the given TokenType with the provided character.
//
// TokenType: type of the token.
// char: character for the token.
// token: the newly created token.
func NewToken(TokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: TokenType, Literal: string(char)}
}

// readIdentifier returns a string representing the identifier read by the lexer.
//
// No parameters.
// Returns a string.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.ReadChar()
	}
	return l.input[position:l.position]
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
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			return tok
		} else {
			tok = NewToken(token.ILLEGAL, l.char)
		}
	}
	l.ReadChar()
	return tok
}
