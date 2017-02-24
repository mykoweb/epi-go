package ch10q03a

type Node struct {
	Val         int
	Left, Right *Node
}

// Given 2 nodes in a binary tree, compute their lower common ancestor (LCA).
func ComputeLca(root, n1, n2 *Node) (numFound int, lca *Node) {
	if root == nil {
		return
	}

	leftNumFound, leftLca := ComputeLca(root.Left, n1, n2)
	if leftNumFound == 2 {
		return leftNumFound, leftLca
	}

	rightNumFound, rightLca := ComputeLca(root.Right, n1, n2)
	if rightNumFound == 2 {
		return rightNumFound, rightLca
	}

	numFound = leftNumFound + rightNumFound
	if root == n1 {
		numFound += 1
	}
	if root == n2 {
		numFound += 1
	}

	if numFound == 2 {
		return numFound, root
	} else {
		return numFound, lca
	}
}
