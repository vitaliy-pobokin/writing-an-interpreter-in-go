package repl

import (
	"bufio"
	"fmt"
	"github.com/vitaliy-pobokin/writing-an-interpreter-in-go/lexer"
	"github.com/vitaliy-pobokin/writing-an-interpreter-in-go/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
