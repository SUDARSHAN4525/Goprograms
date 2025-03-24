package main

import (
	"calculator"
	"fmt"
)

func main() {
	var a, b, choice int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	fmt.Println("1. Add\n2. Subtract\n3. Multiply\n4. Divide")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Result:", calculator.Add(a, b))
	case 2:
		fmt.Println("Result:", calculator.Subtract(a, b))
	case 3:
		fmt.Println("Result:", calculator.Multiply(a, b))
	case 4:
		fmt.Println("Result:", calculator.Divide(a, b))
	default:
		fmt.Println("Invalid choice")
	}
}
