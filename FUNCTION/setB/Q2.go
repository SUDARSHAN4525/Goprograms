package main

import "fmt"

// Function with named return variables
func calculate(a, b int) (sum, diff int) {
	sum = a + b
	diff = a - b
	return
}

func main() {
	sum, diff := calculate(10, 5)
	fmt.Println("Sum:", sum, "Difference:", diff)
}
