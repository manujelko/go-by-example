package main

import (
	"fmt"
	"time"
)

// This function we'll run in a goroutine.
// The done channel will be used to notify another goroutine
// that this function's work is done.
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	// Block until we receive a notification from the worker on
	// the channel.
	<-done
}
