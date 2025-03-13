package main

import (
	"fmt"
	"strconv" // For number to string conversion
)

func main() {
	var num int = 456
	str := strconv.Itoa(num) // Convert integer to string
	fmt.Printf("Integer %d converted to string: \"%s\"\n", num, str)
}
