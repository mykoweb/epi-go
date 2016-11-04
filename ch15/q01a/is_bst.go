package ch15q01a

func isBST(root *Node, low, high int64) bool {
	if root == nil {
		return true
	} else if root.Val < low || root.Val > high {
		return false
	}

	result := isBST(root.Left, low, root.Val)
	return result && isBST(root.Right, root.Val, high)
}

type Node struct {
	Val         int64
	Left, Right *Node
}
