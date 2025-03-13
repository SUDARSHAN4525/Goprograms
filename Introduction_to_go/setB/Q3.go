// 3. WAP to print Fibonacci series of n terms
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&n)

	a, b := 0, 1
	fmt.Println("Fibonacci Series:")
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}
