package main

import "fmt"

// Function to return multiple values
func getPersonDetails() (string, int, string) {
	name := "Varsha"
	age := 22
	city := "Pune"
	return name, age, city
}

func main() {
	name, age, city := getPersonDetails()
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
}
