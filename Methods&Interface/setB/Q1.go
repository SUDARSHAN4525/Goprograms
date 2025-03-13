package main

import "fmt"

// Define the Student struct
type Student struct {
	Name  string
	Age   int
	Grade string
}

// Method to display student details (pointer receiver)
func (s *Student) Show() {
	fmt.Printf("Student Name: %s\n", s.Name)
	fmt.Printf("Age: %d\n", s.Age)
	fmt.Printf("Grade: %s\n", s.Grade)
}

func main() {
	// Create an instance of Student
	student := &Student{
		Name:  "John Doe",
		Age:   20,
		Grade: "A",
	}

	// Call the Show method using pointer receiver
	student.Show()
}
