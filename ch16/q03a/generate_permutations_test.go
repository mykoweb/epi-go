package ch16q03a

import (
	"sort"
	"testing"
)

func TestGeneratePermutations(t *testing.T) {
	var result [][]int
	slice := []int{7, 4, 2, 1}

	GeneratePermutations(0, &slice, &result)

	result[0], result[1] = result[1], result[0]
	sort.Sort(twoDSlice(result))

	if len(result) != factorial(len(slice)) {
		t.Errorf("Expected number of results to be %v, but got %v", factorial(len(slice)), len(result))
	}

	// Verify first permutation is []int{1, 2, 4, 7}
	for i, val := range []int{1, 2, 4, 7} {
		if result[0][i] != val {
			t.Errorf("Expected first permutation index %v to equal %v but got %v", i, val, result[0][i])
		}
	}

	// Verify second permutation is []int{1, 2, 7, 4}
	for i, val := range []int{1, 2, 7, 4} {
		if result[1][i] != val {
			t.Errorf("Expected second permutation index %v to equal %v but got %v", i, val, result[1][i])
		}
	}

	// Verify last permutation is []int{7, 4, 1, 1}
	for i, val := range []int{7, 4, 2, 1} {
		if result[len(result)-1][i] != val {
			t.Errorf("Expected second permutation index %v to equal %v but got %v", i, val, result[len(result)-1][i])
		}
	}
}

type twoDSlice [][]int

func (s twoDSlice) Len() int      { return len(s) }
func (s twoDSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s twoDSlice) Less(i, j int) bool {
	for index, _ := range s[i] {
		if s[i][index] < s[j][index] {
			return true
		} else if s[i][index] > s[j][index] {
			return false
		}
	}
	return false
}

// Assume n is always positive, otherwise return -1.
func factorial(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
