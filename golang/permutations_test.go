package main

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	testCases := []struct {
		caseName string
		input    string
		expected []string
	}{
		{
			caseName: "case 1: one character input",
			input:    "a",
			expected: []string{"a"},
		},
		{
			caseName: "case 2: two characters input",
			input:    "ab",
			expected: []string{"ab", "ba"},
		},
		{
			caseName: "case 3: three characters input",
			input:    "abc",
			expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			caseName: "case 4: four characters input",
			input:    "aabb",
			expected: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
		{
			caseName: "case 5: same two characters",
			input:    "aa",
			expected: []string{"aa"},
		},
		{
			caseName: "case 6: same three characters",
			input:    "aaa",
			expected: []string{"aaa"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actualResult := Permutations(tc.input)
			if !reflect.DeepEqual(actualResult, tc.expected) {
				t.Errorf(`
	case:     %s is Failed
	input:    %s
	expected: %+v, 
	but got:  %+v
	`, tc.caseName, tc.input, tc.expected, actualResult)
			}
		})
	}
}
