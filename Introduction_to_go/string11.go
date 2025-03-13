package main

import (
	"fmt"
	"strings"
)

func main() {
	parts := []string{"Go", "is", "awesome"}
	result := strings.Join(parts, " ")
	fmt.Println("Joined string:", result)
}
