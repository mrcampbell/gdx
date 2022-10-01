package gdx

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Debug(message string, opt ...any) {
	fmt.Println(message, opt)
}

var DEBUG = len(os.Getenv("DEBUG")) > 0

func printYellow(format string, opt ...interface{}) {
	if DEBUG {
		color.Yellow(format, opt)
	}
}

func printPurple(format string, opt ...interface{}) {
	if DEBUG {
		color.Magenta(format, opt)
	}
}

func printGreen(format string, opt ...interface{}) {
	if DEBUG {
		color.Green(format, opt...)
	}
}

func printRed(format string, opt ...interface{}) {
	if DEBUG {
		color.Red(format, opt...)
	}
}

func printBlue(format string, opt ...interface{}) {
	if DEBUG {
		color.Blue(format, opt)
	}
}

func printGrey(format string, opt ...interface{}) {
	if DEBUG {
		fmt.Printf(format, opt...)
		fmt.Println()
	}
}
