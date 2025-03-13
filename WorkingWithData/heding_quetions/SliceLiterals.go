package main

import (
	"fmt"
)

func main() {
	// 1. Create an empty slice
	var emptySlice []int
	fmt.Println("Empty slice:", emptySlice)

	// 2. Declare a slice array
	var sliceArray []int
	fmt.Println("Declared slice array:", sliceArray)

	// 3. Initialize a slice with values using a slice literal
	sliceLiteral := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice initialized with literal:", sliceLiteral)

	// 4. Perform 5 basic operations on slices

	// a. Append elements to the slice
	sliceLiteral = append(sliceLiteral, 60, 70)
	fmt.Println("After appending elements:", sliceLiteral)

	// b. Find the length of the slice
	fmt.Println("Length of the slice:", len(sliceLiteral))

	// c. Find the capacity of the slice
	fmt.Println("Capacity of the slice:", cap(sliceLiteral))

	// d. Slicing a slice (sub-slicing)
	subSlice := sliceLiteral[2:5]
	fmt.Println("Sub-slice from index 2 to 4:", subSlice)

	// e. Iterating through the slice elements
	fmt.Println("Iterating through slice elements:")
	for i, value := range sliceLiteral {
		fmt.Printf("Index %d: %d\n", i, value)
	}
}
