package ch15q03a

import "testing"

func TestFindFirstGreaterThan(t *testing.T) {
	root := &Node{Val: 19}
	root.Left = &Node{Val: 7}
	root.Right = &Node{Val: 43}
	root.Right.Right = &Node{Val: 47}
	root.Right.Left = &Node{Val: 23}
	root.Right.Left.Right = &Node{Val: 37}
	root.Right.Left.Right.Right = &Node{Val: 41}
	root.Right.Left.Right.Left = &Node{Val: 29}
	root.Right.Left.Right.Left.Right = &Node{Val: 31}

	result := FindFirstGreaterThan(root, 23)

	if result != 29 {
		t.Errorf("Expected result to be 29 but got %v", result)
	}
}
