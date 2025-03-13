package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func process(id int) {
	defer wg.Done() // Decrement the counter when Goroutine finishes
	fmt.Printf("Processing ID: %d\n", id)
	time.Sleep(time.Second)
}

func main() {
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the counter for each Goroutine
		go process(i)
	}

	wg.Wait() // Wait until all Goroutines finish
	fmt.Println("All processes completed.")
}
