// 3. Function with parameters and a return value
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(5, 3)
	fmt.Println("Sum:", result)
}
