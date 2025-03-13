package main

import "fmt"

func main() {
	var day int

	fmt.Println("Menu:")
	fmt.Println("1. Monday")
	fmt.Println("2. Tuesday")
	fmt.Println("3. Wednesday")
	fmt.Println("4. Thursday")
	fmt.Println("5. Friday")
	fmt.Println("6. Saturday")
	fmt.Println("7. Sunday")
	fmt.Print("Enter the number for the day: ")
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("Today is Monday.")
	case 2:
		fmt.Println("Today is Tuesday.")
	case 3:
		fmt.Println("Today is Wednesday.")
	case 4:
		fmt.Println("Today is Thursday.")
	case 5:
		fmt.Println("Today is Friday.")
	case 6:
		fmt.Println("Today is Saturday.")
	case 7:
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Invalid day number.")
	}
}
