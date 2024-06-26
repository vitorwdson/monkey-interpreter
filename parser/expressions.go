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

var precedences = map[token.TokenType]Precedence{
	token.EQ:       EQUALS,
	token.NEQ:      EQUALS,
	token.LT:       LESSGREATER,
	token.LTE:      LESSGREATER,
	token.GT:       LESSGREATER,
	token.GTE:      LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}

func (p *Parser) peekPrecedence() Precedence {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) currPrecedence() Precedence {
	if p, ok := precedences[p.currToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) setParseFunctions() {
	p.prefixParseFns[token.IDENTIFIER] = p.parseIdentifier
	p.prefixParseFns[token.INT] = p.parseIntegerLiteral
	p.prefixParseFns[token.TRUE] = p.parseBooleanExpression
	p.prefixParseFns[token.FALSE] = p.parseBooleanExpression
	p.prefixParseFns[token.BANG] = p.parsePrefixExpression
	p.prefixParseFns[token.MINUS] = p.parsePrefixExpression
	p.prefixParseFns[token.LPAREN] = p.parseGroupedExpression
	p.prefixParseFns[token.IF] = p.parseIfExpression
	p.prefixParseFns[token.FUNCTION] = p.parseFunctionLiteral

	p.infixParseFns[token.PLUS] = p.parseInfixExpression
	p.infixParseFns[token.MINUS] = p.parseInfixExpression
	p.infixParseFns[token.ASTERISK] = p.parseInfixExpression
	p.infixParseFns[token.SLASH] = p.parseInfixExpression
	p.infixParseFns[token.GT] = p.parseInfixExpression
	p.infixParseFns[token.GTE] = p.parseInfixExpression
	p.infixParseFns[token.LT] = p.parseInfixExpression
	p.infixParseFns[token.LTE] = p.parseInfixExpression
	p.infixParseFns[token.EQ] = p.parseInfixExpression
	p.infixParseFns[token.NEQ] = p.parseInfixExpression
	p.infixParseFns[token.LPAREN] = p.parseCallExpression
}

func (p *Parser) parseExpression(precedence Precedence) ast.Expression {
	prefix, ok := p.prefixParseFns[p.currToken.Type]
	if !ok {
		msg := fmt.Sprintf("no prefix parse function for %s found", p.currToken.Literal)
		p.errors = append(p.errors, ParserError{
			token:   p.peekToken,
			message: msg,
		})
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix, ok := p.infixParseFns[p.peekToken.Type]
		if !ok {
			return leftExp
		}

		p.nextToken()
		leftExp = infix(leftExp)
	}

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

func (p *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
	}

	p.nextToken()
	exp.Right = p.parseExpression(PREFIX)

	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
		Left:     left,
	}

	precedence := p.currPrecedence()
	p.nextToken()
	exp.Right = p.parseExpression(precedence)

	return exp
}

func (p *Parser) parseBooleanExpression() ast.Expression {
	exp := &ast.Boolean{
		Token: p.currToken,
		Value: p.currToken.Type == token.TRUE,
	}
	return exp
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseIfExpression() ast.Expression {
	exp := &ast.IfExpression{Token: p.currToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	exp.Condition = p.parseGroupedExpression()
	if exp.Condition == nil {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	exp.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		exp.Alternative = p.parseBlockStatement()
	}

	return exp
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	exp := &ast.FunctionLiteral{
		Token:      p.currToken,
		Parameters: []*ast.Identifier{},
	}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	for p.currTokenIs(token.IDENTIFIER) {
		exp.Parameters = append(
			exp.Parameters,
			p.parseIdentifier().(*ast.Identifier),
		)

		if p.peekTokenIs(token.COMMA) {
			p.nextToken()
		}
		p.nextToken()
	}

	if !p.currTokenIs(token.RPAREN) {
		p.peekError(token.RPAREN)
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	exp.Body = p.parseBlockStatement()

	return exp
}

func (p *Parser) parseExpressionList() []ast.Expression {
	expressions := []ast.Expression{}

	fmt.Println(p.currToken)
	expressions = append(expressions, p.parseExpression(LOWEST))
	fmt.Println(p.currToken)
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		expressions = append(expressions, p.parseExpression(LOWEST))
	}

	return expressions
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{
		Token:     p.currToken,
		Function:  function,
		Arguments: []ast.Expression{},
	}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return exp
	}

	p.nextToken()
	exp.Arguments = p.parseExpressionList()

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}
