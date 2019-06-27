package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// messageCredits is a function that checks if a user is eligible to recieve credits for their typed messages
// First this checks to see if the User is a bot and the message does not start with the command prefix if either is true we return
// We check to see if the user has a wumpus, if an error is returned we will return
// If all passes we Update the User's Credits and then push to the Datastore
func messageCredits(session *discordgo.Session, event *discordgo.MessageCreate) {
	if strings.HasPrefix(event.Content, CommandPrefix) || event.Author.Bot {
		return
	}
	UserWumpus, err := GetWumpus(event.Author.ID, true)
	if err != nil {
		return
	}
	UserWumpus.Credits++
	UpdateWumpus(event.Author.ID, UserWumpus)
	return
}
