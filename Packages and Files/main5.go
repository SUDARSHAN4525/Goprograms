package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Define a struct to match the XML structure
type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
	City string `xml:"city"`
}

func main() {
	// Open the XML file
	file, err := os.Open("data.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode the XML file into the structure
	var person Person
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	// Display the structure
	fmt.Printf("Name: %s\nAge: %d\nCity: %s\n", person.Name, person.Age, person.City)
}
