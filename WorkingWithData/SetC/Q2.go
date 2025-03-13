package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {
	var n int
	fmt.Print("Enter the number of employees: ")
	fmt.Scan(&n)

	employees := make([]Employee, n)
	var maxSalary float64
	var maxEmployee Employee

	for i := 0; i < n; i++ {
		fmt.Printf("Enter details for employee %d (ID, Name, Salary):\n", i+1)
		fmt.Scan(&employees[i].ID, &employees[i].Name, &employees[i].Salary)

		if employees[i].Salary > maxSalary {
			maxSalary = employees[i].Salary
			maxEmployee = employees[i]
		}
	}

	fmt.Printf("Employee with the highest salary: %+v\n", maxEmployee)
}
