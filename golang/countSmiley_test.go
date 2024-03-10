package main

import (
	"testing"
)

func TestCountSmileys(t *testing.T) {
	tests := []struct {
		caseName string
		input    []string
		expected int
	}{
		{
			caseName: "case 1: two smileys",
			input:    []string{":)", ";(", ";}", ":-D"},
			expected: 2,
		},
		{
			caseName: "case 2: three smileys",
			input:    []string{";D", ":-(", ":-)", ";~)"},
			expected: 3,
		},
		{
			caseName: "case 3: one smileys",
			input:    []string{";]", ":[", ";*", ":$", ";-D"},
			expected: 1,
		},
		{
			caseName: "case 4: no smileys",
			input:    []string{":[", ";(", ";}", ":-|"},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := CountSmileys(tt.input)
			if result != tt.expected {
				t.Errorf(`
	case:     %s is Failed
	input:    %+v
	expected: %d
	but got:  %d
	`, tt.caseName, tt.input, tt.expected, result)
			}
		})
	}
}
