package parser

import (
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
