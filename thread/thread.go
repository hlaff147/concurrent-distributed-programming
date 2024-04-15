package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	fmt.Println(message)
}

func main() {
	// This program demonstrates the use of goroutines to print messages concurrently.
	// The main function also creates a goroutine to execute its tasks. If the main goroutine
	// finishes before the other goroutines are scheduled, they might not get a chance to run.
	
	go printMessage("Hello from goroutine 1")
	go printMessage("Hello from goroutine 2")

	time.Sleep(time.Second)

	fmt.Println("Hello, world!")
}
