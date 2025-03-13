package main

import "fmt"

func main() {
	var num int = 100
	var ptr1 *int = &num
	var ptr2 **int = &ptr1

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num (using ptr1):", ptr1)
	fmt.Println("Address of ptr1 (using ptr2):", ptr2)
	fmt.Println("Value of num (using ptr2):", **ptr2)
}
