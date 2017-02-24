package ch15q03a

import "math"

type Node struct {
	Val         int
	Left, Right *Node
}

// FindFirstGreaterThan takes a BST and a val and returns the first key that is
// greater than val that would appear in an in-order traversal of the BST.
func FindFirstGreaterThan(root *Node, val int) int {
	result := math.MinInt64
	current := root

	for current != nil {
		if current.Val > val {
			result = current.Val
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return result
}
