package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "test.txt" // File name to read info

	// Get file information
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print file information
	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size(), "bytes")
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
}
