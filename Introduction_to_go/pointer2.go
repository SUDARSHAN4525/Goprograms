package main

import "fmt"

func main() {
	var num int = 10
	var ptr *int = &num

	fmt.Println("Original value of num:", num)
	*ptr = 20
	fmt.Println("Modified value of num using pointer:", num)
}
