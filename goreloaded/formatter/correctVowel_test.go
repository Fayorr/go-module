package formatter

import "testing"

func Test_CorrectVowel(t *testing.T) {

	result := CorrectVowel("this is a antelope")
	if result != "this is an antelope" {
		t.Errorf("Expected %s, got %s", "this is an antelope", result)
	}
}
