package ch08q04a

type Node struct {
	val  int
	next *Node
}

// In go we cannot perform node.next.next without getting a nil pointer
// dereference error on nil pointers. A workaround is to create a Next() method
// that can handle the case where we call Next() on a nil pointer.
func (n *Node) Next() *Node {
	if n == nil {
		return nil
	} else {
		return n.next
	}
}

// FindCycleStart returns a pointer to the node at the start of a cycle if it
// exists. If there is no cycle, it returns nil.
func FindCycleStart(head *Node) *Node {
	// If the linked list is empty or has only one node, then there is no cycle.
	if head == nil || head.Next() == nil {
		return nil
	}

	var cycleLength int
	var trailingPtr, advancedByCycleLengthPtr *Node

	slowPtr := head
	fastPtr := head.Next().Next()

	// The outer for loop advances fastPtr 2 steps for every 1 step slowPtr is
	// advanced. Keep advancing both pointers as such until they point to the
	// same node, or until fastPtr reaches the end of the linked list. If the
	// latter happens, then there was no cycle.
	for fastPtr != nil {
		if slowPtr == fastPtr {
			// Since slowPtr and fastPtr are pointing to the same node, we have found
			// a cycle. Now keep advancing slowPtr until it meets fastPtr again in
			// order to find the cycle length.
			slowPtr = slowPtr.Next()
			cycleLength++
			for slowPtr != fastPtr {
				slowPtr = slowPtr.Next()
				cycleLength++
			}

			// Starting from head, advance the lead pointer by the cycle length.
			trailingPtr = head
			advancedByCycleLengthPtr = head
			for ; cycleLength > 0; cycleLength-- {
				advancedByCycleLengthPtr = advancedByCycleLengthPtr.Next()
			}

			// Now advance both pointers in step until they point to the same node.
			// Once this happens, we are at the start of the cycle.
			for trailingPtr != advancedByCycleLengthPtr {
				trailingPtr = trailingPtr.Next()
				advancedByCycleLengthPtr = advancedByCycleLengthPtr.Next()
			}

			return trailingPtr
		}
		slowPtr = slowPtr.Next()
		fastPtr = fastPtr.Next().Next()
	}

	return nil
}
