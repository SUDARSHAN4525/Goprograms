package main

import "fmt"

func main() {
	var unusedVar int = 10
	fmt.Println("This program will not compile because 'unusedVar' is declared but not used.")
}
