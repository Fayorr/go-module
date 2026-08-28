package main

import (
	"os"
	"strings"
)

func Runner(sentence, banner string) (string, error) {

	var bannerFile string
	switch banner {
	case "shadow":
		bannerFile = "banner/shadow.txt"
	case "thinkertoy":
		bannerFile = "banner/thinkertoy.txt"
	case "standard":
		bannerFile = "banner/standard.txt"
	default:
		bannerFile = "banner/standard.txt"
	}

	// Read the banner file containing ASCII art character definitions
	result, err := os.ReadFile(bannerFile)

	if err != nil {
		return "", err
	}

	trimmedResult := strings.Split(string(result), "\n")

	wordSlice := strings.Split(sentence, "\r\n")

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
	return finalString.String(), nil
}
