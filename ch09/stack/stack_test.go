package stack

import "testing"

func TestStack(t *testing.T) {
	s := new(Stack)

	for i := 8; i > 0; i-- {
		s.Push(i)
	}

	var val interface{}
	val, _ = s.Peek()
	if val != 1 {
		t.Errorf("Expected Peek() to return 1 but got %v", val)
	}

	if s.Len() != 8 {
		t.Errorf("Expected Len() to return 8 but got %v", s.Len())
	}

	for i := 1; i <= 8; i++ {
		val = s.Pop()
		if val != i {
			t.Errorf("Expected Pop() to return %v but got %v", i, s.Pop())
		}
	}

	if s.Pop() != nil {
		t.Errorf("Expected Pop() on an empty stack return nil but got %v", s.Pop())
	}

	val, _ = s.Peek()
	if val != nil {
		t.Errorf("Expected Peek() to return nil but got %v", val)
	}

	if s.Len() != 0 {
		t.Errorf("Expected Len() to return 0 but got %v", s.Len())
	}
}
