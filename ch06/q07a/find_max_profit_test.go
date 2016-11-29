package ch06q07a

import "testing"

func TestFindMaxProfit(t *testing.T) {
	arr := []int64{100, 80, 40, 0}
	result, _ := FindMaxProfit(arr)
	if result != 0 {
		t.Errorf("Expected max profit of 0 but got %v", result)
	}

	arr = []int64{3, 2, 1, 5, 10, 15, 5, 25, 30, 22, 22, 23}
	result, _ = FindMaxProfit(arr)
	if result != 29 {
		t.Errorf("Expected max profit of 29 but got %v", result)
	}
}
