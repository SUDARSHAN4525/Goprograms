// 3. Go program to check even or odd numbers using Goroutine and Channel
package main

import (
	"fmt"
)

func checkEven(ch chan int) {
	for n := range ch {
		if n%2 == 0 {
			fmt.Printf("Even: %d\n", n)
		}
	}
}

func checkOdd(ch chan int) {
	for n := range ch {
		if n%2 != 0 {
			fmt.Printf("Odd: %d\n", n)
		}
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenCh := make(chan int)
	oddCh := make(chan int)

	go checkEven(evenCh)
	go checkOdd(oddCh)

	for _, num := range numbers {
		if num%2 == 0 {
			evenCh <- num
		} else {
			oddCh <- num
		}
	}

	close(evenCh)
	close(oddCh)
}
