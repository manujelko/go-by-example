package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This range iterates over each element as it's received from the queue.
	// Because we close the channel above, the iteration terminates after receiving the two elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}
