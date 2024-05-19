package parser

import (
	"fmt"
	"strconv"

	"github.com/vitorwdson/monkey-interpreter/ast"
	"github.com/vitorwdson/monkey-interpreter/token"
)

type Precedence int

const (
	_ Precedence = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

func (p *Parser) setParseFunctions() {
	p.prefixParseFns[token.IDENTIFIER] = p.parseIdentifier
	p.prefixParseFns[token.INT] = p.parseIntegerLiteral
}

func (p *Parser) parseExpression(precedence Precedence) ast.Expression {
	prefix, ok := p.prefixParseFns[p.currToken.Type]
	if !ok {
		return nil
	}

	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	integer, err := strconv.ParseInt(p.currToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currToken.Literal)

		p.errors = append(p.errors, ParserError{
			token:   p.peekToken,
			message: msg,
		})

		return nil
	}

	return &ast.IntegerLiteral{
		Token: p.currToken,
		Value: integer,
	}
}
