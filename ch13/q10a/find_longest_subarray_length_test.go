package ch13q10a

import "testing"

func TestFindLongestSubarrayLength(t *testing.T) {
	arr := []rune{'f', 's', 'f', 'e', 't', 'w', 'e', 'n', 'w', 'e'}

	result := FindLongestSubarrayLength(arr)

	if result != 5 {
		t.Errorf("Expected the longest subarray length to be 5 but it was %v", result)
	}
}
