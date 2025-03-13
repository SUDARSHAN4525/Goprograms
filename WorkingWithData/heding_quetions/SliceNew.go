package main

import "fmt"

func main() {
	slicePointer := new([]int)

	*slicePointer = []int{10, 20, 30, 40, 50}

	fmt.Println("Slice using new keyword:", *slicePointer)

	*slicePointer = append(*slicePointer, 60, 70)
	fmt.Println("Updated slice:", *slicePointer)
}
