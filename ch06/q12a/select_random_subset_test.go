package ch06q12a

import "testing"

func TestSelectRandSubset(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	randSubArr := SelectRandSubset(arr, 3)
	if len(randSubArr) != 3 {
		t.Errorf("Expected length to be 3 but got", len(randSubArr))
	}

	randSubArr = SelectRandSubset(arr, 6)
	if len(randSubArr) != 6 {
		t.Errorf("Expected length to be 6 but got", len(randSubArr))
	}
}
