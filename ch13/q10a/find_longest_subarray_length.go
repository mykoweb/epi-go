package ch13q10a

type indexPair struct{ l, r int }

// FindLongestSubarrayLength finds the length of the longest subarray within
// arr that has distinct elements. For now we are only dealing with rune slices
// but we will update to take slices of any type.
func FindLongestSubarrayLength(arr []rune) int {
	if len(arr) < 1 {
		return 0
	}

	currSubarrayIndices := indexPair{}
	currLength := 0
	longestLength := 0
	foundRunes := make(map[rune]int)

	for i := 0; i < len(arr); i++ {
		currSubarrayIndices.r = i

		if _, ok := foundRunes[arr[i]]; ok {
			currSubarrayIndices.l = foundRunes[arr[i]] + 1
		} else {
			foundRunes[arr[i]] = i
			currLength = currSubarrayIndices.r - currSubarrayIndices.l + 1
			if currLength > longestLength {
				longestLength = currLength
			}
		}
	}

	return longestLength
}
