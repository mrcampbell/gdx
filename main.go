package main

import (
	"fmt"
	"os"

	"github.com/mrcampbell/gdx/gdx"
)

func main() {
	args := os.Args[1:]
	g := gdx.Goldilox{}

	fmt.Println('&')

	if len(args) > 1 {
		fmt.Println("Usage: gdx [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		g.RunFile(args[0])
		fmt.Printf("%+v\n", g)
	} else {
		g.RunPrompt()
	}
}
