package lexer

import (
	"asteroid/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	type TestTokenType struct {
		expectedType    token.TokenType
		expectedLiteral string
	}

	// tests in as array of TestTokenType
	// each TestTokenType consists of an expected type and an expected literal corresponding to that type

	tests := []TestTokenType{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l = New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
