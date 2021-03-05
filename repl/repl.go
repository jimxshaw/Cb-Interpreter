package repl

import (
	"bufio"
	"cb-interpreter/lexer"
	"cb-interpreter/token"
	"fmt"
	"io"
)

const PROMPT = ">>"

// Start reads from the input source until encountering
// a newline, take the read line and parse it to an
// instance of our lexer and finally print all tokens
// the lexer gives us until EOF.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
