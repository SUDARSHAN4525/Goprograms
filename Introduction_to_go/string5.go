package main

import "fmt"

func main() {
	str := "GoLang"
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	fmt.Println("Reversed string:", reversed)
}
