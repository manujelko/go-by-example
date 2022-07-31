package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` provides access to raw command-line arguments.
	// Not that the first value in this slice is the path to the
	// program, and `os.Args[1:] holds the arguments to the program.`
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
