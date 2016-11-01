package ch19q05a

type Node struct {
	val   int
	edges []*Node
}

func CloneGraph(root *Node) *Node {
	var queue []*Node
	visitedNodes := make(map[*Node]*Node)

	visitedNodes[root] = &Node{val: root.val}
	queue = pushQueue(queue, root)

	for len(queue) > 0 {
		var poppedNode *Node
		poppedNode, queue = popQueue(queue)
		for _, currentNode := range poppedNode.edges {
			if visitedNodes[currentNode] == nil {
				visitedNodes[currentNode] = &Node{val: currentNode.val}
				queue = pushQueue(queue, currentNode)
				visitedNodes[poppedNode].edges = append(visitedNodes[poppedNode].edges, visitedNodes[currentNode])
			}
		}
	}

	return visitedNodes[root]
}

func pushQueue(queue []*Node, node *Node) []*Node {
	return append([]*Node{node}, queue...)
}

func popQueue(queue []*Node) (node *Node, newQueue []*Node) {
	node, newQueue = queue[len(queue)-1], queue[:len(queue)-1]
	return
}
