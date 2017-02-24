package ch16q04a

import (
	"fmt"
	"testing"
)

func TestGeneratePowerSet(t *testing.T) {
	arr := []int{1, 2, 3}
	result := &[][]int{}

	GeneratePowerSet(arr, result)

	if len(*result) != 8 {
		t.Errorf("Expected result length to be 8 but got %v ", len(*result))
	}
	fmt.Println("result is ", *result) // result is  [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]
}
