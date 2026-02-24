package formatter

import "testing"

func TestRunAll(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string 
	}{
		{
			name:     "Harold Wilson Quote (The Final Boss)",
			input:    "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			expected: "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
		{
			name:     "Basic Hex and Bin",
			input:    "1E (hex) files and 10 (bin) years",
			expected: "30 files and 2 years",
		},
		{
			name:     "Punctuation Spacing",
			input:    "I was sitting over there ,and then BAMM !!",
			expected: "I was sitting over there, and then BAMM!!",
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
