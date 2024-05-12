package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex

	// Goroutine 1 tenta adquirir o mutex
	go func() {
		fmt.Println("Goroutine 1 tentando adquirir o mutex")
		mutex.Lock()
		fmt.Println("Goroutine 1 adquiriu o mutex")
	}()

	// Goroutine 2 tenta adquirir o mutex
	go func() {
		fmt.Println("Goroutine 2 tentando adquirir o mutex")
		mutex.Lock()
		fmt.Println("Goroutine 2 adquiriu o mutex")
	}()

	select {}
}
