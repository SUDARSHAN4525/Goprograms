package main

import "fmt"

type Student struct {
	RollNo  int
	Name    string
	Marks   [3]float64
	Total   float64
	Average float64
}

func main() {
	var n int
	fmt.Print("Enter the number of students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter details for student %d (RollNo, Name, Marks1, Marks2, Marks3):\n", i+1)
		fmt.Scan(&students[i].RollNo, &students[i].Name, &students[i].Marks[0], &students[i].Marks[1], &students[i].Marks[2])
		students[i].Total = students[i].Marks[0] + students[i].Marks[1] + students[i].Marks[2]
		students[i].Average = students[i].Total / 3
	}

	fmt.Println("Student Details:")
	for _, student := range students {
		fmt.Printf("RollNo: %d, Name: %s, Total: %.2f, Average: %.2f\n", student.RollNo, student.Name, student.Total, student.Average)
	}
}
