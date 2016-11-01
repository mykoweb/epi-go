package ch06q01a

import (
	"math/rand"
	"time"
)

func DutchFlagSort(a []int, l, r int) {
	if l >= r {
		return
	}

	pivot := DutchFlagPartition(a, l, r)

	DutchFlagSort(a, l, pivot-1)
	DutchFlagSort(a, pivot+1, r)
}

func DutchFlagPartition(a []int, l int, r int) int {
	pivot := randomInt(l, r)
	pivotVal := a[pivot]

	// swap pivot into the first element
	if pivot != l {
		a[pivot], a[l] = a[l], a[pivot]
	}

	i := l + 1 // less than pointer
	j := l + 1 // equal pointer
	k := r     // greater than pointer

	for j <= k {
		if a[j] < pivotVal {
			if i != j {
				a[j], a[i] = a[i], a[j]
			}
			i++
			j++
		} else if a[j] == pivotVal {
			j++
		} else {
			a[j], a[k] = a[k], a[j]
			k--
		}
	}

	a[i-1], a[l] = a[l], a[i-1]
	return i - 1
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
