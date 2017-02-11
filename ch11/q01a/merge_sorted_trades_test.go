package ch11q01a

import "testing"

func TestMergeSortedTrades(t *testing.T) {
	trades1 := []Trade{
		Trade{2, "AAPL", 100, 300},
		Trade{3, "INTC", 40, 1000},
		Trade{9, "GOOD", 500, 3030},
	}
	trades2 := []Trade{
		Trade{1, "BRKA", 100, 300},
		Trade{4, "PG", 25, 1000},
		Trade{7, "IBM", 500, 777},
	}
	trades3 := []Trade{
		Trade{5, "PETS", 100, 300},
		Trade{6, "CVS", 40, 1000},
		Trade{8, "FB", 500, 3030},
	}

	sortedTrades := [][]Trade{trades1, trades2, trades3}

	result := MergeSortedTrades(sortedTrades)

	if len(result) != 9 {
		t.Errorf("Expected result to have 9 trades but it had %v", len(result))
	}

	for i, trade := range result {
		if trade.time != i+1 {
			t.Errorf("Expected trade time %v but got %v", i+1, trade.time)
		}
	}
}
