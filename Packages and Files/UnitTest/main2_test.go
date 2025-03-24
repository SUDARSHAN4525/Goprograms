package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(5, 7)
	expected := 12
	if result != expected {
		t.Errorf("Add(5, 7) = %d; want %d", result, expected)
	}
}
