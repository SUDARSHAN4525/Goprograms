package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	fmt.Println("Length:", len(ch)) // Output: 2
}
