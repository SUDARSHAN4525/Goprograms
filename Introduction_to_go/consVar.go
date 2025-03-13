package main

import "fmt"

func main() {
	// Constants for enumeration
	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)

	// Variables for initialization
	var day int = Tuesday
	var message string

	// Conditional message based on the day
	if day == Monday {
		message = "Start of the work week!"
	} else if day == Tuesday {
		message = "Keep going, it's only Tuesday."
	} else {
		message = "Mid-week vibes!"
	}

	// Printing message
	fmt.Printf("Day %d: %s\n", day, message)
}
