package formatter

import "testing"

func Test_correctQuote(t *testing.T) {
	result := CorrectQuote("Harold Wilson: ' I am an optimist, but an optimist who carries a raincoat. '")
	if result != "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'" {
		t.Errorf("Expected: %s, Got: %s", "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'", result)
	}
}
