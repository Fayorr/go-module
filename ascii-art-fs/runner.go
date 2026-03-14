package main

import (
	"fmt"
	"os"
	"strings"
)

func Runner(sentence, banner string) string {

	var bannerFile string
	switch banner {
	case "shadow":
		bannerFile = "./banner/shadow.txt"
	case "thinkertoy":
		bannerFile = "./banner/thinkertoy.txt"
	case "standard":
		bannerFile = "./banner/standard.txt"
	default:
		bannerFile = "./banner/standard.txt"
	}

	// Read the banner file containing ASCII art character definitions
	result, err := os.ReadFile(bannerFile)

	if err != nil {
		fmt.Printf("Error reading banner file: %s", err)
		os.Exit(1)
	}

	// Split content by newlines
	trimmedResult := strings.Split(string(result), "\n")
	wordSlice := strings.Split(sentence, "\\n")

	var finalString strings.Builder
	for j := 0; j < len(wordSlice); j++ {
		if wordSlice[j] == "" {
			finalString.WriteRune('\n')
			continue
		}

		for i := 1; i <= 8; i++ {
			for _, ch := range wordSlice[j] {
				pos := (int(ch-' ') * 9)
				finalString.WriteString(trimmedResult[pos+i])
			}
			finalString.WriteRune('\n')
		}
	}
	return strings.TrimRight(finalString.String(), "\n")
}
