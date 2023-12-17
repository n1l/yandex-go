package main

import "testing"

func TestAbs(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{input: -1.1, expected: 1.1},
		{input: -1, expected: 1},
		{input: -0, expected: 0},
		{input: 0, expected: 0},
		{input: 1, expected: 1},
		{input: 1.1, expected: 1.1},
	}

	for _, tc := range testCases {
		result := Abs(tc.input)
		if result != tc.expected {
			t.Errorf("Abs() = %f, want %f", result, tc.expected)
		}
	}
}
