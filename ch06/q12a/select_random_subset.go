package ch06q12a

import (
	"math/rand"
	"time"
)

// SelectRandSubset takes an integer slice with unique elements and an integer,
// k, where k < len(arr), and returns a subset of k elements selected randomly
// from the integer slice.
//
// The time complexity is O(k).
func SelectRandSubset(arr []int, k int) []int {
	if k >= len(arr) {
		return arr
	}

	var index int
	localArr := make([]int, len(arr))
	copy(localArr, arr)

	for i := 0; i < k; i++ {
		index = randSelect(i, len(localArr))
		localArr[i], localArr[index] = localArr[index], localArr[i]
	}

	return localArr[:k]
}

func randSelect(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
