package main

import "fmt"

// Function returning two values: sum and difference
func calculate(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	// Call the function and capture both return values
	sum, diff := calculate(10, 5)

	// Print the results
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
}
