package ch09q01

import "testing"

func TestMaxStack(t *testing.T) {
	stack := MaxStack{}

	stack.Push(-100)
	val, _ := stack.Max()
	if val != -100 {
		t.Errorf("Expected -100 but got %v", val)
	}

	stack.Push(100)
	val, _ = stack.Max()
	if val != 100 {
		t.Errorf("Expected 100 but got %v", val)
	}

	stack.Push(64)
	val, _ = stack.Max()
	if val != 100 {
		t.Errorf("Expected 100 but got %v", val)
	}

	stack.Push(212)
	val, _ = stack.Max()
	if val != 212 {
		t.Errorf("Expected 212 but got %v", val)
	}

	stack.Push(212)
	val, _ = stack.Max()
	if val != 212 {
		t.Errorf("Expected 212 but got %v", val)
	}

	stack.Push(32)
	val, _ = stack.Max()
	if val != 212 {
		t.Errorf("Expected 212 but got %v", val)
	}

	val, _ = stack.Pop()
	if val != 32 {
		t.Errorf("Expected 32 but got %v", val)
	}
	val, _ = stack.Max()
	if val != 212 {
		t.Errorf("Expected 212 but got %v", val)
	}

	val, _ = stack.Pop()
	if val != 212 {
		t.Errorf("Expected 212 but got %v", val)
	}
	val, _ = stack.Max()
	if val != 212 {
		t.Errorf("Expected 212 but got %v", val)
	}

	val, _ = stack.Pop()
	if val != 212 {
		t.Errorf("Expected 212 but got %v", val)
	}
	val, _ = stack.Max()
	if val != 100 {
		t.Errorf("Expected 100 but got %v", val)
	}

	val, _ = stack.Pop()
	if val != 64 {
		t.Errorf("Expected 64 but got %v", val)
	}
	val, _ = stack.Max()
	if val != 100 {
		t.Errorf("Expected 100 but got %v", val)
	}

	val, _ = stack.Pop()
	if val != 100 {
		t.Errorf("Expected 100 but got %v", val)
	}
	val, _ = stack.Max()
	if val != -100 {
		t.Errorf("Expected -100 but got %v", val)
	}

	val, _ = stack.Pop()
	if val != -100 {
		t.Errorf("Expected -100 but got %v", val)
	}
	val, ok := stack.Max()
	if ok != false {
		t.Errorf("Expected Max() to return ok false but it did not")
	}
}
