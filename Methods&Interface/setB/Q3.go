package main

import "fmt"

// Method to copy elements from one array to another
func copyArray(source [5]int) [5]int {
	var destination [5]int
	for i, v := range source {
		destination[i] = v
	}
	return destination
}

func main() {
	// Initialize an array
	sourceArray := [5]int{1, 2, 3, 4, 5}

	// Call the method to copy array
	copiedArray := copyArray(sourceArray)

	// Print the original and copied arrays
	fmt.Println("Original Array:", sourceArray)
	fmt.Println("Copied Array:", copiedArray)
}
