package main

import (
	"testing"
)
// 2nOs3HqZd5
func TestRun(t *testing.T) {
	tests := []struct {
		name      string
		flag      string
		subString string
		sentence  string
		expected  string
	}{
		{
			name:      "Audit 1: Color everything red",
			flag:      "--color=red",
			subString: "",
			sentence:  "hello world",
			expected:  "", // Replace with actual output or Golden File
		},
		{
			name:      "Audit 2: Color everything green",
			flag:      "--color=green",
			subString: "",
			sentence:  "1 + 1 = 2",
			expected:  "",
		},
		{
			name:      "Audit 3: Color everything yellow",
			flag:      "--color=yellow",
			subString: "",
			sentence:  "(%&) ??",
			expected:  "",
		},
		{
			name:      "Audit 4: Color second until the last letter",
			flag:      "--color=cyan", // The audit didn't specify a color, so we use cyan
			subString: "ello",
			sentence:  "hello",
			expected:  "",
		},
		{
			name:      "Audit 5: Color the second letter",
			flag:      "--color=magenta", // Using magenta as a placeholder
			subString: "e",
			sentence:  "hello",
			expected:  "",
		},
		{
			name:      "Audit 6: Color just two letters",
			flag:      "--color=white", // Using white as a placeholder
			subString: "ll",
			sentence:  "hello",
			expected:  "",
		},
		{
			name:      "Audit 7: Color specific word with different casing",
			flag:      "--color=orange",
			subString: "GuYs",
			sentence:  "HeY GuYs",
			expected:  "",
		},
		{
			name:      "Audit 8: Color single uppercase letter in symbols",
			flag:      "--color=blue",
			subString: "B",
			sentence:  "RGB()",
			expected:  "",
		},
	}

	content, _ := readFile()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Runner(content, tt.flag, tt.subString, tt.sentence)
			if got != tt.result {
				t.Errorf("Expected:\n%q\n\nGot:\n%q", tt.result, got)
			}
		})
	}
}
