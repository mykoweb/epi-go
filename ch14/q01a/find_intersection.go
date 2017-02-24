package ch14q01a

// FindIntersection returns the non-duplicated intersection of the 2 sorted
// arrays, a and b. This code is optimized for the case where array a is much
// shorter in length than array b. The runtime is O(m log n) where m is the
// length of a and n is the length of b.
func FindIntersection(a, b []int) []int {
	result := []int{}

	for i := 0; i < len(a); i++ {
		if boolBinarySearch(b, a[i]) {
			if len(result) == 0 || result[len(result)-1] != a[i] {
				result = append(result, a[i])
			}
		}
	}

	return result
}

// boolBinarySearch returns true if val was found in the slice, false otherwise.
func boolBinarySearch(arr []int, val int) bool {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] < val {
			left = mid + 1
		} else if arr[mid] > val {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}
