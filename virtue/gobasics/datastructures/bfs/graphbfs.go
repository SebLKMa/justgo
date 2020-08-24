package graphbfs

import (
	"../graph"
)

// nodeQueue the queue of Nodes, package internal, non-concurrency safe
type nodeQueue struct {
	items []graph.Node
}

// New creates a new NodeQueue instance
func (s *nodeQueue) New() *nodeQueue {
	s.items = []graph.Node{}
	return s
}

// Enqueue adds an Node to the end of the queue
func (s *nodeQueue) Enqueue(t graph.Node) {
	s.items = append(s.items, t)
}

// Dequeue removes an Node from the start of the queue
func (s *nodeQueue) Dequeue() *graph.Node {
	item := s.items[0]                // get the first item from slice
	s.items = s.items[1:len(s.items)] // like a substring begin-position=1, end-position=len(str)
	return &item
}

// Front returns the item next in the queue, without removing it
func (s *nodeQueue) Front() *graph.Node {
	item := s.items[0]
	return &item
}

// IsEmpty returns true if the queue is empty
func (s *nodeQueue) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of Nodes in the queue
func (s *nodeQueue) Size() int {
	return len(s.items)
}

// Traverse ... srcNode typically g.nodes[0]
func Traverse(g *graph.ItemGraph, srcNode *graph.Node, f func(*graph.Node)) {
	q := nodeQueue{items: []graph.Node{}}
	n := srcNode
	q.Enqueue(*n)
	visited := make(map[*graph.Node]bool)
	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()
		visited[node] = true
		near := g.GetAdjacents(node) //g.edges[*node]

		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				q.Enqueue(*j)
				visited[j] = true
			}
		}
		if f != nil {
			f(node)
		}
	}
}

// TraverseHops tracks the number of Hops from srcNode
func TraverseHops(g *graph.ItemGraph, srcNode *graph.Node) (parents map[string]string, hops map[string]int) {
	q := nodeQueue{items: []graph.Node{}}
	n := srcNode
	q.Enqueue(*n)
	//visited := make(map[*graph.Node]bool)
	//parents = make(map[*graph.Node]*graph.Node)
	visited := make(map[string]bool)
	parents = make(map[string]string)
	hops = make(map[string]int)
	parents[n.String()] = "" // root's has no parent
	hops[n.String()] = 0     // no hop to self
	for {
		if q.IsEmpty() {
			break
		}

		node := q.Dequeue()

		visited[node.String()] = true

		near := g.GetAdjacents(node) //g.edges[*node]

		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j.String()] {
				q.Enqueue(*j)
				visited[j.String()] = true
				parents[j.String()] = node.String()
				hops[j.String()] = hops[node.String()] + 1
			}
		}
	}
	return
}

// Search ...
func Search(g *graph.ItemGraph, srcNode *graph.Node, dstNode *graph.Node) bool {
	q := nodeQueue{items: []graph.Node{}}
	n := srcNode
	q.Enqueue(*n)
	visited := make(map[*graph.Node]bool)
	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()
		if node.String() == dstNode.String() {
			return true
		}
		visited[node] = true
		near := g.GetAdjacents(node) //g.edges[*node]

		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				q.Enqueue(*j)
				visited[j] = true
			}
		}
	}
	return false
}
