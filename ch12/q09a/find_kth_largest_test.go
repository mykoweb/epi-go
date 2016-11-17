package ch12q09a

import "testing"

func TestFindKthLargest(t *testing.T) {
	var result int
	var expectedResult int
	var arr []int

	for i := 1; i <= 6; i++ {
		arr = []int{3, 6, 2, 1, 5, 4}
		expectedResult = 7 - i
		result = FindKthLargest(arr, i)
		if result != expectedResult {
			t.Errorf("Expected %v but got %v", expectedResult, result)
		}
	}
}
