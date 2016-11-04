package ch07q02a

import "testing"

func TestConvertBase(t *testing.T) {
	result := ConvertBase("615", 7, 13)

	if result != "1A7" {
		t.Errorf("Expected 1A7 but got %v", result)
	}
}
