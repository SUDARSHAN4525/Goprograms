package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if number > 0 {
		fmt.Println("The number is positive.")
	} else {
		fmt.Println("The number is negative.")
	}
}
