package main

import (
	"testing"

	"github.com/ahmedgaber19/pokedexcli/repl"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	}}
	for _, c := range cases {
		actual := repl.CleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				// if they don't match, use t.Errorf to print an error message
				// and fail the test
				t.Errorf("Expected %s but got %s", expectedWord, word)
				t.Fail()
			}
		}
	}
}
