package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, Go!"
	prefix := "Hello"
	if strings.HasPrefix(str, prefix) {
		fmt.Println("String starts with", prefix)
	} else {
		fmt.Println("String does not start with", prefix)
	}
}
