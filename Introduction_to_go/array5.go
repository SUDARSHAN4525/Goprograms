package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	var key int
	found := false

	fmt.Print("Enter the element to search: ")
	fmt.Scan(&key)

	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			fmt.Printf("Element %d found at index %d\n", key, i)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Element not found in the array.")
	}
}
