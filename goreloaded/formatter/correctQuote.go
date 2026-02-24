package formatter

import "strings"

func CorrectQuote(sentence string) string {
	var result strings.Builder

	insideQuote := false
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == '\'' {
			insideQuote = !insideQuote
		}
		if sentence[i] == ' ' && insideQuote {
			if (i-1 >= 0 && sentence[i-1] == '\'') || (i+1 < len(sentence) && sentence[i+1] == '\'') {
				continue
			}
		}
		result.WriteByte(sentence[i])
	}
	return result.String()
}
