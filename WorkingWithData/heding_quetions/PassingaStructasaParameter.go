package main

import "fmt"

// Define a structure
type Rectangle struct {
	Length int
	Width  int
}

// Function to calculate the area (pass by value)
func calculateArea(rect Rectangle) int {
	return rect.Length * rect.Width
}

// Function to modify the dimensions (pass by reference)
func modifyDimensions(rect *Rectangle, newLength, newWidth int) {
	rect.Length = newLength
	rect.Width = newWidth
}

func main() {
	// Create a rectangle
	rect := Rectangle{Length: 10, Width: 5}

	// Calculate area (pass by value)
	area := calculateArea(rect)
	fmt.Println("Area of rectangle:", area)

	// Modify dimensions (pass by reference)
	modifyDimensions(&rect, 20, 10)
	fmt.Println("Modified rectangle:", rect)
}
