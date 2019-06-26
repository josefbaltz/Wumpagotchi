package main

import (
	"github.com/bwmarrin/discordgo"
)

// leftHandler decides which story the user will get depending on the Wumpus' age
// Also decides whether or not the user will be able to claim an egg or not
func leftHandler(UserWumpus Wumpus, event *discordgo.MessageCreate, session *discordgo.Session) {
	if UserWumpus.Age > 9 {
		go sendMessage(session, event, event.ChannelID, UserWumpus.Name+" has something important to tell you, They were accepted into Discordiversity and will be studying Wumpology, basically how to be a True Discord Wumpus, With a full ride scholarship! The wumpus shares how they loved all of the time they spent with you.")
		return
	}
	if UserWumpus.Age > 4 && UserWumpus.Age < 10 {
		go sendMessage(session, event, event.ChannelID, UserWumpus.Name+" wants to talk with you. They enjoyed the time that they spent with you, but want to pursue greener pastures. They’ll be packing their things and leaving soon in search of one.")
		return
	}
	go sendMessage(session, event, event.ChannelID, "You can’t seem to find the wumpus anywhere. You don’t worry too much and head for the fridge for a quick snack. As you head towards the fridge you see a note addressed to "+event.Author.Username+". You open and read, Hey "+event.Author.Username+", I’m sorry but i’ve decided to leave without telling you first, all I wanted was a friend but I’m constantly stressed out living with you.")
	return
}

func claimHandler(UserWumpus Wumpus, event *discordgo.MessageCreate, session *discordgo.Session) {

}
