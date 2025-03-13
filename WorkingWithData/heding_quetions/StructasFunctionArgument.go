package main

import "fmt"

// Define a structure
type Circle struct {
	Radius float64
}

// Function to calculate the area of a circle (pass by value)
func calculateArea(c Circle) float64 {
	return 3.14 * c.Radius * c.Radius
}

// Function to modify the radius of a circle (pass by reference)
func modifyRadius(c *Circle, newRadius float64) {
	c.Radius = newRadius
}

func main() {
	// Create a Circle instance
	circle := Circle{Radius: 5}

	// Pass the struct by value to calculateArea
	area := calculateArea(circle)
	fmt.Printf("Area of the circle: %.2f\n", area)

	// Pass the struct by reference to modifyRadius
	modifyRadius(&circle, 10) // Pass the address of the circle
	fmt.Println("Modified Circle:", circle)

	// Recalculate the area with the updated radius
	newArea := calculateArea(circle)
	fmt.Printf("New area of the circle: %.2f\n", newArea)
}
