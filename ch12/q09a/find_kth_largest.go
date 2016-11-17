package ch12q09a

import (
	"math/rand"
	"time"
)

// Finds the kth largest element in arr. arr is an array of distinct integers.
// k starts from 1, so k = 1 will find the largest element in arr.
func FindKthLargest(arr []int, k int) (result int) {
	left := 0
	right := len(arr) - 1

	var pivotIndex int
	for left <= right {
		pivotIndex = partition(arr, left, right)
		if pivotIndex == k-1 {
			result = arr[pivotIndex]
			return
		} else if pivotIndex < k-1 {
			left = pivotIndex + 1
		} else { // pivotIndex > k-1
			right = pivotIndex - 1
		}
	}

	return
}

// Partitions the array between indices left and right, inclusive, and returns
// the pivot index. All elements before the pivot are greater than the pivot.
func partition(arr []int, left, right int) int {
	if left == right {
		return left
	} else if left > right {
		return -1
	}

	pivotIndex := randInt(left, right)
	pivot := arr[pivotIndex]
	arr[left], arr[pivotIndex] = arr[pivotIndex], arr[left]
	i := left + 1
	for j := left + 1; j <= right; j++ {
		if arr[j] > pivot {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i-1], arr[left] = arr[left], arr[i-1]
	return i - 1
}

// Returns a random integer between min and max.
func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
