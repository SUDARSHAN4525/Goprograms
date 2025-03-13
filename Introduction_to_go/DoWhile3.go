package main

import "fmt"

func main() {
	var n, i int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	i = 1
	for {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
		i++
		if i > 10 {
			break
		}
	}
}
