package ch15q04a

type Node struct {
	Val         int
	Left, Right *Node
}

// FindKLargest returns an int slice containing the k largest elements in a BST.
func FindKLargest(root *Node, k int, kLargest *[]int) {
	if root == nil {
		return
	}

	if len(*kLargest) < k {
		FindKLargest(root.Right, k, kLargest)
		if len(*kLargest) < k {
			*kLargest = append(*kLargest, root.Val)
			FindKLargest(root.Left, k, kLargest)
		}
	}
}
