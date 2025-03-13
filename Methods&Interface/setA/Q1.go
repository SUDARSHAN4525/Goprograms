package main

import (
	"fmt"
	"math"
)

// Shape interface with methods Area and Perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle struct
type Circle struct {
	radius float64
}

// Rectangle struct
type Rectangle struct {
	length, width float64
}

// Circle method to calculate area
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Circle method to calculate perimeter
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Rectangle method to calculate area
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Rectangle method to calculate perimeter
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{radius: 5}
	rectangle := Rectangle{length: 4, width: 6}

	// Print results for Circle
	fmt.Printf("Circle Area: %.2f\n", circle.Area())
	fmt.Printf("Circle Perimeter: %.2f\n", circle.Perimeter())

	// Print results for Rectangle
	fmt.Printf("Rectangle Area: %.2f\n", rectangle.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n", rectangle.Perimeter())
}
