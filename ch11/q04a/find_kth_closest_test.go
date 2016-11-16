package ch11q04a

import "testing"

func TestFindKthClosest(t *testing.T) {
	kthClosest := FindKthClosest("./stars.txt", 3)

	if kthClosest.Size() != 3 {
		t.Errorf("Expected maxHeap size to be 3 but got %v", kthClosest.Size())
	}

	maxStar := kthClosest.ExtractMax()
	if maxStar.Name() != "h" {
		t.Errorf("Expected to get star h but got star %v", maxStar.Name())
	}
	maxStar = kthClosest.ExtractMax()
	if maxStar.Name() != "f" {
		t.Errorf("Expected to get star f but got star %v", maxStar.Name())
	}
	maxStar = kthClosest.ExtractMax()
	if maxStar.Name() != "e" {
		t.Errorf("Expected to get star e but got star %v", maxStar.Name())
	}

	if kthClosest.Size() != 0 {
		t.Errorf("Expected maxHeap size to be 0 but got %v", kthClosest.Size())
	}
}
