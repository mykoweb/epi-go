package ch09q08a

import "testing"

func TestStackSort(t *testing.T) {
	s := new(Stack)
	s.Push(5)
	s.Push(4)
	s.Push(3)
	s.Push(2)
	s.Push(1)
	StackSort(s)

	for i := 5; i > 0; i-- {
		if val, _ := s.Pop(); val != i {
			t.Errorf("Expected %v but got %v", i, val)
		}
	}
}
