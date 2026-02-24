package formatter

import "strings"

func CorrectVowel(sentence string) string {
	// I have a banana and an apple"
	var result strings.Builder
	for i := 0; i < len(sentence); i++ {
		result.WriteByte(sentence[i])
		if (i == 0 || sentence[i-1] == ' ') && sentence[i] == 'a' || sentence[i] == 'A' {
			if i+2 < len(sentence) && (sentence[i+1] == ' ' && strings.ContainsRune("aeiouhAEIOUH", rune(sentence[i+2]))) {
				result.WriteByte('n')
			}
		}
	}

	return result.String()
}
