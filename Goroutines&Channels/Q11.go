package main

import "fmt"

func main() {
	ch := make(chan int) // Unbuffered channel
	go func() {
		ch <- 5
	}()
	fmt.Println(<-ch) // Output: 5
}
