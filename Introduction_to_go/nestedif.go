package main

import "fmt"

func main() {
	var marks int
	fmt.Print("Enter the marks (out of 100): ")
	fmt.Scan(&marks)

	if marks >= 90 {
		fmt.Println("Grade: A")
	} else {
		if marks >= 75 {
			fmt.Println("Grade: B")
		} else {
			if marks >= 50 {
				fmt.Println("Grade: C")
			} else {
				fmt.Println("Grade: F")
			}
		}
	}
}
