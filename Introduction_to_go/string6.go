package main

import "fmt"

func main() {
	str := "Go Programming"
	vowels := 0
	for _, ch := range str {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowels++
		}
	}
	fmt.Println("Number of vowels:", vowels)
}
