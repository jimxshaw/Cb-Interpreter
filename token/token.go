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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

// LookupIdent checks the keywords table to see
// whether the given identifier is a keyword.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
