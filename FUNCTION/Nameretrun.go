package main

import "fmt"

// Function that returns a string
func getName() string {
	name := "Varsha"
	return name
}

func main() {
	// Call the getName function and store the result in a variable
	var name string = getName()

	// Print the returned name
	fmt.Println("Name returned by the function:", name)
}
