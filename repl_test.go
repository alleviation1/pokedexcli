package main

import "testing"


func TestCleanInput(t *testing.T) {
	return
	cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	// add more cases here
 }

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Unequal output lengths")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if expectedWord != word {
				t.Errorf("Words don't match")
			}
		}
	}
}

// func TestStartRepl(t *testing.T) {
// 	return
// 	cases := []struct {
// 		input 	 []string
// 		expected []string
// 	}{
// 		{
// 			input: 	  []string{"help", "exit"}
// 			expected: []string{}
// 		}
// 	}
// }