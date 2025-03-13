package main

import "fmt"

func main() {
	var num, reverse int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	for num > 0 {
		remainder := num % 10
		reverse = reverse*10 + remainder
		num /= 10
	}
	fmt.Println("Reversed number is:", reverse)
}
