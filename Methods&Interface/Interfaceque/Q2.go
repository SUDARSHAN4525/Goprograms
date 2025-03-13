package main

import "fmt"

//Interface with Multiple Structs
// Define an interface
type Shape interface {
	Area() float64
}

// Struct 1: Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Struct 2: Rectangle
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var s Shape

	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	s = c
	fmt.Println("Circle Area:", s.Area())

	s = r
	fmt.Println("Rectangle Area:", s.Area())
}
