package main

import "fmt"

func main() {
	var choice int
	var temp float64

	fmt.Println("Menu:")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter temperature in Celsius: ")
		fmt.Scan(&temp)
		fmt.Println("Temperature in Fahrenheit:", (temp*9/5)+32)
	case 2:
		fmt.Print("Enter temperature in Fahrenheit: ")
		fmt.Scan(&temp)
		fmt.Println("Temperature in Celsius:", (temp-32)*5/9)
	default:
		fmt.Println("Invalid choice.")
	}
}
