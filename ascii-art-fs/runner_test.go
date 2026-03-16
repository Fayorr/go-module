package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	testAll := []struct {
		name     string
		sentence string
		banner   string
		expected string
	}{
		{
			name:     "Audit1",
			sentence: "hello",
			banner:   "standard",
			expected: "./test-files/audit1.txt",
		},
		{
			name:     "Audit2",
			sentence: "hello world",
			banner:   "shadow",
			expected: "./test-files/audit2.txt",
		},
		{
			name:     "Audit3",
			sentence: "nice 2 meet you",
			banner:   "thinkertoy",
			expected: "./test-files/audit3.txt",
		},
		{
			name:     "Audit4",
			sentence: "you & me",
			banner:   "standard",
			expected: "./test-files/audit4.txt",
		},
		{
			name:     "Audit5",
			sentence: "123",
			banner:   "shadow",
			expected: "./test-files/audit5.txt",
		},
		{
			name:     "Audit6",
			sentence: "/(\")",
			banner:   "thinkertoy",
			expected: "./test-files/audit6.txt",
		},
		{
			name:     "Audit7",
			sentence: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			banner:   "shadow",
			expected: "./test-files/audit7.txt",
		},
		{
			name:     "Audit8",
			sentence: "\"#$%&/()*+,-./",
			banner:   "thinkertoy",
			expected: "./test-files/audit8.txt",
		},
		{
			name:     "Audit9",
			sentence: "It's Working",
			banner:   "thinkertoy",
			expected: "./test-files/audit9.txt",
		},
	}

	for _, tt := range testAll {
		expect, err := os.ReadFile(tt.expected)
		expected := strings.TrimRight(string(expect), "\n")
		if err != nil {
			fmt.Println("couldnt read file for tests")
		}
		t.Run(tt.name, func(t *testing.T) {
			got := Runner(tt.sentence, tt.banner)
			if got != expected {
				t.Errorf("Expected:%q, Got:%q, ", string(expect), got)
			}
		})
	}
}
