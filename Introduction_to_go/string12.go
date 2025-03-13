package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go is fun"
	newStr := strings.Replace(str, "fun", "awesome", 1)
	fmt.Println("Updated string:", newStr)
}
