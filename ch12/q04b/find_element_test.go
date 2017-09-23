package ch12q04b

import "testing"

func TestFindElement(t *testing.T) {
	arr := []int{7, 8, 9, 1, 2, 3, 4, 5, 6}

	result := FindElement(arr, 2)
	if arr[result] != 2 {
		t.Errorf("FindMinInCyclicalArray failed to find the min")
	}

	result = FindElement(arr, 8)
	if arr[result] != 8 {
		t.Errorf("FindMinInCyclicalArray failed to find the min")
	}

	result = FindElement(arr, 1)
	if arr[result] != 1 {
		t.Errorf("FindMinInCyclicalArray failed to find the min")
	}

	result = FindElement(arr, 7)
	if arr[result] != 7 {
		t.Errorf("FindMinInCyclicalArray failed to find the min")
	}

	result = FindElement(arr, 6)
	if arr[result] != 6 {
		t.Errorf("FindMinInCyclicalArray failed to find the min")
	}

	result = FindElement(arr, 9)
	if arr[result] != 9 {
		t.Errorf("FindMinInCyclicalArray failed to find the min")
	}
}
