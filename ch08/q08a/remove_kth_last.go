package ch08q08a

// RemoveKthLast removes the kth last element from a singly linked list.
//
// The time complexity is O(n) where n is the length of the linked list.
func RemoveKthLast(root *Node, k int) {
	if root == nil || k < 0 {
		return
	}

	leadPtr := root
	lagPtr := root

	for i := 1; i <= k; i++ {
		leadPtr = leadPtr.Next
		if leadPtr == nil {
			return
		}
	}

	for leadPtr.Next != nil {
		leadPtr = leadPtr.Next
		lagPtr = lagPtr.Next
	}

	lagPtr.Val = lagPtr.Next.Val
	lagPtr.Next = lagPtr.Next.Next
}

type Node struct {
	Val  int
	Next *Node
}
