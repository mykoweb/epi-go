package ch07q11a

import "testing"

func TestMakeSnake(t *testing.T) {
	s := "Hello world!"
	result := MakeSnake(s)
	if result != "e lHlowrdlo!" {
		t.Errorf("Expected e lHlowrdlo! but got %v", result)
	}
}
