package main

import "fmt"

// Function to subtract two integers
func Subtract(a, b int) int {
	return a - b
}

func main() {
	a, b := 10, 4
	fmt.Printf("Subtraction of %d and %d is %d\n", a, b, Subtract(a, b))
}
