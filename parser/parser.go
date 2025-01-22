package parser

import (
	"asteroid/ast"
	"asteroid/lexer"
	"asteroid/token"
	"fmt"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l,
		errors: []string{}}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Parses the program - constructs an AST and an array of statements
func (p *Parser) ParseProgram() *ast.Program {
	// construct the root node of an AST
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// Iterate over every token until we reach EOF
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// parses a single statement
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// check if next token is an identifier
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// if it is indeed an identifier
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// Currently, we skip tokens unless we get to a semicolon
	// effectively skipping over the expression part
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		// if the next token's type does not match our expectation - we add it to the list of errors.
		p.peekError(t)
		return false
	}
}

// Helper fucntions to get all the errors encountered by the parser
func (p *Parser) Errors() []string {
	return p.errors
}

// peek error accepts the wrong token that was seen and adds an error 
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("Expected next token to be '%s' got '%s' instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
