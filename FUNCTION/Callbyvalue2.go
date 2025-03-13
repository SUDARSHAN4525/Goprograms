package main

import "fmt"

func doubleValue(a int) {
	a = a * 2
	fmt.Println("Inside doubleValue, a:", a)
}

func main() {
	value := 15
	fmt.Println("Before doubleValue, value:", value)
	doubleValue(value)
	fmt.Println("After doubleValue, value:", value)
}
