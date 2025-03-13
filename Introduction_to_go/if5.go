package main

import "fmt"

func main() {
	var char string
	fmt.Print("Enter a character: ")
	fmt.Scan(&char)

	if char == "a" || char == "e" || char == "i" || char == "o" || char == "u" ||
		char == "A" || char == "E" || char == "I" || char == "O" || char == "U" {
		fmt.Println("The character is a vowel.")
	}
}
