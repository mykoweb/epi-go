package ch12q04a

import "testing"

func TestFindMinInCyclicalArray(t *testing.T) {
	arr := []int{7, 8, 9, 1, 2, 3, 4, 5, 6}

	result := FindMinInCyclicalArray(arr)

	if arr[result] != 1 {
		t.Errorf("FindMinInCyclicalArray failed to find the min")
	}
}
