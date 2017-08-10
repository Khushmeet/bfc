package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file := os.Args[1]
	code, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}

	compiler := NewCompiler(string(code))
	instructions := compiler.Compile()

	bfm := NewBFMachine(instructions, os.Stdin, os.Stdout)
	bfm.Execute()
}
