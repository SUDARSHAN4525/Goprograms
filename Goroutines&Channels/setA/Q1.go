// 1. Go program using Goroutine and Channel to print sum of squares and cubes of individual digits
package main

import (
	"fmt"
)

func sumOfSquaresAndCubes(num int, ch chan int) {
	sumSquares, sumCubes := 0, 0
	for num > 0 {
		digit := num % 10
		sumSquares += digit * digit
		sumCubes += digit * digit * digit
		num /= 10
	}
	ch <- sumSquares
	ch <- sumCubes
}

func main() {
	num := 123
	ch := make(chan int)

	go sumOfSquaresAndCubes(num, ch)

	squares := <-ch
	cubes := <-ch

	fmt.Printf("Sum of squares = %d\n", squares)
	fmt.Printf("Sum of cubes = %d\n", cubes)
	fmt.Printf("Final sum of squares and cubes = %d\n", squares+cubes)
}
