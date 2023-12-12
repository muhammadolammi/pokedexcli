package main

import "testing"

func TestCleanText(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "How far faith",
			expected: []string{"how", "far", "faith"},
		},
	}

	for _, cs := range cases {
		// test for the lenght of output
		actual := cleanText(cs.input)

		if len(actual) != len(cs.expected) {
			t.Errorf("then len of %v is not equal %v ", cs.input, cs.expected)
		}
		// test for the actual word
		for i := range actual {
			acWord := actual[i]
			exWord := cs.expected[i]
			if acWord != exWord {
				t.Errorf("Error %v is not equal to %v", acWord, exWord)
			}
		}
	}

}
