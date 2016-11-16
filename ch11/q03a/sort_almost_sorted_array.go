package ch11q03a

import "errors"

// Given an array and a guarrantee that any element is no more than k positions
// away from where it would be if the array were sorted, sort the array.
func SortAlmostSortedArray(arr []int, k int) {
	if k >= len(arr) {
		return
	}

	minHeap := MinHeap{}
	for i := 0; i < k; i++ {
		minHeap.Insert(arr[i])
	}

	for i := k; i < len(arr); i++ {
		minHeap.Insert(arr[i])
		arr[i-k], _ = minHeap.ExtractMin()
	}

	for i := len(arr) - k; i < len(arr); i++ {
		arr[i], _ = minHeap.ExtractMin()
	}
}

type MinHeap struct {
	heap []int
	size int
}

func (h *MinHeap) Insert(val int) {
	h.heap = append(h.heap, val)
	h.size++

	currIndex := len(h.heap) - 1
	parentIndex := h.parentIndex(currIndex)

	for h.heap[parentIndex] > h.heap[currIndex] {
		h.heap[parentIndex], h.heap[currIndex] = h.heap[currIndex], h.heap[parentIndex]
		currIndex = parentIndex
		parentIndex = h.parentIndex(currIndex)
	}
}

func (h *MinHeap) ExtractMin() (result int, err error) {
	if h.size <= 0 {
		err = errors.New("ExtractMin() failed on an empty heap")
		// errors.New("ExtractMin() failed on an empty heap")
	} else if h.size == 1 {
		result = h.heap[0]
		h.heap = []int{}
	} else {
		result = h.heap[0]
		h.heap[0], h.heap[len(h.heap)-1] = h.heap[len(h.heap)-1], h.heap[0]
		h.heap = h.heap[:len(h.heap)-1]

		if len(h.heap) > 1 {
			currIndex := 0
			leftIndex, rightIndex := h.childIndices(currIndex)

			for h.heap[leftIndex] < h.heap[currIndex] || h.heap[rightIndex] < h.heap[currIndex] {
				if h.heap[leftIndex] < h.heap[rightIndex] {
					h.heap[currIndex], h.heap[leftIndex] = h.heap[leftIndex], h.heap[currIndex]
					currIndex = leftIndex
				} else {
					h.heap[currIndex], h.heap[rightIndex] = h.heap[rightIndex], h.heap[currIndex]
					currIndex = rightIndex
				}
				leftIndex, rightIndex = h.childIndices(currIndex)
				if leftIndex == -1 {
					break
				}
			}
		}
	}

	h.size--
	return
}

func (h *MinHeap) Min() int {
	return h.heap[0]
}

func (h *MinHeap) Size() int {
	return h.size
}

// If index is 0 then return 0, otherwise find the parent index.
func (h *MinHeap) parentIndex(index int) (result int) {
	if index < 1 {
		result = 0
	} else {
		if index%2 == 0 {
			result = index/2 - 1
		} else {
			result = index / 2
		}
	}
	return
}

// Return both child indices if they are within the array bounds. If only
// the right child index is out of bounds then set the right child index to the
// left child index. If both are out of bounds then return -1 for both child
// indices.
func (h *MinHeap) childIndices(index int) (leftIndex, rightIndex int) {
	leftIndexVal := index*2 + 1
	rightIndexVal := index*2 + 2

	if leftIndexVal > len(h.heap)-1 {
		leftIndex = -1
	} else {
		leftIndex = leftIndexVal
	}
	if rightIndexVal > len(h.heap)-1 {
		if leftIndex != -1 {
			rightIndex = leftIndex
		} else {
			rightIndex = -1
		}
	} else {
		rightIndex = rightIndexVal
	}

	return
}
