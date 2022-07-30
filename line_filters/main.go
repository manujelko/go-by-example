/*
Here's an example line filter in Go that writes a capitalized version of all input text.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Wrapping the unboffered `os.Stdin` with a buffered scanner gives us a
	// convenient `Scan` method that advances the scanner to the next token;
	// which is the next line in the default scanner.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
