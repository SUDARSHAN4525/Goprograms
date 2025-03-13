package main

import "fmt"

type Author struct {
	name string
	age  int
}

func (a Author) details() {
	fmt.Println("Author:", a.name)
	fmt.Println("Age:", a.age)
}

func main() {
	author := Author{name: "J.K. Rowling", age: 58}
	author.details()
}
