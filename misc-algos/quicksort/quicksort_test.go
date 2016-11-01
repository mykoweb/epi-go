package quicksort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestQuicksort(t *testing.T) {
	arrayLen := 64
	array := make([]int, arrayLen)
	for i := 0; i < arrayLen; i++ {
		array[i] = i
	}
	shuffle(array)
	// fmt.Println("Initial array is", array)

	Quicksort(array, 0, arrayLen-1)
	// fmt.Println("Sorted array is", array)

	if sort.IntsAreSorted(array) != true {
		t.Errorf("Expected array to be sorted but it was not")
	}
}

func shuffle(a []int) {
	rand.Seed(time.Now().UnixNano())
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
