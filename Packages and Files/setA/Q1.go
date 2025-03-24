package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Marks int
}

func main() {
	students := []Student{
		{"Amit", 85},
		{"Varsha", 92},
		{"Ravi", 78},
		{"Priya", 88},
	}

	// Sorting students based on marks
	sort.Slice(students, func(i, j int) bool {
		return students[i].Marks > students[j].Marks
	})

	fmt.Println("Sorted Students by Marks:")
	for _, s := range students {
		fmt.Printf("%s: %d\n", s.Name, s.Marks)
	}
}
