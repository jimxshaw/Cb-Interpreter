package token

// TokenType is a string so we can use a variety
// of values for TokenTypes and in turn we can
// distinguish among different types of tokens.
type TokenType string

// Token data structure.
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL signifies a token or character we
	// don't know about.
	ILLEGAL = "ILLEGAL"
	// EOF means "end of file", telling our parser
	// when to stop.
	EOF = "EOF"

	// Identifiers and literals.
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 12345

	// Operators.
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters.
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords.
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
