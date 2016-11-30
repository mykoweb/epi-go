package ch09q12a

import "testing"

func TestMaxQueue(t *testing.T) {
	q := Queue{}

	_, ok := q.Dequeue()
	if ok != false {
		t.Errorf("Dequeueing empty queue should have returned ok false")
	}

	_, ok = q.Max()
	if ok != false {
		t.Errorf("Max() on empty queue should have returned ok false")
	}

	q.Enqueue(5)
	q.Enqueue(4)
	q.Enqueue(3)
	val, ok := q.Max()
	if val != 5 || ok != true {
		t.Errorf("Expected Max() to return 5 but returned %v", val)
	}

	val, ok = q.Dequeue()
	if val != 5 || ok != true {
		t.Errorf("Expected Dequeue() to return 5 but returned %v", val)
	}
	val, ok = q.Max()
	if val != 4 || ok != true {
		t.Errorf("Expected Max() to return 4 but returned %v", val)
	}

	q.Enqueue(6)
	q.Enqueue(1)
	val, ok = q.Max()
	if val != 6 || ok != true {
		t.Errorf("Expected Max() to return 6 but returned %v", val)
	}
	val, ok = q.Dequeue()
	if val != 4 || ok != true {
		t.Errorf("Expected Dequeue() to return 4 but returned %v", val)
	}
	val, ok = q.Max()
	if val != 6 || ok != true {
		t.Errorf("Expected Max() to return 6 but returned %v", val)
	}
	val, ok = q.Dequeue()
	if val != 3 || ok != true {
		t.Errorf("Expected Dequeue() to return 3 but returned %v", val)
	}
	val, ok = q.Max()
	if val != 6 || ok != true {
		t.Errorf("Expected Max() to return 6 but returned %v", val)
	}
	val, ok = q.Dequeue()
	if val != 6 || ok != true {
		t.Errorf("Expected Dequeue() to return 6 but returned %v", val)
	}
	val, ok = q.Max()
	if val != 1 || ok != true {
		t.Errorf("Expected Max() to return 1 but returned %v", val)
	}
	val, ok = q.Dequeue()
	if val != 1 || ok != true {
		t.Errorf("Expected Dequeue() to return 1 but returned %v", val)
	}

	_, ok = q.Max()
	if ok != false {
		t.Errorf("Max() on empty queue should have returned ok false")
	}
	_, ok = q.Dequeue()
	if ok != false {
		t.Errorf("Dequeueing empty queue should have returned ok false")
	}
}
