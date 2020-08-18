package socialnetwork

import (
	"fmt"
	"testing"

	"../datastructures/graph"
)

var tg graph.ItemGraph

func TestClearNetwork(t *testing.T) {
	fmt.Println("TestClearNetwork")
	ClearNetwork()
}

func TestAddFriends(t *testing.T) {
	fmt.Println("TestAddFriends")
	seb, err := AddFriend("seb@example.com")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(seb.String() + " added")
	}

	wendy, err := AddFriend("wendy@example.com")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(wendy.String() + " added")
	}

	can, err := AddFriend("can@example.com")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(can.String() + " added")
	}

	gillian, err := AddFriend("gillian@example.com")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(gillian.String() + " added")
	}

	sherry, err := AddFriend("sherry@example.com")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(sherry.String() + " added")
	}
}

func TestCheckExistError(t *testing.T) {
	fmt.Println("TestCheckExistError")
	nobody, err := AddFriend("nobody@example.com")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(nobody.String() + " added")
	}

	nobody, err = AddFriend("nobody@example.com")
	if err == nil {
		t.Errorf("expected error not nil, actual error is nil")
	} else {
		fmt.Println("error expected ok - " + err.Error())
	}
}

func TestFriendship(t *testing.T) {
	fmt.Println("TestFriendship")
	err := AddFriendship("seb@example.com", "wendy@example.com")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = AddFriendship("seb@example.com", "can@example.com")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = AddFriendship("seb@example.com", "gillian@example.com")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = AddFriendship("gillian@example.com", "wendy@example.com")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = AddFriendship("gillian@example.com", "can@example.com")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = AddFriendship("gillian@example.com", "sherry@example.com")
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(getGraphAsString())
}

func TestGetFriends(t *testing.T) {
	fmt.Println("TestGetFriends")
	friends := GetFriends("seb@example.com")
	for _, friend := range friends {
		fmt.Println(friend)
	}
}

func TestGetCommonFriends(t *testing.T) {
	fmt.Printf("TestGetCommonFriends %s and %s\n", "seb@example.com", "gillian@example.com")
	friends := GetCommonFriends("seb@example.com", "gillian@example.com")
	for _, friend := range friends {
		fmt.Println(friend)
	}

	fmt.Printf("TestGetCommonFriends %s and %s\n", "seb@example.com", "sherry@example.com")
	friends = GetCommonFriends("seb@example.com", "sherry@example.com")
	for _, friend := range friends {
		fmt.Println(friend)
	}
}

func TestAddSubscription(t *testing.T) {
	fmt.Println("TestAddSubscription")
	err := AddSubscription("seb@example.com", "wendy@example.com")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = AddSubscription("seb@example.com", "can@example.com")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = AddSubscription("seb@example.com", "gillian@example.com")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = AddSubscription("gillian@example.com", "seb@example.com")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = AddSubscription("gillian@example.com", "sherry@example.com")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestBlockTarget(t *testing.T) {
	fmt.Println("TestBlockTarget")
	err := BlockTarget("gillian@example.com", "seb@example.com")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = BlockTarget("wendy@example.com", "sherry@example.com")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetActualSubscriptions(t *testing.T) {
	fmt.Println("TestGetActualSubscriptions of seb@example.com")
	actuals := GetActualSubscriptions("seb@example.com")
	for _, subscriber := range actuals {
		fmt.Println(subscriber)
	}
}

/*
func populateGraph() {
	seb := graph.Node{Value: "seb@example.com"}
	wendy := graph.Node{Value: "wendy@example.com"}
	can := graph.Node{Value: "can@example.com"}
	gillian := graph.Node{Value: "gillian@example.com"}
	sherry := graph.Node{Value: "sherry@example.com"}
	nobody := graph.Node{Value: "nobody@example.com"}
	g.AddNode(&seb)
	g.AddNode(&wendy)
	g.AddNode(&can)
	g.AddNode(&gillian)
	g.AddNode(&sherry)
	g.AddNode(&nobody)

	g.AddEdge(&seb, &wendy)
	g.AddEdge(&seb, &can)
	g.AddEdge(&seb, &gillian)

	g.AddEdge(&gillian, &wendy)
	g.AddEdge(&gillian, &can)
	g.AddEdge(&gillian, &sherry)

}

func main() {
	populateGraph()
	g.String()

	seb := graph.Node{Value: "seb@example.com"}
	fmt.Println("Friends of", seb.String())
	sebAdjacents := g.GetAdjacents(&seb)
	for i, node := range sebAdjacents {
		fmt.Println(i, node.String())
	}

	gillian := graph.Node{Value: "gillian@example.com"}
	fmt.Println("Friends of", gillian.String())
	gillianAdjacents := g.GetAdjacents(&gillian)
	for i, node := range gillianAdjacents {
		fmt.Println(i, node.String())
	}

	sherry := graph.Node{Value: "sherry@example.com"}
	fmt.Println("Friends of", sherry.String())
	sherryAdjacents := g.GetAdjacents(&sherry)
	for i, node := range sherryAdjacents {
		fmt.Println(i, node.String())
	}

	nobody := graph.Node{Value: "nobody@example.com"}
	fmt.Println("Friends of", nobody.String())
	nobodyAdjacents := g.GetAdjacents(&nobody)
	for i, node := range nobodyAdjacents {
		fmt.Println(i, node.String())
	}

	commonFriends := g.Intersection(sebAdjacents, gillianAdjacents)
	fmt.Println("Common friends:", seb.String(), gillian.String())
	for i, node := range commonFriends {
		fmt.Println(i, node.String())
	}

	commonFriends = g.Intersection(sebAdjacents, sherryAdjacents)
	fmt.Println("Common friend:s", seb.String(), sherry.String())
	for i, node := range commonFriends {
		fmt.Println(i, node.String())
	}

	commonFriends = g.Intersection(sebAdjacents, nobodyAdjacents)
	fmt.Println("Common friends:", seb.String(), nobody.String())
	for i, node := range commonFriends {
		fmt.Println(i, node.String())
	}
}
*/
