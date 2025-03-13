package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Task %d started\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Task %d finished\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All tasks completed.")
}
