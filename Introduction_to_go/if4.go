package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if number%5 == 0 {
		fmt.Println("The number is divisible by 5.")
	}
}
