package main

import (
	"fmt"
	"time"
)

func main() {
	defer trackTime("main")()

	fmt.Println("Start of main function")
	// Simulate some work
	time.Sleep(2 * time.Second)
	fmt.Println("End of main function")
}

func trackTime(name string) func() {
	start := time.Now()
	fmt.Printf("Start of execution of %s\n", name)
	return func() {
		fmt.Printf("End of execution of %s. Elapsed time: %v\n", name, time.Since(start))
	}
}
