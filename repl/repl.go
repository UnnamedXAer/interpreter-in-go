package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/unnamedxaer/interpreter-in-go/evaluator"
	"github.com/unnamedxaer/interpreter-in-go/lexer"
	"github.com/unnamedxaer/interpreter-in-go/object"
	"github.com/unnamedxaer/interpreter-in-go/parser"
)

const MONKEY_FACE = `
`

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}

}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ren into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
