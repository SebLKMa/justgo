package main

import (
	"container/ring"
	"fmt"
)

var myRing = ring.New(11)

func initRing() {
	fmt.Println(myRing.Len())
	fmt.Println("Empty ring: ", *myRing)

	// populate myRing, using ring Len() as a stop condition for ring data structure
	for i := 0; i < myRing.Len()-1; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}

	myRing.Value = 2 // value already exists
}

// demoRing demonstrates ring iteration
func demoRing() {
	sum := 0
	// calls this function for each element in the ring
	// using ring Do it knows when to stop the ring iteration
	myRing.Do(func(x interface{}) {
		t := x.(int)
		fmt.Print(t, " ")
		sum = sum + t
	})
	fmt.Println("Sum: ", sum)

	for i := 0; i < myRing.Len(); i++ {
		fmt.Print(myRing.Value, " ")
		myRing = myRing.Next()
		//fmt.Print(myRing.Value, " ")
	}
	fmt.Println()
}

// demoRing demonstrates ring over iterate its length
func demoRingOver() {
	sum := 0
	// calls this function for each element in the ring
	// using ring Do it knows when to stop the ring iteration
	myRing.Do(func(x interface{}) {
		t := x.(int)
		fmt.Print(t, " ")
		sum = sum + t
	})
	fmt.Println("Sum: ", sum)

	for i := 0; i < myRing.Len()+2; i++ {
		myRing = myRing.Next()
		fmt.Print(myRing.Value, " ")
	}
	fmt.Println()
}

func main() {
	initRing()
	fmt.Println(myRing.Len())
	demoRing()
	fmt.Println(myRing.Len())
	demoRingOver()
}
