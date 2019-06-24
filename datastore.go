package main

import "cloud.google.com/go/datastore"

//AddWumpus - Adds a new wumpus to the database
func AddWumpus(UserID string) (err error) {
	userKey := datastore.NameKey("User", UserID, nil)
	NewWumpus := Wumpus{
		//TODO: Create Wumpus
	}
	return
}
