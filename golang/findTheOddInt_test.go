package main

import (
	"testing"
)

func TestFindTheOddInt(t *testing.T) {
	tests := []struct {
		caseName string
		input    []string
		expected int
	}{
		{
			caseName: "one odd number",
			input:    []string{"1", "2", "3"},
			expected: 2,
		},
		{
			caseName: "no odd numbers",
			input:    []string{"2", "4", "6"},
			expected: 0,
		},
		{
			caseName: "all odd numbers",
			input:    []string{"1", "3", "5"},
			expected: 3,
		},
	}
	for _, tc := range tests {
		t.Run(tc.caseName, func(t *testing.T) {
			if got := FindTheOddInt(tc.input); got != tc.expected {
				t.Errorf(`
	case:     %s is Failed
	input:    %+v 
	expected: %d
	but got:  %d
	`, tc.caseName, tc.input, tc.expected, got)
			}
		})
	}
}
