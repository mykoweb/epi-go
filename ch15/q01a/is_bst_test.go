package ch15q01a

import (
	"math"
	"testing"
)

func TestIsBST(t *testing.T) {
	var low, high int64
	low = math.MinInt64
	high = math.MaxInt64

	bst := &Node{4, nil, nil}
	bst.Left = &Node{2, nil, nil}
	bst.Left.Left = &Node{1, nil, nil}
	bst.Left.Right = &Node{3, nil, nil}
	bst.Right = &Node{6, nil, nil}
	bst.Right.Left = &Node{5, nil, nil}
	bst.Right.Right = &Node{7, nil, nil}

	result := isBST(bst, low, high)
	if result != true {
		t.Errorf("Expected result to be true but got false")
	}

	bst.Left.Right.Val = 10000
	result = isBST(bst, low, high)
	if result != false {
		t.Errorf("Expected result to be false but got true")
	}
}
