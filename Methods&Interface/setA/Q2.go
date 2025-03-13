package main

import "fmt"

// Define a method on an integer type
type Number int

// Multiply method that multiplies two numbers
func (n Number) Multiply(other Number) Number {
	return n * other
}

func main() {
	// Create two Number variables
	num1 := Number(5)
	num2 := Number(4)

	// Call the Multiply method
	result := num1.Multiply(num2)

	// Print the result
	fmt.Printf("Multiplication result: %d\n", result)
}
