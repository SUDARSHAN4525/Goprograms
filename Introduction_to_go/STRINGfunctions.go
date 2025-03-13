package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2, substring string
	var choice int

	fmt.Println("Enter the first string:")
	fmt.Scan(&str1)

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Print the string")
		fmt.Println("2. Concatenate two strings")
		fmt.Println("3. Find the length of the string")
		fmt.Println("4. Compare two strings")
		fmt.Println("5. Reverse the string")
		fmt.Println("6. Count vowels in the string")
		fmt.Println("7. Convert string to uppercase")
		fmt.Println("8. Convert string to lowercase")
		fmt.Println("9. Check if a substring exists")
		fmt.Println("10. Split the string")
		fmt.Println("11. Join strings")
		fmt.Println("12. Replace substring")
		fmt.Println("13. Trim spaces from the string")
		fmt.Println("14. Count occurrences of a character")
		fmt.Println("15. Check if string starts with a substring")
		fmt.Println("16. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("The string is:", str1)
		case 2:
			fmt.Println("Enter another string to concatenate:")
			fmt.Scan(&str2)
			fmt.Println("Concatenated string:", str1+" "+str2)
		case 3:
			fmt.Println("Length of the string:", len(str1))
		case 4:
			fmt.Println("Enter another string to compare:")
			fmt.Scan(&str2)
			if str1 == str2 {
				fmt.Println("Strings are equal")
			} else {
				fmt.Println("Strings are not equal")
			}
		case 5:
			reversed := ""
			for i := len(str1) - 1; i >= 0; i-- {
				reversed += string(str1[i])
			}
			fmt.Println("Reversed string:", reversed)
		case 6:
			vowels := 0
			for _, ch := range str1 {
				switch ch {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					vowels++
				}
			}
			fmt.Println("Number of vowels:", vowels)
		case 7:
			fmt.Println("Uppercase string:", strings.ToUpper(str1))
		case 8:
			fmt.Println("Lowercase string:", strings.ToLower(str1))
		case 9:
			fmt.Println("Enter a substring to search:")
			fmt.Scan(&substring)
			if strings.Contains(str1, substring) {
				fmt.Println("Substring found")
			} else {
				fmt.Println("Substring not found")
			}
		case 10:
			parts := strings.Split(str1, " ")
			fmt.Println("Split string:", parts)
		case 11:
			parts := []string{"Go", "is", "fun"}
			fmt.Println("Joined string:", strings.Join(parts, " "))
		case 12:
			fmt.Println("Enter the substring to replace:")
			fmt.Scan(&substring)
			fmt.Println("Enter the new substring:")
			var newSubstring string
			fmt.Scan(&newSubstring)
			fmt.Println("Updated string:", strings.Replace(str1, substring, newSubstring, -1))
		case 13:
			fmt.Println("Trimmed string:", strings.TrimSpace(str1))
		case 14:
			fmt.Println("Enter a character to count:")
			var char string
			fmt.Scan(&char)
			fmt.Println("Occurrences of", char, ":", strings.Count(str1, char))
		case 15:
			fmt.Println("Enter a prefix to check:")
			var prefix string
			fmt.Scan(&prefix)
			if strings.HasPrefix(str1, prefix) {
				fmt.Println("String starts with", prefix)
			} else {
				fmt.Println("String does not start with", prefix)
			}
		case 16:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
