package ch08q02a

import "testing"

func TestReverseList(t *testing.T) {
	root := &Node{val: 5}
	root.next = &Node{val: 4}
	root.next.next = &Node{val: 3}
	root.next.next.next = &Node{val: 2}
	root.next.next.next.next = &Node{val: 1}

	reversedRoot := ReverseList(root)

	i := 1
	for reversedRoot != nil {
		if reversedRoot.val != i {
			t.Errorf("Expected node to have value %v but got %v", i, reversedRoot.val)
		}
		reversedRoot = reversedRoot.next
		i += 1
	}
}
