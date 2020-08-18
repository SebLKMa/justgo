package socialnetwork

import (
	"errors"
	"fmt"

	"../datastructures/graph"
)

var g graph.ItemGraph

func getGraphAsString() string {
	return g.String()
}

// AddFriend adds new friend
func AddFriend(email string) (*graph.Node, error) {
	newFriend := graph.Node{Value: email}

	_, found := g.FindNode(&newFriend)
	if found {
		return &graph.Node{}, errors.New("friend already exists")
	}

	g.AddNode(&newFriend)

	return &newFriend, nil
}

// AddFriendship adds new friendship between 2 friends
func AddFriendship(email1 string, email2 string) error {
	friend1 := graph.Node{Value: email1}
	found1, found := g.FindNode(&friend1)
	if !found {
		return fmt.Errorf("Friend does not exist - %s", friend1.String())
	}
	friend2 := graph.Node{Value: email2}
	found2, found := g.FindNode(&friend2)
	if !found {
		return fmt.Errorf("Friend does not exist - %s", friend2.String())
	}

	g.AddEdge(found1, found2)

	return nil
}

// GetFriends retrieves the friends of the specified friend
func GetFriends(email string) (friends []string) {
	person := graph.Node{Value: email}
	adjacents := g.GetAdjacents(&person)
	for _, node := range adjacents {
		friends = append(friends, node.String())
	}
	return
}

// GetCommonFriends retrieves the common friends between 2 friends
func GetCommonFriends(email1 string, email2 string) (friends []string) {
	friend1 := graph.Node{Value: email1}
	friend1Adjacents := g.GetAdjacents(&friend1)

	friend2 := graph.Node{Value: email2}
	friend2Adjacents := g.GetAdjacents(&friend2)

	commonFriends := g.Intersection(friend1Adjacents, friend2Adjacents)
	for _, node := range commonFriends {
		friends = append(friends, node.String())
	}
	return
}

// AddSubscription adds request for subcribing to changes of a target
func AddSubscription(requestorEmail string, targetEmail string) error {
	requestor := graph.Node{Value: requestorEmail}
	requestorFound, found := g.FindNode(&requestor)
	if !found {
		return fmt.Errorf("Friend does not exist - %s", requestor.String())
	}
	target := graph.Node{Value: targetEmail}
	targetFound, found := g.FindNode(&target)
	if !found {
		return fmt.Errorf("Friend does not exist - %s", target.String())
	}

	g.AddSubscription(requestorFound, targetFound)

	return nil
}

// BlockTarget adds request for blocking a target from getting any changes to a target
func BlockTarget(requestorEmail string, targetEmail string) error {
	requestor := graph.Node{Value: requestorEmail}
	requestorFound, found := g.FindNode(&requestor)
	if !found {
		return fmt.Errorf("Friend does not exist - %s", requestor.String())
	}
	target := graph.Node{Value: targetEmail}
	targetFound, found := g.FindNode(&target)
	if !found {
		return fmt.Errorf("Friend does not exist - %s", target.String())
	}

	g.AddBlockedList(requestorFound, targetFound)

	return nil
}

// GetActualSubscriptions returns the actual subscriptions of a specified friend.
// The returned recipients must not contain those that have been blocked by the specified friend.
func GetActualSubscriptions(email string) (recipients []string) {
	person := graph.Node{Value: email}

	originals := g.GetSubscriptions(&person)
	//var actuals []*Node
	isBlocked := false
	for _, subscriber := range originals {
		blockedlist := g.GetBlockedList(subscriber)
		for _, blocked := range blockedlist {
			if blocked.Value == person.Value {
				isBlocked = true
				break
			}
		}
		if !isBlocked {
			//actuals = append(actuals, subscriber)
			recipients = append(recipients, subscriber.String())
		}
		isBlocked = false
	}
	return
}

// ClearNetwork clear all objects in the social network graph
func ClearNetwork() {
	g.ClearAll()
}
