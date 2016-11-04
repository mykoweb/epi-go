package ch06q12a

import (
	"math/rand"
	"time"
)

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
