package main

import "fmt"

// Recursive function to calculate factorial
func factorial(n int) int {
	if n == 0 {
		return 1 // Base case: 0! = 1
	}
	return n * factorial(n-1) // Recursive case
}

func main() {
	number := 5
	result := factorial(number)
	fmt.Printf("Factorial of %d is %d\n", number, result)
}
