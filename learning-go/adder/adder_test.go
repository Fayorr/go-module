package adder

import "testing"

func Test_adder(t *testing.T) {
	result := addNumbers(4, 4)
	if result != 9 {
		t.Error("Error occured, expected 9, got:", result)
	}
}
