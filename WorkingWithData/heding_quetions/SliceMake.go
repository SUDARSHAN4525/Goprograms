package main

import "fmt"

func main() {
	// Declare a slice using make
	sliceUsingMake := make([]int, 5) // Creates a slice with length 5 and default values (0)
	fmt.Println("Slice declared using make:", sliceUsingMake)

	// Initialize a slice with values using a slice literal
	sliceWithLiteral := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice initialized with a slice literal:", sliceWithLiteral)

	// Modify elements in the slice declared using make
	sliceUsingMake[0] = 100
	sliceUsingMake[1] = 200
	fmt.Println("Updated slice (using make):", sliceUsingMake)
}
