package ch16q02a

import "testing"

func TestGenNQueens(t *testing.T) {
	result := GenNQueens(4)
	if len(result) != 2 {
		t.Errorf("Expected length of 2, but got %v", len(result))
	}
}
