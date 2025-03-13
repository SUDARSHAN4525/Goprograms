package main

import "fmt"

func main() {
	var num int = 42
	var floatNum float64 = float64(num) // Type casting
	fmt.Printf("Integer %d converted to float: %.2f\n", num, floatNum)
}
