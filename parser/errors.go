package parser

import "github.com/vitorwdson/monkey-interpreter/token"

type ParserError struct {
	token   token.Token
	message string
}

func (e ParserError) Error() string {
	return e.message
}
