package main

import "testing"

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{10, 4, 6},
		{5, 3, 2},
		{7, 9, -2},
		{0, 5, -5},
	}

	for _, test := range tests {
		result := Subtract(test.a, test.b)
		if result != test.expected {
			t.Errorf("Subtract(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}
