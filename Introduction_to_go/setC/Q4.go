// 4. WAP to check if a number is a single digit
package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num >= -9 && num <= 9 {
		fmt.Println("The number is a single digit.")
	} else {
		fmt.Println("The number is not a single digit.")
	}
}
