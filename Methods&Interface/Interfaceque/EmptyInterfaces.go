package main

import "fmt"

func printValue(value interface{}) {
	fmt.Println("Value:", value)
}

func main() {
	printValue(42)           // Integer
	printValue("Hello, Go!") // String
	printValue(3.14)         // Float
	printValue(true)         // Boolean
}
