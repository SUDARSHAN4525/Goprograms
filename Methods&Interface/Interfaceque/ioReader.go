package main

import (
	"fmt"
	"io"
)

// Custom type that implements io.Reader
type CustomReader struct {
	Data  string
	Index int
}

// Implement the Read method for CustomReader
func (r *CustomReader) Read(p []byte) (n int, err error) {
	if r.Index >= len(r.Data) {
		return 0, io.EOF // End of data
	}

	n = copy(p, r.Data[r.Index:]) // Copy data to the buffer
	r.Index += n                  // Move index forward

	return n, nil
}

func main() {
	// Create an instance of CustomReader
	reader := &CustomReader{Data: "Hello, Go! Implementing io.Reader."}

	// Buffer to store read data
	buffer := make([]byte, 8) // Read 8 bytes at a time

	fmt.Println("Reading Data:")
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Print(string(buffer[:n])) // Print read data
	}
}
