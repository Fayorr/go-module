package main

import (
	"os"
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
			expected:  "./test-files/audit1.txt", // Replace with actual output or Golden File
		},
		{
			name:      "Audit 2: Color everything green",
			flag:      "--color=green",
			subString: "",
			sentence:  "1 + 1 = 2",
			expected:  "./test-files/audit2.txt",
		},
		{
			name:      "Audit 3: Color everything yellow",
			flag:      "--color=yellow",
			subString: "",
			sentence:  "(%&) ??",
			expected:  "./test-files/audit3.txt",
		},
		{
			name:      "Audit 4: Color second until the last letter",
			flag:      "--color=cyan", // The audit didn't specify a color, so we use cyan
			subString: "RNIGERDACT",
			sentence:  "MRNIGERDACT",
			expected:  "./test-files/audit4.txt",
		},
		{
			name:      "Audit 5: Color the second letter",
			flag:      "--color=magenta", // Using magenta as a placeholder
			subString: "e",
			sentence:  "hello",
			expected:  "./test-files/audit5.txt",
		},
		{
			name:      "Audit 6: Color just two letters",
			flag:      "--color=white", // Using white as a placeholder
			subString: "ii",
			sentence:  "ascii",
			expected:  "./test-files/audit6.txt",
		},
		{
			name:      "Audit 7: Color specific word with different casing",
			flag:      "--color=orange",
			subString: "GuYs",
			sentence:  "HeY GuYs",
			expected:  "./test-files/audit7.txt",
		},
	}

	content, _ := readFile()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Runner(content, tt.flag, tt.subString, tt.sentence)
			expectedBytes, err := os.ReadFile(tt.expected)

			if err != nil {
				t.Errorf("Failed to Read File: %v", err)
			}
			expectedString := string(expectedBytes)

			if got != expectedString {
				t.Errorf("Expected:\n%q\n\nGot:\n%q", expectedString, got)
			}
		})
	}
}
