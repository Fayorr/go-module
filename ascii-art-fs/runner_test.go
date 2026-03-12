package main

import (
	"fmt"
	"os"
	"testing"
)

func TestAll(t *testing.T) {
	 testAll := []struct{
			name string
			sentence string
			banner string
			expected string
		}{
			{
			name: "Audit1",
			sentence: "First\nTest",
			banner: "shadow",
			expected: "./test-files/test00.txt",
		},
		{
			name: "Audit2",
			sentence: "hello",
			banner: "standard",
			expected: "./test-files/test01.txt",
		},
		{
			name: "Audit3",
			sentence: "123 -> #$%",
			banner: "standard",
			expected: "./test-files/test02.txt",
		},
		{
			name: "Audit4",
			sentence: "432 -> #$%&@",
			banner: "shadow",
			expected: "./test-files/test03.txt",
		},
	}
	
	for _, tt := range testAll {
		expect, err := os.ReadFile(tt.expected)
		if err != nil {
			fmt.Println("couldnt read file for tests")
		}
		t.Run(tt.name, func(t *testing.T) {
			got := Runner(tt.sentence, tt.banner)
				if got != string(expect) {
					t.Errorf("Expected:%q, Got:%q, ", tt.expected, got)
			}
		})
	}
}