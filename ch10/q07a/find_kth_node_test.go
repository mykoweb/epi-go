package ch10q07a

import "testing"

func TestFindKthNode(t *testing.T) {
	root := &Node{4, nil, nil}
	root.Left = &Node{2, nil, nil}
	root.Right = &Node{6, nil, nil}
	root.Left.Left = &Node{1, nil, nil}
	root.Left.Right = &Node{3, nil, nil}
	root.Right.Left = &Node{5, nil, nil}
	root.Right.Right = &Node{7, nil, nil}

	for i := 1; i < 8; i++ {
		var result *Node
		k := i
		FindKthNode(root, &k, &result)
		if (*result).Val != i {
			t.Errorf("Expected value of %v but got %v", i, (*result).Val)
		}

	}
}
