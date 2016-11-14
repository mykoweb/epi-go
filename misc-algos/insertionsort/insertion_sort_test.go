package insertionsort

import "testing"

func TestInsertionSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1, 0}

	InsertionSort(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] != i {
			t.Errorf("Expected %v but got %v", i, arr[i])
		}
	}
}
