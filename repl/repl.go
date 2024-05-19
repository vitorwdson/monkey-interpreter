package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/vitorwdson/monkey-interpreter/lexer"
	"github.com/vitorwdson/monkey-interpreter/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			for _, err := range p.Errors() {
				io.WriteString(out, "\t"+err.Error()+"\n")
			}
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}
