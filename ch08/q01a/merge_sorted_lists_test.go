package ch08q01a

import "testing"

func TestMergeSortedLists(t *testing.T) {
	list1 := &Node{val: 1}
	list1.next = &Node{val: 3}
	list1.next.next = &Node{val: 5}

	list2 := &Node{val: 0}
	list2.next = &Node{val: 2}
	list2.next.next = &Node{val: 4}
	list2.next.next.next = &Node{val: 6}
	list2.next.next.next.next = &Node{val: 7}

	mergedList := MergeSortedLists(list1, list2)

	index := 0
	for mergedList != nil {
		if mergedList.Val() != index {
			t.Errorf("Expected node val of %v but got %v", index, mergedList.Val())
		}

		mergedList = mergedList.next
		index += 1
	}
}
