package main

import (
	"fmt"
	"os"

	"github.com/mrcampbell/gdx/gdx"
)

func main() {
	args := os.Args[1:]
	g := gdx.Goldilox{}
	fmt.Println('\n')
	fmt.Println('\x00')
	fmt.Println('\x00' >= '0')
	fmt.Println('\x00' <= '9')
	var b byte = byte(59)
	fmt.Println(string(b))

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
