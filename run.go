package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	file := os.Args[1]
	code, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}

	bfm := NewBFMachine(string(code), os.Stdin, os.Stdout)
	bfm.Execute()
}
