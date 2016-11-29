package ch06q08a

import "math"

// BuyAndSellTwice returns the maximum profit you can get if you buy and sell
// a given stock no more than twice within a given date range. An array is
// provided whose elements represent the starting price of a stock on a given
// day. For array indices i and j where i < j, then arr[i] must occur before
// arr[j]. The second buy must occur on the day after the first sell.
func BuyAndSellTwice(arr []int64) int64 {
	if len(arr) < 2 {
		return 0
	}

	var currDiff int64
	var totalProfit int64
	var totalMaxProfit int64 = math.MinInt64
	firstMax := make([]int64, len(arr))
	secondMax := make([]int64, len(arr))
	minVal := arr[0]
	maxVal := arr[len(arr)-1]

	// For the first buy and sell:
	// Find the maximum profit if you sell by day i and buy anywhere from day
	// 0 until day i.
	firstMax[0] = 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < minVal {
			minVal = arr[i]
		}
		currDiff = arr[i] - minVal
		if currDiff > firstMax[i-1] {
			firstMax[i] = currDiff
		} else {
			firstMax[i] = firstMax[i-1]
		}
	}

	// For the second buy and sell:
	// Find the max profit if you buy on day i and sell anywhere from day i until
	// day n-1, where n is the length of the array.
	secondMax[len(arr)-1] = 0
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
		currDiff = maxVal - arr[i]
		if currDiff > secondMax[i+1] {
			secondMax[i] = currDiff
		} else {
			secondMax[i] = secondMax[i+1]
		}

		// Calculate the profit for both buy and sells where the first sell is on
		// day i and the second buy is on day i+1.
		totalProfit = firstMax[i] + secondMax[i+1]
		if totalProfit > totalMaxProfit {
			totalMaxProfit = totalProfit
		}
	}

	return totalMaxProfit
}
