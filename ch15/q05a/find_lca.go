package ch15q05a

type Node struct {
	Val         int
	Left, Right *Node
}

// FindLca returns a pointer to the node that is the lowest common ancestor
// (LCA) of two nodes, n1 and n2, in a BST. Assume the values in each node are
// distinct. Also assume that n1.Val < n2.Val.
func FindLca(root, n1, n2 *Node) *Node {
	if root == nil {
		return nil
	}

	var result *Node

	if n1.Val == root.Val || n2.Val == root.Val {
		return root
	} else if n1.Val < root.Val && root.Val < n2.Val {
		return root
	} else if n1.Val < root.Val && n2.Val < root.Val {
		result = FindLca(root.Left, n1, n2)
	} else {
		result = FindLca(root.Right, n1, n2)
	}

	return result
}
