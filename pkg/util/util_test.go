package util

import "testing"

func TestCollapseWhitespace(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Basic cases
		{"Hello   World", "Hello World"},
		{"  Leading and trailing   ", "Leading and trailing"},
		{"Multiple\tspaces\nand\nnewlines", "Multiple\tspaces\nand\nnewlines"},
		{"Tab\tand space", "Tab\tand space"},

		// Punctuation spacing
		{"Hello , world !", "Hello, world!"},
		{"Wait  ...  what ?", "Wait... what?"},
		{"One , two ; three : done .", "One, two; three: done."},

		// Only whitespace
		{"    ", ""},
		{"\t\n\r", ""},

		// No whitespace
		{"NoChange", "NoChange"},

		// Mixed content
		{"  Hello \t world  ! How are you  ? ", "Hello \t world! How are you?"},

		// Multiple punctuation
		{"Hello  ,  world  !  ", "Hello, world!"},

		// Unicode spaces
		{"Hello\u00A0World", "Hello World"}, // \u00A0 = non-breaking space

		// Already formatted
		{"Perfectly formatted sentence.", "Perfectly formatted sentence."},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := CollpaseWhitespace(test.input)
			if result != test.expected {
				t.Errorf("Expected '%s' but got '%s'", test.expected, result)
			}
		})
	}
}
