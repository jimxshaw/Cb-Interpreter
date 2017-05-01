package lexer

// Lexer transforms source code to tokens.
// For simplicity, our lexer will only read ASCII as each of its
// character is only ever 1 byte long.
type Lexer struct {
	input        string
	position     int  // Current position in input (points to current char).
	readPosition int  // Current reading position in input (after current char).
	ch           byte // Current char being examined.
}

// New is essentially a constructor function.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// The purpose of this helper method is to give us the next character and
// advance our position in the input string.
func (l *Lexer) readChar() {
	// Check if we've reached the end of input. If yes then assign ch to
	// 0, which essentially means end of file.
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// If it's not the end of input then assign ch the next character.
		l.ch = l.input[l.readPosition]
	}

	// Assign the position we've just read to the current position and
	// increment the current reading position by 1.
	l.position = l.readPosition
	l.readPosition++
}
