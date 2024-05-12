package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var counter int

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			counter++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			counter++
		}
	}()

	wg.Wait()

	fmt.Println("Valor final do contador:", counter)
}
