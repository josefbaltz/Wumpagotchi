package main

import (
	"fmt"

	"cloud.google.com/go/datastore"
)

//UpdateWumpus - Adds a new wumpus to the database
func UpdateWumpus(UserID string, NewWumpus Wumpus) (err error) {
	userKey := datastore.NameKey("User", UserID, nil)
	if _, err := gcp.Put(ctx, userKey, &NewWumpus); err != nil {
		fmt.Println("==Warning==\nFailed to add Wumpus to Datastore")
		return err
	}
	return
}

//GetWumpus - Retrieves a wumpus from the database
func GetWumpus(UserID string) (UserWumpus Wumpus, err error) {
	userKey := datastore.NameKey("User", UserID, nil)
	if err := gcp.Get(ctx, userKey, &UserWumpus); err != nil {
		fmt.Println("==Warning==\nFailed to retrieve Wumpus from Datastore")
		return UserWumpus, err
	}
	return UserWumpus, nil
}
