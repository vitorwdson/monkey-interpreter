package main

import (
	"fmt"
	"os"

	"github.com/vitorwdson/monkey-interpreter/repl"
)

func main() {
	fmt.Println("Starting MonkeyLang interpreter...")
	fmt.Println()
	repl.Start(os.Stdin, os.Stdout)
}
