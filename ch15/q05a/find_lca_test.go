package ch15q05a

import "testing"

func TestFindLca(t *testing.T) {
	root := &Node{Val: 19}
	root.Left = &Node{Val: 7}
	root.Left.Left = &Node{Val: 3}
	root.Left.Left.Left = &Node{Val: 2}
	root.Left.Left.Right = &Node{Val: 5}
	root.Left.Right = &Node{Val: 11}
	root.Left.Right.Right = &Node{Val: 17}
	root.Left.Right.Right.Left = &Node{Val: 13}
	root.Right = &Node{Val: 43}
	root.Right.Left = &Node{Val: 23}

	n1 := root.Left.Left
	n2 := root.Left.Right.Right

	result := FindLca(root, n1, n2)
	if result.Val != 7 {
		t.Errorf("FindLca returned the wrong node")
	}

	n2 = root
	result = FindLca(root, n1, n2)
	if result.Val != 19 {
		t.Errorf("FindLca returned the wrong node")
	}

	n2 = root.Right.Left
	result = FindLca(root, n1, n2)
	if result.Val != 19 {
		t.Errorf("FindLca returned the wrong node")
	}
}
