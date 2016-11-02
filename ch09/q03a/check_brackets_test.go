package ch09q03a

import "testing"

func TestCheckBrackets(t *testing.T) {
	s := "[()[]{()()}]"
	if CheckBrackets(s) == false {
		t.Errorf("Expected true but got false for", s)
	}

	s = "{]"
	if CheckBrackets(s) == true {
		t.Errorf("Expected false but got true for", s)
	}

	s = "({}[)"
	if CheckBrackets(s) == true {
		t.Errorf("Expected false but got true for", s)
	}

	s = ""
	if CheckBrackets(s) == false {
		t.Errorf("Expected true but got false for", s)
	}
}
