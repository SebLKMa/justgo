package graphbfs

import (
	"fmt"
	"testing"

	"../graph"
)

var g graph.ItemGraph

func populateGraph() {
	nA := graph.Node{"A"}
	nB := graph.Node{"B"}
	nC := graph.Node{"C"}
	nD := graph.Node{"D"}
	nE := graph.Node{"E"}
	nF := graph.Node{"F"}
	nG := graph.Node{"G"}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)
	g.AddNode(&nG)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)

	g.AddSubscription(&nA, &nB)
	g.AddSubscription(&nA, &nC)
	g.AddSubscription(&nA, &nE)

	g.AddBlockedList(&nC, &nA)
	g.AddBlockedList(&nC, &nF)
}

// TestAddNodes test add nodes to graph
func TestAddNodes(t *testing.T) {
	fmt.Println("TestAddNodes")
	populateGraph()
	s := g.String()
	fmt.Println(s)
}

// TestFindNode test find node function
func TestFindNode(t *testing.T) {
	fmt.Println("TestFindNode")
	checkNode := graph.Node{"F"}
	foundNode, found := g.FindNode(&checkNode)
	if found {
		fmt.Printf("Node found %v\n", foundNode)
	} else {
		fmt.Printf("Node not found %v\n", checkNode)
	}

	checkNode = graph.Node{"53b08cd2-db9d-11ea-87d0-0242ac130003"}
	foundNode, found = g.FindNode(&checkNode)
	if found {
		fmt.Printf("Node found %v\n", foundNode)
	} else {
		fmt.Printf("Node not found %v\n", checkNode)
	}
}

// TestGetAdjacents test get adjacent vertices(nodes) of a vertex
func TestGetAdjacents(t *testing.T) {
	nA := graph.Node{"A"}
	fmt.Println("TestGetAdjacents", nA.String())
	adjacents := g.GetAdjacents(&nA)
	for i, node := range adjacents {
		fmt.Println(i, node.String())
	}
}

// TestTraverseBfs test breadth first search traversal of graph nodes
func TestTraverseBfs(t *testing.T) {
	fmt.Println("TestTraverseBfs")
	nA := graph.Node{"A"}
	Traverse(&g, &nA, func(n *graph.Node) {
		fmt.Printf("%v\n", n)
	})
}

// TestTraverseHopsBfs test breadth first search traversal of graph nodes
func TestTraverseHopsBfs(t *testing.T) {
	fmt.Println("TestTraverseHopsBfs")
	nA := graph.Node{"A"}
	parents, hops := TraverseHops(&g, &nA)
	for key, value := range parents {
		//parent := ""
		//if value != nil {
		//	parent = value.String()
		//}
		fmt.Printf("%s : %s\n", key, value)
	}
	fmt.Println("==============================")
	for key, value := range hops {
		fmt.Printf("%s : %d\n", key, value)
	}
}

// TestSearchBfs test breadth first search traversal of graph nodes
func TestSearchBfs(t *testing.T) {
	fmt.Println("TestSearchBfs")
	nSrc := graph.Node{"D"}
	nDst := graph.Node{"A"}
	if Search(&g, &nSrc, &nDst) {
		fmt.Printf("Found %v -> %v\n", nSrc, nDst)
	} else {
		fmt.Printf("Not Found %v -> %v\n", nSrc, nDst)
	}

	nDst = graph.Node{"F"}
	if Search(&g, &nSrc, &nDst) {
		fmt.Printf("Found %v -> %v\n", nSrc, nDst)
	} else {
		fmt.Printf("Not Found %v -> %v\n", nSrc, nDst)
	}

	nDst = graph.Node{"G"}
	if Search(&g, &nSrc, &nDst) {
		fmt.Printf("Found %v -> %v\n", nSrc, nDst)
	} else {
		fmt.Printf("Not Found %v -> %v\n", nSrc, nDst)
	}
}
