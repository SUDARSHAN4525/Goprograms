package main

import "fmt"

func main() {
	var num, largest int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	for {
		digit := num % 10
		if digit > largest {
			largest = digit
		}
		num /= 10
		if num == 0 {
			break
		}
	}
	fmt.Println("The largest digit is:", largest)
}
