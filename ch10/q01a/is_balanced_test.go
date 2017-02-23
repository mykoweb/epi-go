package ch10q01a

import "testing"

func TestIsBalanced(t *testing.T) {
	bt := &Node{}
	bt.left = &Node{}
	bt.left.left = &Node{}
	bt.right = &Node{}

	balanced, _ := IsBalanced(bt)

	if balanced == false {
		t.Errorf("IsBalanced expected to return true for balanced binary tree but it did not")
	}

	bt.left.left.left = &Node{}

	balanced, _ = IsBalanced(bt)

	if balanced == true {
		t.Errorf("IsBalanced expected to return false for unbalanced binary tree but it did not")
	}
}
