// 5. Anonymous function
package main

import "fmt"

func main() {
	square := func(x int) int {
		return x * x
	}
	fmt.Println("Square of 4:", square(4))
}
