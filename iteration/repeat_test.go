package iteration

import "testing"

func TestRepeat(t *testing.T) {
	result := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if result != expected {
		t.Errorf("I was expected '%s' but got  '%s'", expected, result)
	}
}
