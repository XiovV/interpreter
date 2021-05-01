package main

import (
	"github.com/XiovV/interpreter/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
