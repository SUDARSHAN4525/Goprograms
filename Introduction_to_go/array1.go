package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Array elements are:")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
