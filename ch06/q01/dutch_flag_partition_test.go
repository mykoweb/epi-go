package ch06q01a

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestDutchFlagPartition(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 2, 2, 2, 2, 2, 2, 2, 2}
	shuffle(arr)

	pivot := DutchFlagPartition(arr, 0, len(arr)-1)
	pivotVal := arr[pivot]

	for i := 0; i < pivot; i++ {
		if arr[i] > pivotVal {
			t.Errorf("Array was not partitioned correctly")
		}
	}

	for i := pivot + 1; i < len(arr); i++ {
		if arr[i] < pivotVal {
			t.Errorf("Array was not partitioned correctly")
		}
	}
}

func TestDutchFlagSort(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 2, 2, 2, 2, 2, 2, 2, 2}
	shuffle(arr)

	DutchFlagSort(arr, 0, len(arr)-1)

	if sort.IntsAreSorted(arr) != true {
		t.Errorf("Expected array to be sorted but it was not")
	}
}

func shuffle(a []int) {
	rand.Seed(time.Now().UnixNano())
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
