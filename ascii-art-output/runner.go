package main

import (
	"fmt"
	"os"
	"strings"
)

func Runner(flag, sentence, banner string) string {

	var bannerFile string 

	if banner == "shadow" {
		bannerFile = "shadow.txt"
	} else if banner == "thinkertoy" {
		bannerFile = "thinkertoy.txt"
	} else if banner == "standard" {
		bannerFile = "standard.txt"
	} else {
		return ""
	}
	}
	
	result, err := os.ReadFile(bannerFile)
	fileName := strings.TrimPrefix(flag, "--output=")


	if err != nil {
		fmt.Errorf("Error reading file: %s", err)
	}

	trimmedResult := strings.Split(result, "\n")

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
	return finalString.String()
}
