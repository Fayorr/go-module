package formatter

import "testing"

func TestRunAll(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Audit 1",
			input:    "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			expected: "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			name:     "Audit 2",
			input:    "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			expected: "I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			name:     "Audit 3",
			input:    "Don not be sad ,because sad backwards is das . And das not good",
			expected: "Don not be sad, because sad backwards is das. And das not good",
		},
		{
			name:     "Audit 4",
			input:    "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			expected: "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
	}
	// Loop through every test case in the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Call our runner function
			result := Runner(tt.input)

			// Check if the result matches what the auditors expect
			if result != tt.expected {
				t.Errorf("\nTest Failed: %s\nExpected: %s\nGot:      %s", tt.name, tt.expected, result)
			}
		})
	}
}
