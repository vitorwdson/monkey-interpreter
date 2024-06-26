package lexer

import (
	"github.com/vitorwdson/monkey-interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readNumber() string {
	startPosition := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		tok = token.New(token.LPAREN, "(")
	case ')':
		tok = token.New(token.RPAREN, ")")
	case '{':
		tok = token.New(token.LBRACE, "{")
	case '}':
		tok = token.New(token.RBRACE, "}")
	case ',':
		tok = token.New(token.COMMA, ",")
	case ';':
		tok = token.New(token.SEMICOLON, ";")

	case '+':
		tok = token.New(token.PLUS, "+")
	case '-':
		tok = token.New(token.MINUS, "-")
	case '*':
		tok = token.New(token.ASTERISK, "*")
	case '/':
		tok = token.New(token.SLASH, "/")

	case '=':
		tok = l.getComposedToken(l.ch, '=', token.ASSIGN, token.EQ)
	case '!':
		tok = l.getComposedToken(l.ch, '=', token.BANG, token.NEQ)
	case '>':
		tok = l.getComposedToken(l.ch, '=', token.GT, token.GTE)
	case '<':
		tok = l.getComposedToken(l.ch, '=', token.LT, token.LTE)

	case 0:
		tok = token.New(token.EOF, "")

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) getComposedToken(ch byte, next byte, normalType, composedType token.TokenType) token.Token {
	if next != l.peekChar() {
		return token.New(normalType, string(ch))
	}

	l.readChar()
	literal := string(ch) + string(next)
	return token.New(composedType, literal)
}

func isLetter(ch byte) bool {
	isLowercase := ch >= 'a' && ch <= 'z'
	isUppercase := ch >= 'A' && ch <= 'Z'
	return isLowercase || isUppercase || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == '\n' || ch == '\r' || ch == '\t' || ch == ' '
}
