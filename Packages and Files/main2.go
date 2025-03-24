package main

import "fmt"

// Function to add two integers
func Add(a, b int) int {
	return a + b
}

func main() {
	a, b := 5, 7
	fmt.Printf("Sum of %d and %d is %d\n", a, b, Add(a, b))
}
