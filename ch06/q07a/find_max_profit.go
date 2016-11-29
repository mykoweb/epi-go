package ch06q07a

import (
	"errors"
	"math"
)

// FindMaxProfit finds the maximum profit based on the stock prices provided
// from the given array. Each element in the array represents the starting
// stock price on a given day. For indices i and j where i < j, the day for the
// stock price in arr[i] occurs BEFORE the day for the stock price in arr[j].
//
// This is not stated in the original problem, but assume we can both buy and
// sell on the same day.
//
// The time complexity is O(n) where n is the length of the array.
func FindMaxProfit(arr []int64) (int64, error) {
	if len(arr) < 2 {
		return 0, errors.New("FindMaxProfit: needs at least 2 elements in array")
	}

	var minVal int64 = arr[0]
	var maxProfit int64 = math.MinInt64
	var currDiff int64

	for _, val := range arr {
		if val < minVal {
			minVal = val
		}
		currDiff = val - minVal
		if currDiff > maxProfit {
			maxProfit = currDiff
		}
	}

	return maxProfit, nil
}
