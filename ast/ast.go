package ast

import "asteroid/token"

// Every node in the AST must provide the TokenLiteral() function.
// This token literal function returns the literal value of the token it is associated with.
type Node interface {
	TokenLiteral() string
}

// This interface ensures that every statement has a statementNode() method
type Statement interface {
	Node
	statementNode()
}

// This interface ensures that every expression has an expressionNode() method
type Expression interface {
	Node
	expressionNode()
}

// A program is an array of statements.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	// if the program consists of at least one statement
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // token it is associated with
	Name  *Identifier // variable name
	Value Expression  // value of the variable eg. "5"
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string      // the value of the identifier eg. "x"
}

// The identifier of a let statement is considered to be an expression.
// this keeps things simple
// identifier in other parts of the language do produce values and are hence expressions
// this keeps the number of types of nodes small
// the identifier here represents the name of a variable.
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
