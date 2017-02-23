package ch08q01a

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

func (n *Node) Val() int {
	return n.val
}

// MergeSortedLists merges 2 lists that are both sorted into a single list.
func MergeSortedLists(aPtr, bPtr *Node) *Node {
	var mergedList *Node
	mergedPtr := mergedList

	for aPtr != nil && bPtr != nil {
		if bPtr == nil || aPtr.Val() < bPtr.Val() {
			mergedPtr = &Node{val: aPtr.Val()}
			aPtr = aPtr.Next()
		} else if aPtr == nil || aPtr.Val() >= bPtr.Val() {
			mergedPtr = &Node{val: bPtr.Val()}
			bPtr = bPtr.Next()
		}
		mergedPtr = mergedPtr.Next()
	}

	return mergedList
}
