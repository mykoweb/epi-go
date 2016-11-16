package ch11q04a

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func FindKthClosest(filename string, k int) MaxHeap {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var name string
	var x, y, z float64
	var newStar Star
	i := 0
	maxHeap := MaxHeap{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmpResult := strings.Split(scanner.Text(), ",")
		name = tmpResult[0]
		x, err = strconv.ParseFloat(tmpResult[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		y, _ = strconv.ParseFloat(tmpResult[2], 64)
		z, _ = strconv.ParseFloat(tmpResult[3], 64)
		newStar = NewStar([3]float64{x, y, z}, name)
		if i < k {
			maxHeap.Insert(newStar)
			i++
		} else {
			if newStar.Dist() < maxHeap.MaxDist() {
				maxHeap.ExtractMax()
				maxHeap.Insert(newStar)
			}
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return maxHeap
}

type MaxHeap struct {
	heap []Star
	size int
}

func (h *MaxHeap) Insert(s Star) {
	h.heap = append(h.heap, s)
	h.size++
	currIndex := len(h.heap) - 1
	parentIndex := h.parentIndex(currIndex)
	for h.heap[parentIndex].Dist() < h.heap[currIndex].Dist() {
		h.heap[parentIndex], h.heap[currIndex] = h.heap[currIndex], h.heap[parentIndex]
		currIndex = parentIndex
		parentIndex = h.parentIndex(currIndex)
	}
}

func (h *MaxHeap) ExtractMax() Star {
	if h.Size() < 1 {
		return Star{}
	}

	result := NewStar(h.heap[0].Coord(), h.heap[0].Name())
	if h.Size() == 1 {
		h.heap = []Star{}
		h.size = 0
		return result
	}

	h.heap[len(h.heap)-1], h.heap[0] = h.heap[0], h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.size--
	currIndex := 0
	leftChild, rightChild := h.childIndices(currIndex)
	for h.heap[leftChild].Dist() > h.heap[currIndex].Dist() || h.heap[rightChild].Dist() > h.heap[currIndex].Dist() {
		if h.heap[leftChild].Dist() > h.heap[rightChild].Dist() {
			h.heap[currIndex], h.heap[leftChild] = h.heap[leftChild], h.heap[currIndex]
			currIndex = leftChild
		} else {
			h.heap[currIndex], h.heap[rightChild] = h.heap[rightChild], h.heap[currIndex]
			currIndex = rightChild
		}
		leftChild, rightChild = h.childIndices(currIndex)
	}

	return result
}

func (h *MaxHeap) MaxDist() float64 {
	return h.heap[0].Dist()
}

func (h *MaxHeap) Size() int {
	return h.size
}

// If curr is 0 then return 0, otherwise find the parent index.
func (h *MaxHeap) parentIndex(curr int) int {
	if curr == 0 {
		return 0
	}
	if curr%2 == 0 {
		return curr/2 - 1
	} else {
		return curr / 2
	}
}

// Return both child indices if they are within the array bounds. If the heap
// array length is 1 or less, then return an index of 0 for both child indices.
// If either of the child indices goes out of bounds cap it to the max index
// of the heap array.
func (h *MaxHeap) childIndices(parentIndex int) (leftIndex, rightIndex int) {
	if h.Size() <= 1 {
		return 0, 0
	}
	leftIndex = 2*parentIndex + 1
	rightIndex = 2*parentIndex + 2
	heapLength := len(h.heap)
	if leftIndex >= heapLength {
		leftIndex = heapLength - 1
	}
	if rightIndex >= heapLength {
		rightIndex = heapLength - 1
	}
	return
}

type Star struct {
	coord [3]float64
	dist  float64
	name  string
}

func NewStar(coord [3]float64, name string) Star {
	s := Star{}
	s.coord = coord
	s.dist = math.Sqrt(math.Pow(coord[0], 2) + math.Pow(coord[1], 2) + math.Pow(coord[2], 2))
	s.name = name
	return s
}

func (s *Star) Coord() [3]float64 {
	return s.coord
}

func (s *Star) Dist() float64 {
	return s.dist
}

func (s *Star) Name() string {
	return s.name
}
