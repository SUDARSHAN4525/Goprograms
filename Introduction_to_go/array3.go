package main

import "fmt"

func main() {
	arr := [5]int{10, 25, 5, 40, 15}
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Println("Maximum element in the array is:", max)
}
