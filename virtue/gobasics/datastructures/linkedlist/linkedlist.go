package main

import (
	"fmt"
)

// Node definition for a linked list
type Node struct {
	Value   int
	NextPtr *Node
}

var root = new(Node) // cannot assign to untyped nil, has to be initialized as empty struct ptr
var count = 0

// Initialize list to nil
func Initialize() {
	root = nil
	count = 0
}

// AddNode adds the new value to the end of the list
func AddNode(newValue int) {
	addNode(root, newValue)
}

// addNode adds the new value to the end of the list
func addNode(nodePtr *Node, newValue int) int {
	newNodePtr := &Node{newValue, nil} // new node will always to appended to end, its next ptr always nil

	// list is empty
	if root == nil {
		nodePtr = newNodePtr
		root = nodePtr // set the root ptr
		count++
		return 0
	}

	// duplicate value
	if newValue == nodePtr.Value {
		fmt.Println("Node value already exists:", newValue)
		return -1
	}

	// we only append to end
	if nodePtr.NextPtr == nil {
		nodePtr.NextPtr = newNodePtr
		count++
		return 1
	}

	// otherwise, recursively forward to the next node until end of list
	return addNode(nodePtr.NextPtr, newValue)
}

// Traverse and print the values of the list from beginning to end
func Traverse() {
	traverse(root)
}

// traverse and print the values of the list from beginning to end
func traverse(nodePtr *Node) {
	if nodePtr == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for nodePtr != nil {
		fmt.Printf("%d -> ", nodePtr.Value)
		nodePtr = nodePtr.NextPtr // move to next node
	}
	fmt.Println()
}

// Lookup checks if the value exists in the list
func Lookup(value int) bool {
	return lookup(root, value)
}

// lookup checks if the value exists in the list
func lookup(nodePtr *Node, value int) bool {
	// list is empty
	if root == nil {
		//nodePtr = &Node{value, nil}
		//root = nodePtr
		return false
	}

	// found
	if value == nodePtr.Value {
		return true
	}

	// reached end of list, not found
	if nodePtr.NextPtr == nil {
		return false
	}

	// move on to next node
	return lookup(nodePtr.NextPtr, value)
}

// Size returns the number of nodes in the list
func Size() int {
	return count
}

func main() {
	fmt.Println(root)
	//root = nil
	Initialize()
	Traverse()
	fmt.Printf("Size %d\n", Size()) // 0

	AddNode(1)
	AddNode(-1)
	Traverse()
	fmt.Printf("Size %d\n", Size()) // 2

	AddNode(10)
	AddNode(5)
	AddNode(45)
	AddNode(5)
	AddNode(5)
	Traverse()
	fmt.Printf("Size %d\n", Size()) // 2 + 3

	AddNode(42)
	Traverse()
	fmt.Printf("Size %d\n", Size()) // 2 + 3 + 1

	value := 42
	if Lookup(value) {
		fmt.Printf("%d Node exists\n", value)
	} else {
		fmt.Printf("%d Node does not exists\n", value)
	}

	value = 62
	if Lookup(value) {
		fmt.Printf("%d Node exists\n", value)
	} else {
		fmt.Printf("%d Node does not exists\n", value)
	}
}
