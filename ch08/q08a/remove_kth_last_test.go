package ch08q08a

import "testing"

func TestRemoveKthLast(t *testing.T) {
	linkedList := &Node{1, nil}
	appendNode(linkedList, Node{2, nil})
	appendNode(linkedList, Node{3, nil})
	appendNode(linkedList, Node{4, nil})
	appendNode(linkedList, Node{4, nil})
	appendNode(linkedList, Node{5, nil})
	appendNode(linkedList, Node{6, nil})

	RemoveKthLast(linkedList, 2)

	ptr := linkedList
	i := 1
	for ptr.Next != nil {
		if ptr.Val != i {
			t.Errorf("Expected node to have val %v but got %v", i, ptr.Val)
		}
		ptr = ptr.Next
		i++
	}
}

func appendNode(root *Node, node Node) {
	ptr := root
	for ptr.Next != nil {
		ptr = ptr.Next
	}

	ptr.Next = &node
}
