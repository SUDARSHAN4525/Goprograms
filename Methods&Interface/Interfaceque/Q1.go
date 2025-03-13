package main

import "fmt"

//Basic Interface Implementation
// Define an interface
type Speaker interface {
	Speak() string
}

// Define a struct that implements the interface
type Person struct {
	Name string
}

// Implement the Speak method for Person
func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

func main() {
	var s Speaker // Declare an interface variable

	p := Person{"Varsha"}
	s = p // Assigning struct value to the interface

	fmt.Println(s.Speak()) // Output: Hello, my name is Varsha
}
