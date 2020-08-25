# Initialization
curl --header "Content-Type: application/json" --request POST http://localhost:8282/clearnetwork

echo User story 1a.
echo As a user, I need an API to create a friend connection between two email addresses.
curl --header "Content-Type: application/json" --request POST --data '{"Email": "seb@bla.com"}' http://localhost:8282/addfriend
curl --header "Content-Type: application/json" --request POST --data '{"Email": "wendy@bla.com"}' http://localhost:8282/addfriend
curl --header "Content-Type: application/json" --request POST --data '{"Email": "can@bla.com"}' http://localhost:8282/addfriend
curl --header "Content-Type: application/json" --request POST --data '{"Email": "gillian@bla.com"}' http://localhost:8282/addfriend
curl --header "Content-Type: application/json" --request POST --data '{"Email": "dennis@bla.com"}' http://localhost:8282/addfriend
curl --header "Content-Type: application/json" --request POST --data '{"Email": "sherry@bla.com"}' http://localhost:8282/addfriend

echo User story 1b.
echo As a user, I need an API to create a friend connection between two email addresses.
curl --header "Content-Type: application/json" --request POST --data '{"friends": ["seb@bla.com", "wendy@bla.com"]}' http://localhost:8282/addfriendship
curl --header "Content-Type: application/json" --request POST --data '{"friends": ["seb@bla.com", "can@bla.com"]}' http://localhost:8282/addfriendship
curl --header "Content-Type: application/json" --request POST --data '{"friends": ["seb@bla.com", "gillian@bla.com"]}' http://localhost:8282/addfriendship
curl --header "Content-Type: application/json" --request POST --data '{"friends": ["gillian@bla.com", "wendy@bla.com"]}' http://localhost:8282/addfriendship
curl --header "Content-Type: application/json" --request POST --data '{"friends": ["gillian@bla.com", "can@bla.com"]}' http://localhost:8282/addfriendship
curl --header "Content-Type: application/json" --request POST --data '{"friends": ["gillian@bla.com", "sherry@bla.com"]}' http://localhost:8282/addfriendship

echo User story 2.
echo As a user, I need an API to retrieve the friends list for an email address.
curl --header "Content-Type: application/json" --request GET --data '{"Email": "seb@bla.com"}' http://localhost:8282/getfriends

echo User story 3.
echo As a user, I need an API to retrieve the common friends list between two email addresses.
curl --header "Content-Type: application/json" --request GET --data '{"friends": ["seb@bla.com", "gillian@bla.com"]}' http://localhost:8282/getcommonfriends
curl --header "Content-Type: application/json" --request GET --data '{"friends": ["seb@bla.com", "sherry@bla.com"]}' http://localhost:8282/getcommonfriends

echo User story 4.
echo As a user, I need an API to subscribe to updates from an email address.
curl --header "Content-Type: application/json" --request POST --data '{"requestor": "seb@bla.com", "target": "wendy@bla.com"}' http://localhost:8282/addsubscription
curl --header "Content-Type: application/json" --request POST --data '{"requestor": "seb@bla.com", "target": "can@bla.com"}' http://localhost:8282/addsubscription
curl --header "Content-Type: application/json" --request POST --data '{"requestor": "seb@bla.com", "target": "gillian@bla.com"}' http://localhost:8282/addsubscription
curl --header "Content-Type: application/json" --request POST --data '{"requestor": "gillian@bla.com", "target": "seb@bla.com"}' http://localhost:8282/addsubscription
curl --header "Content-Type: application/json" --request POST --data '{"requestor": "gillian@bla.com", "target": "sherry@bla.com"}' http://localhost:8282/addsubscription

echo User story 5.
echo As a user, I need an API to block updates from an email address.
curl --header "Content-Type: application/json" --request POST --data '{"requestor": "gillian@bla.com", "target": "seb@bla.com"}' http://localhost:8282/blocktarget
# {"success":true,"error":""}
curl --header "Content-Type: application/json" --request POST --data '{"requestor": "wendy@bla.com", "target": "sherry@bla.com"}' http://localhost:8282/blocktarget
# {"success":true,"error":""}

echo User story 6.
echo As a user, I need an API to retrieve all email addresses that can receive updates from an email address.
echo Note that I am demonstrating the actual subscriptions are those whom the specified friend has subscribed to 
echo and the specified friend has not been blocked by them.
curl --header "Content-Type: application/json" --request GET --data '{"email": "seb@bla.com"}' http://localhost:8282/getactualsubscriptions
# {"success":true,"recipients":["wendy@bla.com","can@bla.com"]#}


