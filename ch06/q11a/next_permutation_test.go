package ch06q11a

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	perm := []int{6, 2, 1, 5, 4, 3, 0}
	result := NextPermutation(perm)
	fmt.Println("Result is", result)
	expected_result := []int{6, 2, 3, 0, 1, 4, 5}

	for i := 0; i < len(result); i++ {
		if result[i] != expected_result[i] {
			t.Errorf("Expected %v but got %v", expected_result[i], result[i])
		}
	}

	perm = []int{3, 2, 1} // Largest permutation
	result = NextPermutation(perm)
	if len(result) != 0 {
		t.Errorf("Expected empty slice but did not get this")
	}
}
