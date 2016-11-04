package ch10q07a

// Pointers to an object of type T means that the object is being passed by
// reference to the function. However, the function will actually receive a
// copy of the pointer itself. If you want to update the pointer value from
// within the function, you need to pass a pointer to the pointer.
//
// http://stackoverflow.com/questions/8768344/what-are-pointers-to-pointers-good-for/8769095#8769095
func FindKthNode(root *Node, k *int, result **Node) {
	if root == nil || *result != nil {
		return
	}

	FindKthNode(root.Left, k, result)
	*k--
	if *k == 0 && *result == nil {
		*result = root
	}
	FindKthNode(root.Right, k, result)
}

type Node struct {
	Val         int
	Left, Right *Node
}
