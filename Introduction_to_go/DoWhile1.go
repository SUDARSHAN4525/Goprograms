package main

import "fmt"

func main() {
	num := 1
	for {
		fmt.Println(num)
		num++
		if num > 10 {
			break
		}
	}
}
