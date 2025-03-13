package main

import "fmt"

type Rectangle struct {
	length, width int
}

func (r Rectangle) area() int {
	return r.length * r.width
}

func (r *Rectangle) setDimensions(length, width int) {
	r.length = length
	r.width = width
}

func main() {
	rect := Rectangle{length: 5, width: 4}
	rect.setDimensions(10, 6)
	fmt.Println("Area:", rect.area())
}
