package formatter

import "testing"

// This is so exciting (up, 2)

func Test_CorrectCase(t *testing.T) {
	result := CorrectCase("This is so exciting (up, 2)")
	correctResult := "This is SO EXCITING"
	if result != correctResult {
		t.Errorf("Expected: %s, Got: %s", correctResult, result)
	}
}
