package ch06q08a

import "testing"

func TestBuyAndSellTwice(t *testing.T) {
	arr := []int64{12, 11, 13, 9, 12, 8, 14, 13, 15}
	result := BuyAndSellTwice(arr)
	if result != 10 {
		t.Errorf("Expected max profit of 10 but got %v", result)
	}

	arr = []int64{8, 7, 6, 5}
	result = BuyAndSellTwice(arr)
	if result != 0 {
		t.Errorf("Expected max profit of 0 but got %v", result)
	}
}
