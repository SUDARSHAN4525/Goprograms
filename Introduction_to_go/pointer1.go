package main

import "fmt"

func main() {
	var num int = 42
	var ptr *int = &num

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", ptr)
	fmt.Println("Value at the address:", *ptr)
}
