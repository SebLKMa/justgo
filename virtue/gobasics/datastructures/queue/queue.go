// doubly linked list as queue
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
func Enqueue(newValue int) {
	addNode(head, newValue)
}

// addNode adds the new value to the front of the list
func addNode(nodePtr *Node, newValue int) {
	newNodePtr := &Node{newValue, nil, nil}

	// list is empty
	if head == nil {
		//nodePtr = newNodePtr
		head = newNodePtr
		tail = newNodePtr
		count++
		return
	}

	oldhead := head                  // temp is old head
	head = newNodePtr                // new head
	head.NextPtr = oldhead           // new head's next ptr
	oldhead.PreviousPtr = newNodePtr // old head previous ptr
	count++
}

// Dequeue pops value from the tail
func Dequeue() (bool, int) {
	return removeNode()
}

// removeNode pops value from the tail
func removeNode() (bool, int) {
	// list is empty
	if tail == nil {
		return false, 0
	}

	value := tail.Value
	if count == 1 {
		head = nil
		tail = nil
	} else {
		tail = tail.PreviousPtr // new tail points to tail's previous
		tail.NextPtr = nil      // new tail next ptr set to nil
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
		fmt.Printf("-> %d", nodePtr.Value)
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

	Enqueue(10)
	fmt.Printf("Size %d\n", Size()) // 1
	Traverse()

	ok, v := Dequeue()
	if ok {
		fmt.Println("Dequeue: ", v)
	}
	Traverse()
	fmt.Printf("Size %d\n", Size()) // 0

	Enqueue(1)
	Enqueue(2)
	fmt.Printf("Size %d\n", Size()) // 2
	Traverse()

	ok, v = Dequeue()
	if ok {
		fmt.Println("Dequeue: ", v)
	}
	ok, v = Dequeue()
	if ok {
		fmt.Println("Dequeue: ", v)
	}
	Traverse()

	for i := 0; i < 5; i++ {
		Enqueue(i)
	}
	Traverse()
	fmt.Printf("Size %d\n", Size()) // 5

	ok, v = Dequeue()
	if ok {
		fmt.Println("Dequeue: ", v)
	}
	fmt.Printf("Size %d\n", Size()) // 4

	ok, v = Dequeue()
	if ok {
		fmt.Println("Dequeue: ", v)
	}
	fmt.Printf("Size %d\n", Size()) // 3

	Traverse()

}
