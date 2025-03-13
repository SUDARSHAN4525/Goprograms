// 5. WAP to explain the new function
package main

import "fmt"

func main() {
	ptr := new(int)
	*ptr = 42

	fmt.Println("Value stored at ptr:", *ptr)
	fmt.Println("Address stored in ptr:", ptr)
}
