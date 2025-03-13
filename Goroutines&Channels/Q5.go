package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func process(id int) {
	defer wg.Done() // Decrements the counter when goroutine completes
	fmt.Printf("Processing %d\n", id)
}

func main() {
	for i := 1; i <= 4; i++ {
		wg.Add(1) // Increments the counter
		go process(i)
	}
	wg.Wait() // Waits until counter becomes zero
	fmt.Println("All processes completed.")
}
