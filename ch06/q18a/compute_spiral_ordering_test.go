package ch06q18a

import (
	"fmt"
	"testing"
)

func TestComputeSpiralOrdering(t *testing.T) {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	result := ComputeSpiralOrdering(matrix)

	fmt.Println("The result is", result) // The result is [1 2 3 6 9 8 7 4 5]

	matrix = [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}

	result = ComputeSpiralOrdering(matrix)

	fmt.Println("The result is", result) // The result is [1 2 3 4 8 12 16 15 14 13 9 5 6 7 11 10]
}
