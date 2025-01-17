package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// Define the token types - these are the types of tokens that the lexer can produce.
// Very very preliminary - we will add more as we go along.
const (
	// miscellaenous
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
