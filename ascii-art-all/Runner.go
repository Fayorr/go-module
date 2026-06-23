package main

import (
	"fmt"
	"os"
	"strings"
)

func Runner(input, banner string) (string, error) {

	bannerPath := ""
	// banner
	switch banner {
	case "standard":
		{
			bannerPath = "./banner/standard.txt"
		}
	case "shadow":
		{
			bannerPath = "./banner/shadow.txt"
		}
	case "thinkertoy":
		{
			bannerPath = "./banner/thinkertoy.txt"
		}
	default:
		bannerPath = "./banner/standard.txt"
	}

	fileContent, err := os.ReadFile(bannerPath)

	if err != nil {
		fmt.Println("Error reading file content")
		return "", err
	}

	bannerLines := strings.Split(strings.ReplaceAll(string(fileContent), "\r\n", "\n"), "\n")

	// input
	wordSlice := strings.Split(input, "\\n")

	var result strings.Builder
	for i := 0; i < len(wordSlice); i++ {
		if wordSlice[i] == "" {
			result.WriteRune('\n')
			continue
		}
		for k := 0; k < 8; k++ {
			for j := 0; j < len(wordSlice[i]); j++ {
				pos := int(rune(wordSlice[i][j])- ' ') * 9 + k +1
				result.WriteString(bannerLines[pos])
			}
			result.WriteByte('\n')
		}
	}
	return result.String(), nil
}
