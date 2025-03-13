package main

import "fmt"

func main() {
	// Declaring constants
	const pi float64 = 3.14159
	const greeting string = "Hello, World!"

	// Declaring variables
	var radius int = 5
	var area float64

	// Calculating area of a circle
	area = pi * float64(radius*radius) // Type casting required for multiplication

	// Printing results
	fmt.Println(greeting)
	fmt.Printf("The area of a circle with radius %d is %.2f\n", radius, area)
}
