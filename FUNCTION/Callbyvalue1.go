package main

import "fmt"

func increment(x int) {
	x++
	fmt.Println("Inside increment, x:", x)
}

func main() {
	num := 10
	fmt.Println("Before increment, num:", num)
	increment(num)
	fmt.Println("After increment, num:", num)
}
