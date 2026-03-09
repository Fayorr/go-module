package main

import (
	"strings"
)

func Runner(result, flag, subString string, sentence string) string {

	colorDict := map[string]string{
		"reset":   "\033[0m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"purple":  "\033[35m",
		"cyan":    "\033[36m",
		"gray":    "\033[37m",
		"white":   "\033[97m",
		"orange":  "\033[38;5;208m",

		// --- HEX CODES ---
		"#000000": "\033[30m",       // Black
		"#ff0000": "\033[31m",       // Red
		"#00ff00": "\033[32m",       // Green
		"#ffff00": "\033[33m",       // Yellow
		"#0000ff": "\033[34m",       // Blue
		"#ff00ff": "\033[35m",       // Magenta
		"#00ffff": "\033[36m",       // Cyan
		"#ffffff": "\033[37m",       // White
		"#ffa500": "\033[38;5;208m", // Orange

	}

	colorName := strings.TrimPrefix(flag, "--color=")

	ansiColor := colorDict[colorName]
	ansiReset := colorDict["reset"]

	trimmedResult := strings.Split(result, "\n")

	wordSlice := strings.Split(sentence, "\\n")

	var finalString strings.Builder
	for j := 0; j < len(wordSlice); j++ {
		if wordSlice[j] == "" {
			finalString.WriteRune('\n')
			continue
		}
		// define the map with length of current word in sentence
		colorMap := make([]bool, len(wordSlice[j]))

		// check if theres a subString and colours everything if there isn't
		if subString == "" {
			for k := 0; k < len(colorMap); k++ {
				colorMap[k] = true
			}
		} else {
			for k := 0; k < len(wordSlice[j]); k++ {
				if strings.HasPrefix(wordSlice[j][k:], subString) {
					for l := 0; l < len(subString); l++ {
						colorMap[k+l] = true
					}
				}
			}
		}

		for i := 1; i <= 8; i++ {
			for index, ch := range wordSlice[j] {
				pos := (int(ch-' ') * 9)
				if colorMap[index] == true {
					finalString.WriteString(ansiColor)
				}
				finalString.WriteString(trimmedResult[pos+i])

				if colorMap[index] == true {
					finalString.WriteString(ansiReset)
				}
			}
			finalString.WriteRune('\n')
		}
	}
	return finalString.String()
}
