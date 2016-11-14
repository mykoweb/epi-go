package ch10q10q

func RecreateBT(inOrder, preOrder []string, root *Node) {
	if len(preOrder) < 1 {
		return
	}

	rootVal := preOrder[0]
	rootIndex := 0
	for i, val := range inOrder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	root.Val = rootVal
	root.Left = &Node{}
	root.Right = &Node{}

	RecreateBT(inOrder[:rootIndex], preOrder[1:rootIndex+1], root.Left)
	RecreateBT(inOrder[rootIndex+1:], preOrder[rootIndex+1:], root.Right)

	if root.Left.Val == "" {
		root.Left = nil
	}
	if root.Right.Val == "" {
		root.Right = nil
	}
}

type Node struct {
	Val         string
	Left, Right *Node
}
