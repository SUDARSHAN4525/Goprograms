package main

import (
	"encoding/xml"
	"testing"
)

// Test data for XML parsing
const xmlData = `
<person>
    <name>Varsha</name>
    <age>21</age>
    <city>Pune</city>
</person>
`

func TestXMLParsing(t *testing.T) {
	var person Person
	err := xml.Unmarshal([]byte(xmlData), &person)
	if err != nil {
		t.Fatalf("Failed to parse XML: %v", err)
	}

	if person.Name != "Varsha" || person.Age != 21 || person.City != "Pune" {
		t.Errorf("Unexpected parsed data: %+v", person)
	}
}
