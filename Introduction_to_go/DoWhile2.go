package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	i := 1
	for {
		sum += i
		i++
		if i > n {
			break
		}
	}
	fmt.Println("Sum of numbers from 1 to", n, "is:", sum)
}
