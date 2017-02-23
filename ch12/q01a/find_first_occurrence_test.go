package ch12q01a

import "testing"

func TestFindFirstOccurrence(t *testing.T) {
	arr := []int{-4, -4, -1, 2, 2, 2, 5, 7, 8, 9, 9, 9, 9}

	index := FindFirstOccurrence(arr, 100)

	if index != -1 {
		t.Errorf("Expected FindFirstOccurrence to return -1 but got %v", index)
	}

	index = FindFirstOccurrence(arr, 2)

	if index != 3 {
		t.Errorf("Expected FindFirstOccurrence to return 3 but got %v", index)
	}

	index = FindFirstOccurrence(arr, 9)

	if index != 9 {
		t.Errorf("Expected FindFirstOccurrence to return 9 but got %v", index)
	}
}
