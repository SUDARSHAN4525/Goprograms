package main

import "fmt"

// Function to find the square of a number
func Square(n int) int {
	return n * n
}

func main() {
	num := 5
	fmt.Printf("Square of %d is %d\n", num, Square(num))
}
