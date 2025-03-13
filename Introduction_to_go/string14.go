package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "GoLangGo"
	char := "o"
	count := strings.Count(str, char)
	fmt.Println("Occurrences of", char, ":", count)
}
