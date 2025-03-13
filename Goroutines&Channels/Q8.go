package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func process(data int) {
	defer wg.Done()
	fmt.Printf("Processing data: %d\n", data)
	time.Sleep(time.Millisecond * 500)
}

func main() {
	data := []int{10, 20, 30, 40}

	for _, value := range data {
		wg.Add(1)
		go process(value)
	}

	wg.Wait()
	fmt.Println("All data processed.")
}
