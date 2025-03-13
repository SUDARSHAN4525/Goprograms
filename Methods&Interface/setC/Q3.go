package main

import "fmt"

// Define an interface for the printer
type Printer interface {
	Print() string
}

// Define another interface for the scanner
type Scanner interface {
	Scan() string
}

// Define a struct combining Printer and Scanner interfaces (embedded interfaces)
type AllInOnePrinter struct {
	Model string
}

// AllInOnePrinter implements the Print method (Printer interface)
func (a AllInOnePrinter) Print() string {
	return a.Model + " is printing."
}

// AllInOnePrinter implements the Scan method (Scanner interface)
func (a AllInOnePrinter) Scan() string {
	return a.Model + " is scanning."
}

func main() {
	// Create an instance of AllInOnePrinter
	printerScanner := AllInOnePrinter{Model: "HP LaserJet"}

	// Demonstrate the embedded interfaces
	fmt.Println(printerScanner.Print()) // Calls Print method from Printer interface
	fmt.Println(printerScanner.Scan())  // Calls Scan method from Scanner interface
}
