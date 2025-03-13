package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("Anonymous:", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	time.Sleep(3 * time.Second) // Main Goroutine waits to let other Goroutine finish
}
