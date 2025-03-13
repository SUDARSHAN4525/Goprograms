package main

//not run this program
import (
	"fmt"
)

func main() {
	var value interface{} // Empty interface can hold any type

	// Assign different types to the interface
	value = 100
	checkType(value)

	value = "Hello, Go!"
	checkType(value)

	value = 3.14
	checkType(value)

	value = true
	checkType(value)

	value = []int{1, 2, 3}
	checkType(value)
}

// Function to demonstrate Type Assertion and Type Switch
func checkType(i interface{}) {
	// Type Assertion
	num, ok := i.(int) // Trying to assert that i is an int
	if ok {
		fmt.Println("Type Assertion: The value is an integer:", num)
	} else {
		fmt.Println("Type Assertion failed!")
	}

	// Type Switch
	switch v := i.(type) {
	case int:
		fmt.Println("Type Switch: Integer with value", v)
	case string:
		fmt.Println("Type Switch: String with value", v)
	case float64:
		fmt.Println("Type Switch: Float with value", v)
	case bool:
		fmt.Println("Type Switch: Boolean with value", v)
	default:
		fmt.Println("Type Switch: Unknown type")
	}

	fmt.Println()
}
