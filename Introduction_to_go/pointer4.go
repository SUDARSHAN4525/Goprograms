package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	ptr := &arr[0]

	fmt.Println("Array elements accessed using pointer:")
	for i := 0; i < len(arr); i++ {
		fmt.Println(*ptr)
		ptr = &arr[i+1] // Move to the next element
	}
}
