package ch10q10q

import "testing"

func TestRecreateBT(t *testing.T) {
	root := &Node{}

	inOrder := []string{"F", "B", "A", "E", "H", "C", "D", "I", "G"}
	preOrder := []string{"H", "B", "F", "E", "A", "C", "D", "G", "I"}
	RecreateBT(inOrder, preOrder, root)

	if root.Val != "H" {
		t.Errorf("Expected H but got %v", root.Val)
	}
	if root.Left.Val != "B" {
		t.Errorf("Expected B but got %v", root.Left.Val)
	}
	if root.Left.Left.Val != "F" {
		t.Errorf("Expected F but got %v", root.Left.Left.Val)
	}
	if root.Left.Left.Left != nil {
		t.Errorf("Expected nil but got %v", root.Left.Left.Left)
	}
	if root.Left.Left.Right != nil {
		t.Errorf("Expected nil but got %v", root.Left.Left.Right)
	}
	if root.Left.Right.Val != "E" {
		t.Errorf("Expected E but got %v", root.Left.Right.Val)
	}
	if root.Left.Right.Left.Val != "A" {
		t.Errorf("Expected A but got %v", root.Left.Right.Left.Val)
	}
	if root.Right.Val != "C" {
		t.Errorf("Expected C but got %v", root.Right.Val)
	}
}
