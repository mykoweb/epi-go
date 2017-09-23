package ch12q04a

// FindMinInCyclicalArray returns the index of the smallest element in a
// cyclically sorted array. Assume that all elements are distinct.
func FindMinInCyclicalArray(arr []int) int {
	left := 0
	right := len(arr) - 1

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > arr[len(arr)-1] {
			left = mid + 1
		} else { // arr[mid] < arr[len(arr)-1]
			right = mid
		}
	}

	return left
}
