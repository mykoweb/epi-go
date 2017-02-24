package ch10q03a

import "testing"

func TestComputeLca(t *testing.T) {
	root := &Node{Val: 314}
	root.Right = &Node{Val: 6}
	root.Right.Left = &Node{Val: 2}
	root.Right.Left.Right = &Node{Val: 1}
	root.Right.Left.Right.Left = &Node{Val: 401}
	root.Right.Left.Right.Right = &Node{Val: 257}
	root.Right.Left.Right.Left.Right = &Node{Val: 641}
	root.Right.Right = &Node{Val: 271}

	n1 := root.Right.Left.Right.Left.Right
	n2 := root.Right.Left.Right.Right

	numFound, lca := ComputeLca(root, n1, n2)

	if numFound != 2 {
		t.Errorf("Expected numFound to equal 2 but got %v", numFound)
	}

	if lca.Val != 1 {
		t.Errorf("Expected LCA node to have a value of 1 but got %v", lca.Val)
	}
}
