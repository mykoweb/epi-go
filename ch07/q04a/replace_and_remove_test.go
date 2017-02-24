package ch07q04a

import "testing"

func TestReplaceAndRemove(t *testing.T) {
	str := "acdbbca"

	result := ReplaceAndRemove(str)

	if result != "ddcdcdd" {
		t.Errorf("Expected to get 'ddcdcdd' but got %v", result)
	}
}
