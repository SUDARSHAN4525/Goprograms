package main

import "fmt"

func main() {
	var isGoFun, isHard bool = true, false
	fmt.Printf("Go is fun AND hard: %t\n", isGoFun && isHard)
	fmt.Printf("Go is fun OR hard: %t\n", isGoFun || isHard)
}
