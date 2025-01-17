package lexer

import "asteroid/token"

type Lexer struct {
	input        string
	position     int
	readPosition int  //points to the next position from which we are going to read.
	ch           byte //position where we last read.
}

// New returns an object of type Lexer and takes an input string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/*
readChar is used to read through the input string.
Only supports UTF-8.
readChar is bound to the *Lexer type. Thus only a pointer to a Lexer object can call this function.
As this is a pointer *Lexer as opposed to lexer it can modify values associated with the object that calls it as opposed to working on a copy.
This function cannot be used as a standalone function - only invoked by a *Lexer object.
*/
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// if we have reached the end of input - set the ch to 0 - ASCII code for null.
		// Either nothing was read or we have reached the end of file
		l.ch = 0
	} else {
		// if we have not reached the end of the file - we set the current character to that at readPosition
		l.ch = l.input[l.readPosition]
	}
	// update the pointers
	// current position is updated to the position that was just read (readPosition)
	l.position = l.readPosition
	// readPosition is icnremented by one.
	l.readPosition += 1
}

// returns the next token one by one - each token consists of a type and literal.
// The literal is the value of the token.
// The type is the type of the token.
func (l *Lexer) nextToken() token.Token {
	var tok token.Token
	// check the current character and assign the token type accordingly
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		// Default case - EOF
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

// newToken is a helper function to create a new token.
// It takes a tokenType and a character and returns a token object.
// Note - this can also be done via a constructor function in token/token.go
// We are doing it here to keep the lexer and token packages separate and to maintain consistency with the book.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
