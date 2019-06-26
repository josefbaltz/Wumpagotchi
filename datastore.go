package main

import (
	"fmt"

	"cloud.google.com/go/datastore"
)

//UpdateWumpus - Adds a new wumpus to the database
func UpdateWumpus(UserID string, NewWumpus Wumpus) (err error) {
	userKey := datastore.NameKey("User", UserID, nil)
	if _, err := gcp.Put(ctx, userKey, &NewWumpus); err != nil {
		fmt.Println("==Warning==\nFailed to update Wumpus in Datastore")
		return err
	}
	return
}

//GetWumpus - Retrieves a wumpus from the database
func GetWumpus(UserID string, ignoreWarning bool) (UserWumpus Wumpus, err error) {
	userKey := datastore.NameKey("User", UserID, nil)
	if err := gcp.Get(ctx, userKey, &UserWumpus); err != nil {
		if ignoreWarning == true {
			return UserWumpus, err
		}
		fmt.Println("==Warning==\nFailed to retrieve Wumpus from Datastore")
		return UserWumpus, err
	}
	return UserWumpus, nil
}

//DeleteWumpus - Deletes a wumpus from the database, NOT TO BE CONFUSED WITH MARKING ONE AS LEFT
func DeleteWumpus(UserID string) (err error) {
	userKey := datastore.NameKey("User", UserID, nil)
	if err := gcp.Delete(ctx, userKey); err != nil {
		fmt.Println("==Warning==\nFailed to delete Wumpus form Datastore")
		return err
	}
	return
}
