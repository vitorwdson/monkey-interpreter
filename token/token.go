package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Row     int
	Col     int
}

const (
	// Special tokens
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + Literals
	IDENTIFIER TokenType = "IDENTIFIER"
	INT        TokenType = "INT"

	// Operators
	ASSIGN   TokenType = "ASSIGN"
	PLUS     TokenType = "PLUS"
	MINUS    TokenType = "MINUS"
	BANG     TokenType = "BANG"
	ASTERISK TokenType = "ASTERISK"
	SLASH    TokenType = "SLASH"

	// Comparison
	EQ  TokenType = "EQ"
	NEQ TokenType = "NEQ"
	GT  TokenType = "GT"
	GTE TokenType = "GTE"
	LT  TokenType = "LT"
	LTE TokenType = "LTE"

	// Delimiters
	COMMA     TokenType = "COMMA"
	SEMICOLON TokenType = "SEMICOLON"
	LPAREN    TokenType = "LPAREN"
	RPAREN    TokenType = "RPAREN"
	LBRACE    TokenType = "LBRACE"
	RBRACE

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
)

func New(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}
