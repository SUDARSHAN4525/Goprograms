// 2. Go program to execute 5 Goroutines simultaneously
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generateNumbers(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
		time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go generateNumbers(i, &wg)
	}

	wg.Wait()
}
