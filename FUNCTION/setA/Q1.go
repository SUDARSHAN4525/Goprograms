package main

import "fmt"

// Function to add two numbers
func add(a, b int) int {
	return a + b
}

func main() {
	num1, num2 := 10, 20
	fmt.Println("Addition of", num1, "and", num2, "is:", add(num1, num2))
}
