package main

import (
	"strings"
)

func Runner(result string, arguments []string) string {
	trimmedResult := strings.Split(result, "\n")

	word := strings.Join(arguments, " ")

	wordSlice := strings.Split(word, "\\n")

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
