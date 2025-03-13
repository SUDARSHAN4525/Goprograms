// 1. Go program to create a buffered channel and find capacity and length
package main

import "fmt"

func main() {
	// Create a buffered channel with capacity of 5
	ch := make(chan int, 5)

	// Store values in the channel
	ch <- 10
	ch <- 20
	ch <- 30

	// Print capacity and length
	fmt.Println("Channel capacity:", cap(ch))
	fmt.Println("Channel length before reading:", len(ch))

	// Read values from channel
	fmt.Println("Reading from channel:")
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Print modified length after reading
	fmt.Println("Channel length after reading:", len(ch))
}
