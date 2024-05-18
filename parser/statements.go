package parser

import (
	"github.com/vitorwdson/monkey-interpreter/ast"
	"github.com/vitorwdson/monkey-interpreter/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	}

	return nil
}

func (p *Parser) parseLetStatement() ast.Statement {
	statement := &ast.LetStatement{Token: p.currToken}

	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	statement.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: get the expression assigned to the variable
	for p.currToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseReturnStatement() ast.Statement {
	statement := &ast.ReturnStatement{Token: p.currToken}
	p.nextToken()

	// TODO: get the expression assigned to the variable
	for p.currToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return statement
}
