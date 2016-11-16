package ch11q03a

import (
	"sort"
	"testing"
)

func TestSortAlmostSortedArray(t *testing.T) {
	arr := []int{3, -1, 2, 6, 4, 5, 8}
	SortAlmostSortedArray(arr, 2)

	if sort.IntsAreSorted(arr) != true {
		t.Errorf("Expected array %v to be sorted but it was not", arr)
	}
}
