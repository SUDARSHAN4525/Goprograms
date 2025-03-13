package main

import "fmt"

// Directly recursive function to print numbers from n to 1
func printNumbers(n int) {
	if n == 0 {
		return // Base case to stop recursion
	}
	fmt.Println(n)
	printNumbers(n - 1) // Recursive call
}

func main() {
	fmt.Println("Numbers from 5 to 1:")
	printNumbers(5)
}
