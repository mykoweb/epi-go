package ch12q01a

// FindFirstOccurence returns the index of the first occurence of val in a
// sorted slice.
func FindFirstOccurrence(arr []int, val int) int {
	left := 0
	right := len(arr) - 1
	result := -1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] < val {
			left = mid + 1
		} else if arr[mid] > val {
			right = mid - 1
		} else { // arr[mid] == val
			result = mid
			right = mid - 1
		}
	}

	return result
}
