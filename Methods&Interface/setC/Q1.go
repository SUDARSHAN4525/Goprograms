package main

import "fmt"

// Define an interface
type myInterface interface{}

func main() {
	var x myInterface

	// Assigning an integer to x
	x = 42

	// Using type assertion to extract the value from the interface
	if value, ok := x.(int); ok {
		fmt.Printf("x is of type int with value: %d\n", value)
	} else {
		fmt.Println("x is not of type int")
	}

	// Assigning a string to x
	x = "Hello, Go!"

	// Using type assertion to extract the value from the interface
	if value, ok := x.(string); ok {
		fmt.Printf("x is of type string with value: %s\n", value)
	} else {
		fmt.Println("x is not of type string")
	}
}
