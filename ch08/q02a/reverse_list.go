package ch08q02a

type Node struct {
	val  int
	next *Node
}

// ReverseList reverses a singly-linked list and returns a pointer to that
// reversed list.
func ReverseList(root *Node) *Node {
	var prev, curr, next *Node
	curr = root

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev
}
