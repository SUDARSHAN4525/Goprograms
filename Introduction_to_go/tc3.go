package main

import (
	"fmt"
	"strconv" // For string to number conversion
)

func main() {
	var str string = "123"
	num, err := strconv.Atoi(str) // Convert string to integer
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
	} else {
		fmt.Printf("String \"%s\" converted to integer: %d\n", str, num)
	}
}
