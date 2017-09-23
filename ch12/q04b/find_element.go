package ch12q04b

// FindElement returns the index of k in a cyclically sorted array. Assume that
// all elements are distinct.
func FindElement(arr []int, k int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		if k == arr[mid] {
			return mid
		} else if k == arr[len(arr)-1] {
			return len(arr) - 1
		} else if arr[mid] < arr[len(arr)-1] {
			if arr[mid] < k && k < arr[len(arr)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if arr[mid] > arr[len(arr)-1] {
			if arr[len(arr)-1] < k && k < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}
