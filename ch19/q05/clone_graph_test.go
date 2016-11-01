package ch19q05a

import "testing"

func TestCloneGraph(t *testing.T) {
	graph := &Node{val: 1}
	graph.edges = append(graph.edges, &Node{val: 2})
	graph.edges = append(graph.edges, &Node{val: 3})
	graph.edges[1].edges = append(graph.edges[1].edges, &Node{val: 4})

	newGraph := CloneGraph(graph)

	if graph.val != newGraph.val {
		t.Errorf("Root node expected to have value %v but got %v", graph.val, newGraph.val)
	}

	for i := 0; i < 2; i++ {
		if graph.edges[i] == newGraph.edges[i] {
			t.Errorf("Nodes should not have the same address between original and cloned graphs")
		}
		if graph.edges[i].val != newGraph.edges[i].val {
			t.Errorf("Node expected to have value %v but got %v", graph.edges[i].val, newGraph.edges[i].val)
		}
	}
}
