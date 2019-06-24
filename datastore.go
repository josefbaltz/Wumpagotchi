package main

import (
	"fmt"

	"cloud.google.com/go/datastore"
)

//AddWumpus - Adds a new wumpus to the database
func AddWumpus(UserID string, NewWumpus Wumpus) (err error) {
	userKey := datastore.NameKey("User", UserID, nil)
	if _, err := gcp.Put(ctx, userKey, &NewWumpus); err != nil {
		fmt.Println("==Warning==\nFailed to add Wumpus to Datastore")
		return err
	}
	return
}
