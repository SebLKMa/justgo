// doubly linked list as stack
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

// Enqueue adds the new value to the head
func Push(newValue int) {
	addNode(head, newValue)
}

// addNode adds the new value to the front of the list
func addNode(nodePtr *Node, newValue int) {
	newNodePtr := &Node{newValue, nil, nil}

	// list is empty
	if head == nil {
		head = newNodePtr
		tail = newNodePtr
		count++
		return
	}

	temp := head
	head = newNodePtr
	head.NextPtr = temp
	temp.PreviousPtr = newNodePtr
	count++
}

// Pop pops value to the front
func Pop() (bool, int) {
	return removeNode()
}

// removeNode pops value to the head
func removeNode() (bool, int) {
	// list is empty
	if head == nil {
		return false, 0
	}

	value := head.Value
	if count == 1 {
		head = nil
		tail = nil
	} else {
		head = head.NextPtr
		head.PreviousPtr = nil
	}

	count--

	return true, value
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

	ok, v := Pop()
	if ok {
		fmt.Println("Pop: ", v)
	} else {
		fmt.Println("Pop failed!")
	}

	Push(100)
	Traverse()
	Push(200)
	Traverse()
	//fmt.Printf("Size %d\n", Size()) // 0

	for i := 0; i < 10; i++ {
		Push(i)
	}

	for i := 0; i < 15; i++ {
		ok, v := Pop()
		if ok {
			fmt.Print(v, " ")
		} else {
			break // empty stack reached after pops
		}
	}
	fmt.Println()
	Traverse()
}
