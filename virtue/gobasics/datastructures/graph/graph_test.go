package graph

import (
	"fmt"
	"testing"
)

var g ItemGraph

func populateGraph() {
	nA := Node{"A"}
	nB := Node{"B"}
	nC := Node{"C"}
	nD := Node{"D"}
	nE := Node{"E"}
	nF := Node{"F"}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

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
	checkNode := Node{"F"}
	foundNode, found := g.FindNode(&checkNode)
	if found {
		fmt.Printf("Node found %v\n", foundNode)
	} else {
		fmt.Printf("Node not found %v\n", checkNode)
	}

	checkNode = Node{"53b08cd2-db9d-11ea-87d0-0242ac130003"}
	foundNode, found = g.FindNode(&checkNode)
	if found {
		fmt.Printf("Node found %v\n", foundNode)
	} else {
		fmt.Printf("Node not found %v\n", checkNode)
	}
}

// TestGetAdjacents test get adjacent vertices(nodes) of a vertex
func TestGetAdjacents(t *testing.T) {
	nA := Node{"A"}
	fmt.Println("TestGetAdjacents", nA.String())
	adjacents := g.GetAdjacents(&nA)
	for i, node := range adjacents {
		fmt.Println(i, node.String())
	}
}

// TestGetSubscriptions test get the subscribers of a node
func TestGetSubscriptions(t *testing.T) {
	nA := Node{"A"}
	fmt.Println("TestGetSubscriptions", nA.String())
	results := g.GetSubscriptions(&nA)
	for i, node := range results {
		fmt.Println(i, node.String())
	}
}

// TestGetBlockedList test get the blocked list of a node
func TestGetBlockedList(t *testing.T) {
	nC := Node{"C"}
	fmt.Println("TestGetBlockedList", nC.String())
	results := g.GetBlockedList(&nC)
	for i, node := range results {
		fmt.Println(i, node.String())
	}
}

// TestGetActualSubscriptions get the actual subscriptions of a node (subscribed nodes minus its blocked nodes)
func TestGetActualSubscriptions(t *testing.T) {
	nA := Node{"A"}
	fmt.Println("TestGetActualSubscriptions", nA.String())
	originals := g.GetSubscriptions(&nA)
	var actuals []*Node
	isBlocked := false
	for _, subscriber := range originals {
		blockedlist := g.GetBlockedList(subscriber)
		for _, blocked := range blockedlist {
			if blocked.Value == nA.Value {
				isBlocked = true
				break
			}
		}
		if !isBlocked {
			actuals = append(actuals, subscriber)
		}
		isBlocked = false
	}

	for i, node := range actuals {
		fmt.Println(i, node.String())
	}
}

func TestClearAll(t *testing.T) {
	nodeCount := len(g.nodes)
	edgeCount := len(g.edges)
	subscriptionCount := len(g.subscriptions)
	blockedCount := len(g.blockedlist)
	fmt.Println("Before clear all")
	fmt.Println("node count: ", nodeCount)
	fmt.Println("edge count: ", edgeCount)
	fmt.Println("subs count: ", subscriptionCount)
	fmt.Println("blocked count: ", blockedCount)
	g.subscriptions = make(map[Node][]*Node)

	g.ClearAll()

	nodeCount = len(g.nodes)
	edgeCount = len(g.edges)
	subscriptionCount = len(g.subscriptions)
	blockedCount = len(g.blockedlist)
	fmt.Println("After clear all")
	fmt.Println("node count: ", nodeCount)
	fmt.Println("edge count: ", edgeCount)
	fmt.Println("subs count: ", subscriptionCount)
	fmt.Println("blocked count: ", blockedCount)
}
