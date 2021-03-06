package lexer

import "cb/token"

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
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

// NextToken looks at the current character being examined and return
// a token depending on which character it is.
func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	lex.skipWhitespace()

	switch lex.ch {
	case '=':
		// If the current char is = then peek at the next char.
		// If the next char is also = then it's an Equal token
		// rather than an Assign token.
		// The char is saved in a local variable before calling
		// readChar() because we don't want to lose the current
		// char and can safely advance the lexer so it leaves
		// the NextToken() with lex.position and lex.readPosition
		// in the right state.
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			literal := string(ch) + string(lex.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, lex.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, lex.ch)
	case '(':
		tok = newToken(token.LPAREN, lex.ch)
	case ')':
		tok = newToken(token.RPAREN, lex.ch)
	case '{':
		tok = newToken(token.LBRACE, lex.ch)
	case '}':
		tok = newToken(token.RBRACE, lex.ch)
	case ',':
		tok = newToken(token.COMMA, lex.ch)
	case '+':
		tok = newToken(token.PLUS, lex.ch)
	case '-':
		tok = newToken(token.MINUS, lex.ch)
	case '!':
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			literal := string(ch) + string(lex.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, lex.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, lex.ch)
	case '/':
		tok = newToken(token.SLASH, lex.ch)
	case '<':
		tok = newToken(token.LT, lex.ch)
	case '>':
		tok = newToken(token.GT, lex.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lex.ch) {
			tok.Literal = lex.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(lex.ch) {
			tok.Type = token.INT
			tok.Literal = lex.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lex.ch)
		}
	}

	// Before returning the token we advance our pointers into the input
	// so when we call NextToken() again the ch field is already updated.
	lex.readChar()
	return tok
}

// The purpose of this helper method is to give us the next character and
// advance our position in the input string.
func (lex *Lexer) readChar() {
	// Check if we've reached the end of input. If yes then assign ch to
	// 0, which essentially means end of file.
	if lex.readPosition >= len(lex.input) {
		lex.ch = 0
	} else {
		// If it's not the end of input then assign ch the next character.
		lex.ch = lex.input[lex.readPosition]
	}

	// Assign the position we've just read to the current position and
	// increment the current reading position by 1.
	lex.position = lex.readPosition
	lex.readPosition++
}

// This function doesn't increment anything, unlike readChar function.
// This function peeks ahead in the input, while not moving around in it,
// so as we'd know what a call to readChar() would return.
func (lex *Lexer) peekChar() byte {
	if lex.readPosition >= len(lex.input) {
		return 0
	}
	return lex.input[lex.readPosition]
}

func (lex *Lexer) readNumber() string {
	position := lex.position
	for isDigit(lex.ch) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

// Method that read in an identifier and advances our lexer's
// positions until it hits a non-letter character.
func (lex *Lexer) readIdentifier() string {
	position := lex.position
	for isLetter(lex.ch) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

func (lex *Lexer) skipWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
		lex.readChar()
	}
}

// Helper function to initialize our various tokens.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Helper function checks whether the given argument is a letter.
// The underscore _ symbol is allowed, too, which means we treat it
// also as a letter and allow it in identifiers and keywords.
// E.g. variable names like my_var can be used.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
