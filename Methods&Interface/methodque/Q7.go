package main

import "fmt"

type Author struct {
	name string
	age  int
}

func (a Author) greet() {
	fmt.Println("Hello,", a.name)
}

func main() {
	author := Author{name: "George Orwell", age: 46}
	author.greet()
}
