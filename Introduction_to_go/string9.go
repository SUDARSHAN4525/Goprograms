package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, Go!"
	sub := "Go"
	if strings.Contains(str, sub) {
		fmt.Println("Substring found")
	} else {
		fmt.Println("Substring not found")
	}
}
