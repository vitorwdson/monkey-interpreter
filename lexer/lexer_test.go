package lexer

import (
	"testing"

	"github.com/vitorwdson/monkey-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"
	tests := []token.Token{
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.SEMICOLON, Literal: ";"},
	}

	lexer := New(input)
	for i, expected := range tests {
		result := lexer.NextToken()

		if expected.Type != result.Type {
			t.Fatalf(
				"tests[%d] - incorrect token type: expected %q, got %q",
				i,
				expected.Type,
				result.Type,
			)
		}

		if expected.Literal != result.Literal {
			t.Fatalf(
				"tests[%d] - incorrect token literal: expected %q, got %q",
				i,
				expected.Literal,
				result.Literal,
			)
		}
	}
}
