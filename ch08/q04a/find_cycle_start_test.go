package ch08q04a

import "testing"

func TestFindCycleStart(t *testing.T) {
	// Test with a linked list with a cycle.
	head := &Node{}
	head.next = &Node{}
	head.next.next = &Node{}
	head.next.next.next = &Node{}
	cycleStart := head.next.next.next // Make this the start of the cycle

	cycleStart.next = &Node{}
	cycleStart.next.next = &Node{}
	cycleStart.next.next.next = &Node{}
	cycleStart.next.next.next.next = cycleStart // Create the cycle

	result := FindCycleStart(head)

	if result != cycleStart {
		t.Errorf("Expected to find the start of the cycle but we did not")
	}

	// Test with a linked list with no cycle.
	head = &Node{}
	head.next = &Node{}
	head.next.next = &Node{}
	head.next.next.next = &Node{}

	result = FindCycleStart(head)

	if result != nil {
		t.Errorf("Expected to find no cycle but we found a cycle")
	}
}
