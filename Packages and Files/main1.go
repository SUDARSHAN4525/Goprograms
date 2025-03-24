package main

import (
	"fmt"
	"rectangle"
)

func main() {
	var length, width float64

	fmt.Print("Enter length of rectangle: ")
	fmt.Scan(&length)
	fmt.Print("Enter width of rectangle: ")
	fmt.Scan(&width)

	area := rectangle.Area(length, width)
	fmt.Printf("Area of Rectangle: %.2f\n", area)
}
