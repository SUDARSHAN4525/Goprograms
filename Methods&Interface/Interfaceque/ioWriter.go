package main

import (
	"fmt"
	"io"
)

// Custom type that implements io.Writer
type CustomWriter struct{}

// Implement the Write method for CustomWriter
func (w *CustomWriter) Write(p []byte) (n int, err error) {
	fmt.Print("Writing Data: ", string(p)) // Print received data
	return len(p), nil                     // Return number of bytes written
}

func main() {
	// Create an instance of CustomWriter
	writer := &CustomWriter{}

	// Data to write
	data := "Hello, Go! Implementing io.Writer.\n"

	// Writing to CustomWriter
	io.WriteString(writer, data) // Using io.WriteString to write a string
}
