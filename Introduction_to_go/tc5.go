package main

import (
	"fmt"
	"strconv" // For string to number conversion
)

func main() {
	var str string = "3.14"
	floatNum, err := strconv.ParseFloat(str, 64) // Convert string to float
	if err != nil {
		fmt.Println("Error converting string to float:", err)
	} else {
		fmt.Printf("String \"%s\" converted to float: %.2f\n", str, floatNum)
	}
}
