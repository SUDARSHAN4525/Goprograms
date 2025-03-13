package main

import "fmt"

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	c := &Circle{radius: 5}
	fmt.Println("Area:", c.area())
}
