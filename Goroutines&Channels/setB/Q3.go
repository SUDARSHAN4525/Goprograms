// 3. Go program to create and close a channel using for range loop
package main

import "fmt"

func sendData(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go sendData(ch)

	fmt.Println("Reading from channel:")
	for value := range ch {
		fmt.Println(value)
	}
}
