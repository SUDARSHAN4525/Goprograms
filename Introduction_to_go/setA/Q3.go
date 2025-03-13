// 3. WAP to swap numbers without a temporary variable
package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("After swapping, a:", a, "b:", b)
}
