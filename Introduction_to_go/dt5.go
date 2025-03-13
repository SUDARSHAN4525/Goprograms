package main

import "fmt"

func main() {
	var i int = 42
	var f float64 = float64(i)
	fmt.Printf("Integer %d as float: %.2f\n", i, f)
}
