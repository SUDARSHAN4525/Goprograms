package main

import "fmt"

func main() {
	ch := make(chan int) // Create a new channel
	go func() {
		ch <- 10
	}()
	fmt.Println(<-ch) // Receive data from channel
}
