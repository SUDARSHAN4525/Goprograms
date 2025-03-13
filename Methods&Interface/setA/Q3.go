package main

import "fmt"

// Define Author struct
type Author struct {
	Name  string
	Age   int
	Books []string
}

// Method to display information about the author
func (a Author) Show() {
	fmt.Printf("Author: %s\n", a.Name)
	fmt.Printf("Age: %d\n", a.Age)
	fmt.Printf("Books: %v\n", a.Books)
}

func main() {
	// Create an Author instance
	author := Author{
		Name:  "J.K. Rowling",
		Age:   58,
		Books: []string{"Harry Potter and the Sorcerer's Stone", "Harry Potter and the Chamber of Secrets"},
	}

	// Call the Show method on the author instance
	author.Show()
}
