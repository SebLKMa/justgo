package main

import (
	"fmt"
)

// Node definition for a doubly linked list
type Node struct {
	Value       int
	PreviousPtr *Node
	NextPtr     *Node
}

var head = new(Node) // cannot use of untyped nil, initialized as empty struct ptr
var tail = new(Node) // cannot use of untyped nil, initialized as empty struct ptr
var count = 0

// Initialize list to nil
func Initialize() {
	head = nil
	tail = nil
	count = 0
}

// AddNode adds the new value to the end of the list
func AddNode(newValue int) {
	addNode(head, newValue)
}

// addNode adds the new value to the end of the list
func addNode(nodePtr *Node, newValue int) int {
	// new node
	newNodePtr := &Node{newValue, nil, nil}

	// list is empty
	if head == nil {
		nodePtr = newNodePtr
		head = nodePtr // head's previous is always nil
		tail = nodePtr
		count++
		return 0
	}

	// duplicate value
	if newValue == nodePtr.Value {
		fmt.Println("Node value already exists:", newValue)
		return -1
	}

	// if we are at end of list, append to end
	if nodePtr.NextPtr == nil {
		newNodePtr.PreviousPtr = nodePtr // just update the new node's previous
		nodePtr.NextPtr = newNodePtr
		tail = newNodePtr // remember to update tail's pointer
		count++
		return 1
	}

	// otherwise, recursively forward to the next node until end of list
	return addNode(nodePtr.NextPtr, newValue)
}

// Traverse and print the values of the list from beginning to end
func Traverse() {
	traverse(head)
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

// Reverse and print the values of the list from end to beginning
func Reverse() {
	reverse(tail)
}

// reverse and print the values of the list from end to beginning
func reverse(nodePtr *Node) {
	if nodePtr == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for nodePtr != nil {
		fmt.Printf("%d -> ", nodePtr.Value)
		nodePtr = nodePtr.PreviousPtr // move to previous node
	}
	fmt.Println()
}

// Lookup checks if the value exists in the list
func Lookup(value int) bool {
	return lookup(head, value)
}

// lookup checks if the value exists in the list
func lookup(nodePtr *Node, value int) bool {
	// list is empty
	if head == nil {
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
	fmt.Println(head)
	//root = nil
	Initialize()
	Traverse()
	fmt.Printf("Size %d\n", Size()) // 0

	AddNode(1)
	AddNode(1)
	Traverse()
	fmt.Printf("Size %d\n", Size()) // 1

	AddNode(10)
	AddNode(5)
	AddNode(45)
	AddNode(0)
	AddNode(0)
	Traverse()
	fmt.Printf("Size %d\n", Size()) // 1 + 4

	AddNode(42)
	Traverse()
	fmt.Printf("Size %d\n", Size()) // 5 + 1

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

	Reverse()
}
