package formatter

func Runner(sentence string) string {
	step1 := CorrectCase(sentence)
	step2 := CorrectVowel(step1)
	step3 := CorrectPunctuations(step2)
	step4 := CorrectQuote(step3)

	return step4
}
