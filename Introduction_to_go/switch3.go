package main

import "fmt"

func main() {
	var choice int
	var num1, num2 float64

	fmt.Println("Menu:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	switch choice {
	case 1:
		fmt.Println("Result:", num1+num2)
	case 2:
		fmt.Println("Result:", num1-num2)
	case 3:
		fmt.Println("Result:", num1*num2)
	case 4:
		if num2 != 0 {
			fmt.Println("Result:", num1/num2)
		} else {
			fmt.Println("Error: Division by zero!")
		}
	default:
		fmt.Println("Invalid choice.")
	}
}
