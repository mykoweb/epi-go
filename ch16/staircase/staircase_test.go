package staircase

import "testing"

func TestStaircase(t *testing.T) {
	result := Staircase(4)
	if result != 7 {
		t.Errorf("Expected result to be 7 but got %v", result)
	}
}
