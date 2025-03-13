package main

import "fmt"

// Define an interface
type Printer interface {
	Print()
}

// Struct with a value receiver
type Document struct {
	text string
}

// Method with a **value receiver** (works only with values)
func (d Document) Print() {
	fmt.Println("Document:", d.text)
}

// Struct with a pointer receiver
type Report struct {
	data string
}

// Method with a **pointer receiver** (works with pointers)
func (r *Report) Print() {
	fmt.Println("Report:", r.data)
}

func main() {
	// Using Document (value receiver)
	doc := Document{"Go Interfaces"}
	var p1 Printer = doc // ✅ Allowed (value works with value receiver)
	p1.Print()

	// Using Report (pointer receiver)
	rep := Report{"Annual Report"}
	// var p2 Printer = rep // ❌ Not allowed (Report has pointer receiver)

	var p2 Printer = &rep // ✅ Allowed (pointer works with pointer receiver)
	p2.Print()
}
