package main

import "fmt"

func incrementByReference(x *int) {
	*x++
	fmt.Println("Inside incrementByReference, x:", *x)
}

func main() {
	num := 20
	fmt.Println("Before incrementByReference, num:", num)
	incrementByReference(&num)
	fmt.Println("After incrementByReference, num:", num)
}
