package ch15q04a

import (
	"fmt"
	"testing"
)

func TestFindKLargest(t *testing.T) {
	root := &Node{Val: 19}
	root.Left = &Node{Val: 7}
	root.Left.Left = &Node{Val: 3}
	root.Left.Right = &Node{Val: 11}
	root.Right = &Node{Val: 43}
	root.Right.Right = &Node{Val: 47}
	root.Right.Right.Right = &Node{Val: 53}
	root.Right.Left = &Node{Val: 23}
	root.Right.Left.Right = &Node{Val: 37}
	root.Right.Left.Right.Right = &Node{Val: 41}
	root.Right.Left.Right.Left = &Node{Val: 29}
	root.Right.Left.Right.Left.Right = &Node{Val: 31}

	kLargest := []int{}
	FindKLargest(root, 5, &kLargest)
	fmt.Println("k largest is", kLargest) // k largest is [53 47 43 41 37]
}
