package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	fmt.Println("Capacity:", cap(ch)) // Output: 3
}
