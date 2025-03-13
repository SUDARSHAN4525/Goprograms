package main

import "fmt"

func main() {
	var n, factorial int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	factorial = 1
	i := 1
	for i <= n {
		factorial *= i
		i++
	}
	fmt.Println("Factorial of", n, "is:", factorial)
}
