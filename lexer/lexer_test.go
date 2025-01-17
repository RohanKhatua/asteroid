package lexer

import (
	"asteroid/models"
	"asteroid/token"
	"testing"
)

// TestNextToken tests the nextToken function of the lexer
// Modification made - added multiple test cases to test different inputs in one go.
func TestNextToken(t *testing.T) {

	// Define multiple test cases with different inputs
	nextTokenTestList := []models.TestInput{
		{
			Input: `=+(){},;`,
			Tests: []models.TestTokenType{
				{ExpectedType: token.ASSIGN, ExpectedLiteral: "="},
				{ExpectedType: token.PLUS, ExpectedLiteral: "+"},
				{ExpectedType: token.LPAREN, ExpectedLiteral: "("},
				{ExpectedType: token.RPAREN, ExpectedLiteral: ")"},
				{ExpectedType: token.LBRACE, ExpectedLiteral: "{"},
				{ExpectedType: token.RBRACE, ExpectedLiteral: "}"},
				{ExpectedType: token.COMMA, ExpectedLiteral: ","},
				{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
				{ExpectedType: token.EOF, ExpectedLiteral: ""},
			},
		},
		{
			Input: `let five = 5;
					let ten = 10;
					let add = fn(x, y) {
						x + y;
					};
					let result = add(five, ten);`,
			Tests: []models.TestTokenType{
				{ExpectedType: token.LET, ExpectedLiteral: "let"},
				{ExpectedType: token.IDENT, ExpectedLiteral: "five"},
				{ExpectedType: token.ASSIGN, ExpectedLiteral: "="},
				{ExpectedType: token.INT, ExpectedLiteral: "5"},
				{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
				{ExpectedType: token.LET, ExpectedLiteral: "let"},
				{ExpectedType: token.IDENT, ExpectedLiteral: "ten"},
				{ExpectedType: token.ASSIGN, ExpectedLiteral: "="},
				{ExpectedType: token.INT, ExpectedLiteral: "10"},
				{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
				{ExpectedType: token.LET, ExpectedLiteral: "let"},
				{ExpectedType: token.IDENT, ExpectedLiteral: "add"},
				{ExpectedType: token.ASSIGN, ExpectedLiteral: "="},
				{ExpectedType: token.FUNCTION, ExpectedLiteral: "fn"},
				{ExpectedType: token.LPAREN, ExpectedLiteral: "("},
				{ExpectedType: token.IDENT, ExpectedLiteral: "x"},
				{ExpectedType: token.COMMA, ExpectedLiteral: ","},
				{ExpectedType: token.IDENT, ExpectedLiteral: "y"},
				{ExpectedType: token.RPAREN, ExpectedLiteral: ")"},
				{ExpectedType: token.LBRACE, ExpectedLiteral: "{"},
				{ExpectedType: token.IDENT, ExpectedLiteral: "x"},
				{ExpectedType: token.PLUS, ExpectedLiteral: "+"},
				{ExpectedType: token.IDENT, ExpectedLiteral: "y"},
				{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
				{ExpectedType: token.RBRACE, ExpectedLiteral: "}"},
				{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
				{ExpectedType: token.LET, ExpectedLiteral: "let"},
				{ExpectedType: token.IDENT, ExpectedLiteral: "result"},
				{ExpectedType: token.ASSIGN, ExpectedLiteral: "="},
				{ExpectedType: token.IDENT, ExpectedLiteral: "add"},
				{ExpectedType: token.LPAREN, ExpectedLiteral: "("},
				{ExpectedType: token.IDENT, ExpectedLiteral: "five"},
				{ExpectedType: token.COMMA, ExpectedLiteral: ","},
				{ExpectedType: token.IDENT, ExpectedLiteral: "ten"},
				{ExpectedType: token.RPAREN, ExpectedLiteral: ")"},
				{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
				{ExpectedType: token.EOF, ExpectedLiteral: ""},
			},
		},
		// Add more test cases as needed
		// {
		// 	input: `your_input_here`,
		// 	tests: []TestTokenType{
		// 		{token.YOUR_TOKEN_TYPE, "your_literal"},
		// 		// ...more tokens...
		// 	},
		// },
	}

	for _, testInput := range nextTokenTestList {
		// Using the subtest feature of testing to run multiple test cases in one go.
		// This will help in identifying which test case failed and which passed.
		// Allows for verbose output.
		t.Run(testInput.Input, func(t *testing.T) {
			l := New(testInput.Input)

			for i, tt := range testInput.Tests {
				tok := l.nextToken()

				if tok.Type != tt.ExpectedType {
					t.Errorf("tests[%d] - token type wrong, expected=%q, got=%q", i, tt.ExpectedType, tok.Type)
				}

				if tok.Literal != tt.ExpectedLiteral {
					t.Errorf("tests[%d] - literal wrong, expected=%q, got=%q", i, tt.ExpectedLiteral, tok.Literal)
				}
			}
		})
	}

}
