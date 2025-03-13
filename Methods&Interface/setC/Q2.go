package main

import (
	"fmt"
	"sort"
)

// Define the Student struct
type Student struct {
	RollNo     int
	Name       string
	Percentage float64
}

// Method to display student information in descending order of percentage
func (students *[]Student) SortAndDisplay() {
	// Sort the students slice by percentage in descending order
	sort.SliceStable(*students, func(i, j int) bool {
		return (*students)[i].Percentage > (*students)[j].Percentage
	})

	// Display sorted student information
	for _, student := range *students {
		fmt.Printf("RollNo: %d, Name: %s, Percentage: %.2f%%\n", student.RollNo, student.Name, student.Percentage)
	}
}

func main() {
	// Creating a slice of students
	students := []Student{
		{RollNo: 1, Name: "Alice", Percentage: 89.5},
		{RollNo: 2, Name: "Bob", Percentage: 92.3},
		{RollNo: 3, Name: "Charlie", Percentage: 85.7},
		{RollNo: 4, Name: "David", Percentage: 76.2},
	}

	// Calling SortAndDisplay method to display students in descending order of percentage
	(&students).SortAndDisplay()
}
