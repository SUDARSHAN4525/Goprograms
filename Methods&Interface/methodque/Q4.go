package main

import "fmt"

type Number int

func (n Number) multiply(factor int) int {
	return int(n) * factor
}

func main() {
	num := Number(10)
	result := num.multiply(5)
	fmt.Println("Result:", result)
}
