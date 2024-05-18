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
	MINUS
	BANG
	ASTERISK
	SLASH

	// Comparison
	GT
	LT

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
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)

func New(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}
