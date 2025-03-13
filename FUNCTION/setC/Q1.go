package main

import "fmt"

// Function to modify the value (call by value)
func modifyValue(a int) {
	a = 20
	fmt.Println("Inside modifyValue function, a =", a)
}

func main() {
	num := 10
	fmt.Println("Before calling modifyValue, num =", num)
	modifyValue(num)
	fmt.Println("After calling modifyValue, num =", num)
}
