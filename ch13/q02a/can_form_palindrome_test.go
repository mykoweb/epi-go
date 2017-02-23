package ch13q02a

import "testing"

func TestCanFormPalindrom(t *testing.T) {
	result := CanFormPalindrome("edified")

	if result == false {
		t.Errorf("CanFormPalindrome expected true for 'edified' but got false")
	}

	result = CanFormPalindrome("lleev")

	if result == false {
		t.Errorf("CanFormPalindrome expected true for 'lleev' but got false")
	}

	result = CanFormPalindrome("abcd")

	if result == true {
		t.Errorf("CanFormPalindrome expected false for 'abcd' but got true")
	}
}
