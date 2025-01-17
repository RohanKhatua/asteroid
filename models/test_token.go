package models

import "asteroid/token"

type TestTokenType struct {
	ExpectedType    token.TokenType
	ExpectedLiteral string
}
