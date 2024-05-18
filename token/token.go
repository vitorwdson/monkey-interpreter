package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special tokens
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + Literals
	IDENTIFIER TokenType = "IDENTIFIER"
	INT        TokenType = "INT"

	// Operators
	ASSIGN TokenType = "ASSIGN"
	PLUS   TokenType = "PLUS"

	// Delimiters
	COMMA     TokenType = "COMMA"
	SEMICOLON TokenType = "SEMICOLON"
	LPAREN    TokenType = "LPAREN"
	RPAREN    TokenType = "RPAREN"
	LBRACE    TokenType = "LBRACE"
	RBRACE    TokenType = "RBRACE"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)
