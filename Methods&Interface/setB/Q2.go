package main

import "fmt"

// Define an empty interface
type myInterface interface{}

func main() {
	var x myInterface

	// Assign different types to x and use type switch
	x = 42
	switch v := x.(type) {
	case int:
		fmt.Printf("x is an int with value: %d\n", v)
	case string:
		fmt.Printf("x is a string with value: %s\n", v)
	default:
		fmt.Println("Unknown type")
	}

	x = "hello"
	switch v := x.(type) {
	case int:
		fmt.Printf("x is an int with value: %d\n", v)
	case string:
		fmt.Printf("x is a string with value: %s\n", v)
	default:
		fmt.Println("Unknown type")
	}
}
