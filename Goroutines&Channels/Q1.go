package main

import (
	"fmt"
	"time"
)

func display(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go display("Hello") // New Goroutine
	display("World")    // Main Goroutine
}
