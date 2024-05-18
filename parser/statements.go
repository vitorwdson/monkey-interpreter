package parser

import (
	"github.com/vitorwdson/monkey-interpreter/ast"
	"github.com/vitorwdson/monkey-interpreter/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.currToken}

	if p.peekToken.Type != token.IDENTIFIER {
		return nil
	}
	p.nextToken()

	statement.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	if p.peekToken.Type != token.ASSIGN {
		return nil
	}
	p.nextToken()

	// TODO: get the expression assigned to the variable
	for p.currToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return statement
}
