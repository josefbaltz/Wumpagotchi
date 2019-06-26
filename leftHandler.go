package main

import (
	"github.com/bwmarrin/discordgo"
)

func leftHandler(UserWumpus Wumpus, event *discordgo.MessageCreate, session *discordgo.Session) {
	if UserWumpus.Age >= 14 {
		sendMessage(session, event, event.ChannelID, "")
		return
	}
	if UserWumpus.Age > 9 {
		sendMessage(session, event, event.ChannelID, "")
		return
	}
	if UserWumpus.Age > 4 && UserWumpus.Age < 10 {
		sendMessage(session, event, event.ChannelID, "")
		return
	}
	sendMessage(session, event, event.ChannelID, "")
	return
}
