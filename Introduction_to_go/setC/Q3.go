// 3. WAP to accept user choice and perform arithmetic operations
package main

import "fmt"

func main() {
	var a, b, choice int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	fmt.Println("Choose operation: 1-Add, 2-Subtract, 3-Multiply, 4-Divide")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Result:", a+b)
	case 2:
		fmt.Println("Result:", a-b)
	case 3:
		fmt.Println("Result:", a*b)
	case 4:
		if b != 0 {
			fmt.Println("Result:", a/b)
		} else {
			fmt.Println("Division by zero error.")
		}
	default:
		fmt.Println("Invalid choice.")
	}
}
