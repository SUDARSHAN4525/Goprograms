package main

import "fmt"

// Function A calls Function B
func functionA(n int) {
	if n <= 0 {
		return // Base case to stop recursion
	}
	fmt.Println("Function A:", n)
	functionB(n - 1) // Call to Function B
}

// Function B calls Function A
func functionB(n int) {
	if n <= 0 {
		return // Base case to stop recursion
	}
	fmt.Println("Function B:", n)
	functionA(n - 1) // Call to Function A
}

func main() {
	fmt.Println("Indirect Recursion Output:")
	functionA(5)
}
