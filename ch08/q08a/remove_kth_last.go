package ch08q08a

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
