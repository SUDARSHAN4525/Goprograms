package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{34, 7, 23, 32, 5, 62}
	sort.Ints(arr)
	fmt.Println("Sorted Array:", arr)
}
