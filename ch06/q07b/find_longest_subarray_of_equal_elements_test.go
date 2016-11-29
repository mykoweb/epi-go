package ch07q07b

import "testing"

func TestFindLongestSubarrayOfEqualElements(t *testing.T) {
	arr := []int64{1, 2, 3, 4, 7, 7, 7, 5, 4, 4, 0, 1}
	result := FindLongestSubarrayOfEqualElements(arr)
	if result != 3 {
		t.Errorf("Expected result of 3 but got %v", result)
	}

	arr = []int64{9, 8, 7}
	result = FindLongestSubarrayOfEqualElements(arr)
	if result != 1 {
		t.Errorf("Expected result of 1 but got %v", result)
	}

	arr = []int64{7, 6, 5, 5, 5, 4, 3, 2, 2, 2, 2, 2, 2, 3, 3, 4, 5}
	result = FindLongestSubarrayOfEqualElements(arr)
	if result != 6 {
		t.Errorf("Expected result of 6 but got %v", result)
	}
}
