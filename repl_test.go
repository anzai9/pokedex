package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "    help    ",
			expected: []string{"help"},
		},
		{
			input:    "    help    help",
			expected: []string{"help", "help"},
		},
		{
			input:    "    help    Help HELP",
			expected: []string{"help", "help", "help"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length don't match: actual  %v, expected %v", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			if word != c.expected[i] {
				t.Errorf("cleanInput(%s) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
