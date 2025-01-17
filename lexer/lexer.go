package lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New returns an object of type Lexer and takes an input string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// if we have reached the end of input - set the ch to 0 - ASCII code for null.
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.position += 1
}
