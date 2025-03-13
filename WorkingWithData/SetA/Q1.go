package main

import "fmt"

func main() {
	arr := []int{34, 7, 23, 32, 5, 62}
	smallest, largest := arr[0], arr[0]

	for _, num := range arr {
		if num < smallest {
			smallest = num
		}
		if num > largest {
			largest = num
		}
	}
	fmt.Printf("Smallest: %d, Largest: %d\n", smallest, largest)
}
