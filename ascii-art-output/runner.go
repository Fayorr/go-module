package main

import (
	"fmt"
	"os"
	"strings"
)

func Runner(sentence, banner string) string {

	var bannerFile string 

	if banner == "shadow" {
		bannerFile = "./banner/shadow.txt"
	} else if banner == "thinkertoy" {
		bannerFile = "./banner/thinkertoy.txt"
	} else if banner == "standard" {
		bannerFile = "./banner/standard.txt"
	} else {
		return ""
	}
	


	result, err := os.ReadFile(bannerFile)

	if err != nil {
		fmt.Printf("Error reading file: %s", err)
		os.Exit(1)
	}

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
	return finalString.String()
}
