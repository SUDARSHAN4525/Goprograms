package main

import "fmt"

func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func main() {
	result1 := applyOperation(10, 5, add)
	result2 := applyOperation(10, 5, subtract)

	fmt.Println("Addition result:", result1)
	fmt.Println("Subtraction result:", result2)
}
