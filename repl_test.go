package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single word",
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			name:     "multiple words",
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "mixed case",
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			name:     "leading/trailing spaces",
			input:    "  leading AND trailing   ",
			expected: []string{"leading", "and", "trailing"},
		},
		{
			name:     "multiple spaces between words",
			input:    "too    many   spaces",
			expected: []string{"too", "many", "spaces"},
		},
		{
			name:     "tabs and newlines",
			input:    "some\twords\nhere",
			expected: []string{"some", "words", "here"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := cleanInput(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("cleanInput(%q): expected %v, got %v", tc.input, tc.expected, actual)
			}
		})
	}
}
