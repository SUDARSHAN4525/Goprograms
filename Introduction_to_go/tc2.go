package main

import "fmt"

func main() {
	var decimal float64 = 45.67
	var integer int = int(decimal) // Type casting
	fmt.Printf("Float %.2f converted to integer: %d\n", decimal, integer)
}
