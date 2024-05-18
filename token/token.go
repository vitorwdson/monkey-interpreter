package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
	Row     int
	Col     int
}

const (
	// Special tokens
	ILLEGAL TokenType = iota
	EOF

	// Identifiers + Literals
	IDENTIFIER
	INT

	// Operators
	ASSIGN
	PLUS

	// Delimiters
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	FUNCTION
	LET
)

func New(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}
