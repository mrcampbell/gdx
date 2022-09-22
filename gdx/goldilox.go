package gdx

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/kr/pretty"
)

type Goldilox struct {
	hasError bool
	scanner  Scanner
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (g *Goldilox) run(input string) error {
	g.scanner = *NewScanner(input)
	tokens, err := g.scanner.ScanTokens()
	pretty.Println(tokens, err)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (g *Goldilox) RunFile(path string) {
	b, err := os.ReadFile(path)

	check(err)

	g.run(string(b))

	if g.hasError {
		os.Exit(65)
	}
}

func (g *Goldilox) RunPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			fmt.Println("\nGoodbye!")
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		err = g.run(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		g.hasError = false
	}
}

func (g *Goldilox) RaiseError(line int, message string) {
	g.Report(line, "", message)
}

func (g *Goldilox) Report(line int, where string, message string) {
	g.hasError = true
	fmt.Printf("[line %d] Error%s: %s", line, where, message)
}
