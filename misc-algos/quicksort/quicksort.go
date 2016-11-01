package quicksort

import (
	"math/rand"
	"time"
)

func Quicksort(array []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := partition(array, l, r)

	Quicksort(array, l, pivot-1)
	Quicksort(array, pivot+1, r)
}

func partition(array []int, l, r int) int {
	pivot := randomInt(l, r)

	if pivot != l {
		array[pivot], array[l] = array[l], array[pivot]
	}
	pivotVal := array[l]

	i := l + 1
	for j := l + 1; j <= r; j++ {
		if array[j] < pivotVal {
			if i != j {
				array[j], array[i] = array[i], array[j]
			}
			i++
		}
	}
	array[l], array[i-1] = array[i-1], array[l]
	return i - 1
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
