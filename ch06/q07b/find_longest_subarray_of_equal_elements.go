package ch07q07b

import "math"

// FindLongestSubarrayOfEqualElements returns the length of the longest subarray
// of equal elements.
func FindLongestSubarrayOfEqualElements(arr []int64) int {
	if len(arr) < 1 {
		return 0
	}

	var previousVal int64 = math.MaxInt64
	currLength := 0
	maxLength := 1
	previousValIndex := 0

	for i, val := range arr {
		if val != previousVal {
			previousVal = val
			previousValIndex = i
		} else {
			currLength = i - previousValIndex + 1
			if currLength > maxLength {
				maxLength = currLength
			}
		}
	}

	return maxLength
}
