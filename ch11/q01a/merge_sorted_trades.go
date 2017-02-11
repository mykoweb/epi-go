package ch11q01a

import "container/heap"

type Trade struct {
	time   int
	symbol string
	shares int
	prices int // represented as US cents
}

// Given several slices of Trade slices sorted in terms of increasing time,
// merge these slices into a single slice. Assume each of the sorted Trade
// slices are of the same length.
func MergeSortedTrades(sortedTrades [][]Trade) []Trade {
	result := []Trade{}
	numSlices := len(sortedTrades)
	minHeap := make(MinHeap, numSlices)

	// Initialize minHeap with the first element from each of the sorted
	// trade slices.
	for i := 0; i < numSlices; i++ {
		minHeap[i] = sortedTrades[i][0]
	}
	heap.Init(&minHeap)

	for i := 1; i < len(sortedTrades[0]); i++ {
		for j := 0; j < numSlices; j++ {
			trade := heap.Pop(&minHeap).(Trade)
			result = append(result, trade)
			heap.Push(&minHeap, sortedTrades[j][i])
		}
	}

	// minHeap still has len(minHeap) Trades left to pop. Pop them and append to
	// the result.
	for _ = range minHeap {
		trade := heap.Pop(&minHeap).(Trade)
		result = append(result, trade)
	}

	return result
}

type MinHeap []Trade

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	trade := x.(Trade)
	*h = append(*h, trade)
}

func (h *MinHeap) Pop() interface{} {
	trade := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return trade
}
