package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	sn "../gobasics/socialnetwork"
)

// Email definition used in json data
type Email struct {
	Email string `json:"email"`
}

// Emails definition used in json data
type Emails struct {
	Emails []string `json:"friends"`
}

// RequestorTarget definition used in json data
type RequestorTarget struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

// Result definition used in json data
type Result struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// User story 1a.
// As a user, I need an API to create a friend connection between two email addresses.
func addFriendHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)

	w.Header().Set("Content-Type", "application/json")
	resp := Result{Success: true}
	var req Email

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errmsg := fmt.Sprintf("HTTP %d - %s", http.StatusBadRequest, err.Error())
		http.Error(w, errmsg, http.StatusBadRequest)
		return
	}

	log.Printf("Request data: %s\n", req.Email)
	friend, err := sn.AddFriend(req.Email)
	if err != nil {
		resp = Result{Success: false, Error: err.Error()}
	} else {
		log.Println(friend.String() + " added")
	}

	json.NewEncoder(w).Encode(resp)
}

// User story 1b.
// As a user, I need an API to create a friend connection between two email addresses.
func addFriendshipHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)

	w.Header().Set("Content-Type", "application/json")
	resp := Result{Success: true}
	var req Emails

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errmsg := fmt.Sprintf("HTTP %d - %s", http.StatusBadRequest, err.Error())
		http.Error(w, errmsg, http.StatusBadRequest)
		return
	}

	if len(req.Emails) != 2 {
		resp = Result{Success: false, Error: "There must be exactly 2 emails in this request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	for _, email := range req.Emails {
		log.Printf("Request data: %s\n", email)
	}

	err = sn.AddFriendship(req.Emails[0], req.Emails[1])
	if err != nil {
		resp = Result{Success: false, Error: err.Error()}
	}

	json.NewEncoder(w).Encode(resp)
}

// User story 2.
// As a user, I need an API to retrieve the friends list for an email address.
func getFriendsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)

	w.Header().Set("Content-Type", "application/json")
	var req Email

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errmsg := fmt.Sprintf("HTTP %d - %s", http.StatusBadRequest, err.Error())
		http.Error(w, errmsg, http.StatusBadRequest)
		return
	}

	log.Printf("Request data: %s\n", req.Email)

	type LocalResult struct {
		Success bool     `json:"success"`
		Friends []string `json:"friends"`
		Count   int      `json:"count"`
	}
	resultFriends := sn.GetFriends(req.Email)

	resp := LocalResult{}
	if resultFriends == nil {
		resp = LocalResult{Success: false, Friends: resultFriends, Count: 0}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp = LocalResult{Success: true, Friends: resultFriends, Count: len(resultFriends)}

	json.NewEncoder(w).Encode(resp)
}

// User story 3.
// As a user, I need an API to retrieve the common friends list between two email addresses.
func getCommonFriendsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)

	w.Header().Set("Content-Type", "application/json")
	var req Emails

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errmsg := fmt.Sprintf("HTTP %d - %s", http.StatusBadRequest, err.Error())
		http.Error(w, errmsg, http.StatusBadRequest)
		return
	}

	type LocalResult struct {
		Success bool     `json:"success"`
		Friends []string `json:"friends"`
		Count   int      `json:"count"`
		Error   string   `json:"error"`
	}
	resp := LocalResult{}
	if len(req.Emails) != 2 {
		resp = LocalResult{Success: false, Friends: nil, Count: 0, Error: "There must be exactly 2 emails in this request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	for _, email := range req.Emails {
		log.Printf("Request data: %s\n", email)
	}

	resultFriends := sn.GetCommonFriends(req.Emails[0], req.Emails[1])

	if resultFriends == nil {
		// still success though no friend
		resp = LocalResult{Success: true, Friends: resultFriends, Count: 0}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp = LocalResult{Success: true, Friends: resultFriends, Count: len(resultFriends)}

	json.NewEncoder(w).Encode(resp)
}

// User story 4.
// As a user, I need an API to subscribe to updates from an email address.
func addSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)

	w.Header().Set("Content-Type", "application/json")
	resp := Result{Success: true}
	var req RequestorTarget

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errmsg := fmt.Sprintf("HTTP %d - %s", http.StatusBadRequest, err.Error())
		http.Error(w, errmsg, http.StatusBadRequest)
		return
	}

	if req.Requestor == "" || req.Target == "" {
		resp = Result{Success: false, Error: "There must be exactly 2 emails in this request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	err = sn.AddSubscription(req.Requestor, req.Target)
	if err != nil {
		resp = Result{Success: false, Error: err.Error()}
	}

	json.NewEncoder(w).Encode(resp)
}

// User story 5.
// As a user, I need an API to block updates from an email address.
func blockTargetHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)

	w.Header().Set("Content-Type", "application/json")
	resp := Result{Success: true}
	var req RequestorTarget

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errmsg := fmt.Sprintf("HTTP %d - %s", http.StatusBadRequest, err.Error())
		http.Error(w, errmsg, http.StatusBadRequest)
		return
	}

	if req.Requestor == "" || req.Target == "" {
		resp = Result{Success: false, Error: "There must be exactly 2 emails in this request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	err = sn.BlockTarget(req.Requestor, req.Target)
	if err != nil {
		resp = Result{Success: false, Error: err.Error()}
	}

	json.NewEncoder(w).Encode(resp)
}

// User story 6.
// As a user, I need an API to retrieve all email addresses that can receive updates from an email address.
// Note that I am demonstrating the actual subscriptions are those whom the specified friend has subscribed to and the specified friend has not been blocked by them.
func getActualSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)

	w.Header().Set("Content-Type", "application/json")
	var req Email

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errmsg := fmt.Sprintf("HTTP %d - %s", http.StatusBadRequest, err.Error())
		http.Error(w, errmsg, http.StatusBadRequest)
		return
	}

	log.Printf("Request data: %s\n", req.Email)

	type LocalResult struct {
		Success    bool     `json:"success"`
		Recipients []string `json:"recipients"`
	}
	actuals := sn.GetActualSubscriptions(req.Email)

	resp := LocalResult{}
	if actuals == nil {
		resp = LocalResult{Success: false, Recipients: actuals}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp = LocalResult{Success: true, Recipients: actuals}

	json.NewEncoder(w).Encode(resp)
}

func clearNetworkHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)

	w.Header().Set("Content-Type", "application/json")
	resp := Result{Success: true}

	sn.ClearNetwork()

	json.NewEncoder(w).Encode(resp)
}

func main() {

	portPtr := flag.String("port", "8282", "port number")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/addfriend", addFriendHandler)
	mux.HandleFunc("/addfriendship", addFriendshipHandler)
	mux.HandleFunc("/getfriends", getFriendsHandler)
	mux.HandleFunc("/getcommonfriends", getCommonFriendsHandler)
	mux.HandleFunc("/addsubscription", addSubscriptionHandler)
	mux.HandleFunc("/blocktarget", blockTargetHandler)
	mux.HandleFunc("/getactualsubscriptions", getActualSubscriptionsHandler)
	mux.HandleFunc("/clearnetwork", clearNetworkHandler)

	port := *portPtr
	host := "0.0.0.0:" + port
	log.Println(host + " up and listening")

	err := http.ListenAndServe(host, mux)
	if err != nil {
		log.Fatal("Error creating server. ", err)
	}
}
