package formatter

import "testing"

func Test_correctPunctuations(t *testing.T) {
	result := CorrectQuote("I am exactly how they describe me: ' awesome '")
	correctResult := "I am exactly how they describe me: 'awesome'"
	if result != correctResult {
		t.Errorf("Expected: %s, Got: %s", correctResult, result)
	}
}
