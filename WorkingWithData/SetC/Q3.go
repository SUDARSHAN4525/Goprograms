package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}

	// Append
	slice = append(slice, 60, 70)
	fmt.Println("After Append:", slice)

	// Remove
	index := 2 // Remove element at index 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("After Remove:", slice)

	// Copy
	copiedSlice := make([]int, len(slice))
	copy(copiedSlice, slice)
	fmt.Println("Copied Slice:", copiedSlice)
}
