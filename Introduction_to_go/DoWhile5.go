package main

import "fmt"

func main() {
	var num, count int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	for {
		count++
		num /= 10
		if num == 0 {
			break
		}
	}
	fmt.Println("The number of digits is:", count)
}
