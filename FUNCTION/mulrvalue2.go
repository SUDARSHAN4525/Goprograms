package main

import "fmt"

// Function returning a boolean and a string
func validateAge(age int) (bool, string) {
	if age >= 18 {
		return true, "Eligible for voting"
	}
	return false, "Not eligible for voting"
}

func main() {
	// Call the function and capture both return values
	status, message := validateAge(16)

	// Print the results
	fmt.Println("Status:", status)
	fmt.Println("Message:", message)
}
