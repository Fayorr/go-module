package formatter

import "strings"

func CorrectPunctuations(sentence string) string {
	var result strings.Builder

	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' && i+1 < len(sentence) && strings.ContainsRune(".,!?:;", rune(sentence[i+1])) {
			continue
		}
		result.WriteByte(sentence[i])
		if strings.ContainsRune(".,!?:;", rune(sentence[i])) && i+1 < len(sentence) && !strings.ContainsRune(".,!?:; ", rune(sentence[i+1])) {
			result.WriteByte(' ')
		}
	}
	return result.String()
}
