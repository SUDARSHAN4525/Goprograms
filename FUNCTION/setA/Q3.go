package main

import "fmt"

// Function to check if a number is palindrome
func isPalindrome(n int) bool {
	original, reversed, remainder := n, 0, 0
	for n != 0 {
		remainder = n % 10
		reversed = reversed*10 + remainder
		n /= 10
	}
	return original == reversed
}

func main() {
	num := 121
	if isPalindrome(num) {
		fmt.Println(num, "is a palindrome.")
	} else {
		fmt.Println(num, "is not a palindrome.")
	}
}
