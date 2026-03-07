package main

import (
	"strings"
)

// func colors() {
// 	color := map[string]string{
// 		"reset":   "\033[0m",
// 		"red":     "\033[31m",
// 		"green":   "\033[32m",
// 		"yellow":  "\033[33m",
// 		"blue":    "\033[34m",
// 		"magenta": "\033[35m",
// 		"cyan":    "\033[36m",
// 		"gray":    "\033[37m",
// 		"white":   "\033[97m",
// 	}

// }

func Runner(result, flag, subString string, sentence string) string {

	colorName := strings.TrimPrefix(flag, "--color=")

	for _, ch := range 


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
