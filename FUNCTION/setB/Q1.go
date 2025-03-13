package main

import "fmt"

// Function to swap two numbers using pointers
func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 5, 10
	fmt.Println("Before swapping: x =", x, "y =", y)
	swap(&x, &y)
	fmt.Println("After swapping: x =", x, "y =", y)
}
