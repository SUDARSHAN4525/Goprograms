// 5. WAP to check if the first string is a substring of another
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Print("Enter the main string: ")
	fmt.Scan(&str2)
	fmt.Print("Enter the substring to check: ")
	fmt.Scan(&str1)

	if strings.Contains(str2, str1) {
		fmt.Println("The first string is a substring of the second string.")
	} else {
		fmt.Println("The first string is not a substring of the second string.")
	}
}
