package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go is fun"
	parts := strings.Split(str, " ")
	fmt.Println("Split string:", parts)
}
