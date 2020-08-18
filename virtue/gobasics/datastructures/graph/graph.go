package graph

import (
	"fmt"
	"sync"
)

// AbstractType is the placeholder type for a concrete value.
type AbstractType interface{}

// Item the type of the binary search tree
type Item AbstractType

// Node a single node that composes the tree
type Node struct {
	Value Item
}

// String returns the node value as string
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

// ItemGraph the Items graph
type ItemGraph struct {
	nodes []*Node          // We only want to demonstrate a graph data structure here. For performance, one can use set (map[Node][]bool)
	edges map[Node][]*Node // Map of adjacents list

	// Map of publisher/subscriber this Key(publisher) make it will publish to Values(subscribers)
	subscriptions map[Node][]*Node
	// Map of Key(blocker) not interested in getting notification or friend by anyone from this Values(blocked)
	blockedlist map[Node][]*Node

	lock sync.RWMutex
}

// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

// FindNode return the node if it exists in graph
func (g *ItemGraph) FindNode(n *Node) (*Node, bool) {
	g.lock.RLock()
	defer g.lock.RUnlock()
	for i := range g.nodes {
		if g.nodes[i].Value == n.Value {
			// Found!
			return g.nodes[i], true
		}
	}
	return &Node{}, false
}

// AddEdge adds an edge to the graph
func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
	g.lock.Unlock()
}

// ClearAll edges and vertices of the underlying graph
func (g *ItemGraph) ClearAll() {
	g.lock.Lock()
	g.blockedlist = make(map[Node][]*Node)
	g.subscriptions = make(map[Node][]*Node)
	g.edges = make(map[Node][]*Node)
	g.nodes = nil
	g.lock.Unlock()
}

// GetAdjacents returns the adjacent nodes of a specified node
func (g *ItemGraph) GetAdjacents(node *Node) []*Node {
	g.lock.RLock()
	var adjacents []*Node

	for i := 0; i < len(g.nodes); i++ {
		if g.nodes[i].String() == node.String() {
			near := g.edges[*g.nodes[i]]
			for j := 0; j < len(near); j++ {
				adjacents = append(adjacents, near[j])
			}
			break
		}
	}

	g.lock.RUnlock()
	return adjacents
}

// String returns the edges in the graph as a string
func (g *ItemGraph) String() string {
	g.lock.RLock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	g.lock.RUnlock()
	return s
}

// AddSubscription to the graph
func (g *ItemGraph) AddSubscription(n1, n2 *Node) {
	g.lock.Lock()
	if g.subscriptions == nil {
		g.subscriptions = make(map[Node][]*Node)
	}
	g.subscriptions[*n1] = append(g.subscriptions[*n1], n2)
	g.lock.Unlock()
}

// GetSubscriptions returns the subscribers of a specified node
func (g *ItemGraph) GetSubscriptions(node *Node) []*Node {
	g.lock.RLock()
	subscribers := g.subscriptions[*node] // dereference to get the node value
	g.lock.RUnlock()
	return subscribers
}

// AddBlockedList to the graph
func (g *ItemGraph) AddBlockedList(n1, n2 *Node) {
	g.lock.Lock()
	if g.blockedlist == nil {
		g.blockedlist = make(map[Node][]*Node)
	}
	g.blockedlist[*n1] = append(g.blockedlist[*n1], n2)
	g.lock.Unlock()
}

// GetBlockedList returns the blocked list of a specified node
func (g *ItemGraph) GetBlockedList(node *Node) []*Node {
	g.lock.RLock()
	blocked := g.blockedlist[*node] // dereference to get the node value
	g.lock.RUnlock()
	return blocked
}

// Intersection returns the intersection between 2 sets of nodes.
func (g *ItemGraph) Intersection(a []*Node, b []*Node) (inter []*Node) {
	// interacting on the smallest list first can potentailly be faster...but not by much, worse case is the same
	low, high := a, b
	if len(a) > len(b) {
		low = b
		high = a
	}

	done := false
	for i, l := range low {
		for j, h := range high {
			// get future index values
			f1 := i + 1
			f2 := j + 1
			if l.String() == h.String() {
				inter = append(inter, h)
				if f1 < len(low) && f2 < len(high) {
					// if the future values aren't the same then that's the end of the intersection
					if low[f1] != high[f2] {
						done = true
					}
				}
				// we don't want to interate on the entire list everytime, so remove the parts we already looped on will make it faster each pass
				high = high[:j+copy(high[j:], high[j+1:])]
				break
			}
		}
		// nothing in the future so we are done
		if done {
			break
		}
	}
	return
}
