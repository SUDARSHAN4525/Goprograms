package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 5, 10
	fmt.Println("Before swapping: x =", x, "y =", y)

	swap(&x, &y)
	fmt.Println("After swapping: x =", x, "y =", y)
}
