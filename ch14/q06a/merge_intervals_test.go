package ch14q06a

import (
	"fmt"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	var sortedIntervals = []Interval{
		Interval{-4, -1}, Interval{0, 2}, Interval{3, 6},
		Interval{7, 9}, Interval{11, 12}, Interval{14, 17},
	}

	result := mergeIntervals(sortedIntervals, Interval{1, 8})

	if len(result) != 4 {
		t.Errorf("Expected result to have length 4 but had length %v", len(result))
	}

	if result[1][0] != 0 {
		t.Errorf("Expected merged interval first val to be 0 but was %v", result[1][0])
	}
	if result[1][1] != 9 {
		t.Errorf("Expected merged interval second val to be 9 but was %v", result[1][1])
	}
	fmt.Println("The result is", result)
}
