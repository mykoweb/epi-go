package ch14q01a

import "testing"

func TestFindIntersection(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	result := FindIntersection(a, b)

	if len(result) != 0 {
		t.Errorf("Expected no intersection but got one")
	}

	a = []int{4, 5, 6}
	b = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}

	result = FindIntersection(a, b)

	if len(result) != 3 {
		t.Errorf("Expected result length to be 3, but got %v", len(result))
	}

	for i := 0; i < len(result); i++ {
		if result[i] != i+4 {
			t.Errorf("Expected %v but got %v", i+4, result[i])
		}
	}
}
