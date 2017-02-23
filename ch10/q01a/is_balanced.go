package ch10q01a

type Node struct {
	val         int
	left, right *Node
}

// IsBalanced returns true if a binary tree is balance, false otherwise.
func IsBalanced(root *Node) (balanced bool, height int) {
	if root == nil {
		return true, 0
	}

	leftBalanced, leftHeight := IsBalanced(root.left)
	rightBalanced, rightHeight := IsBalanced(root.right)

	balanced = leftBalanced && rightBalanced
	if !balanced {
		return false, 0
	}

	balanced = abs(leftHeight-rightHeight) <= 1
	if leftHeight > rightHeight {
		height = leftHeight + 1
	} else {
		height = rightHeight + 1
	}

	return balanced, height
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
