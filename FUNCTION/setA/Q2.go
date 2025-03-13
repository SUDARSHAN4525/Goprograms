package main

import "fmt"

// Recursive function to calculate the sum of digits
func sumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + sumOfDigits(n/10)
}

func main() {
	num := 1234
	fmt.Println("Sum of digits of", num, "is:", sumOfDigits(num))
}
