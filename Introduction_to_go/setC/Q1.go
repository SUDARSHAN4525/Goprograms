// SET C

// 1. WAP to concatenate two strings using pointers
package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := " World"

	result := *(&str1) + *(&str2)
	fmt.Println("Concatenated string:", result)
}
