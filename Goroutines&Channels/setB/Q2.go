// 2. Go program to read and write Fibonacci series to a channel
package main

import "fmt"

func generateFibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	n := 10

	go generateFibonacci(n, ch)

	fmt.Println("Fibonacci series:")
	for num := range ch {
		fmt.Println(num)
	}
}
