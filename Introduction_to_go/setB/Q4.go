// 4. WAP to illustrate pointer to pointer concept
package main

import "fmt"

func main() {
	var x int = 10
	var ptr *int = &x
	var pptr **int = &ptr

	fmt.Println("Value of x:", x)
	fmt.Println("Value at ptr:", *ptr)
	fmt.Println("Value at pptr:", **pptr)
}
