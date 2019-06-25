package main

import "github.com/bwmarrin/discordgo"

func messageCredits(session *discordgo.Session, event *discordgo.MessageCreate) {
	if UserWumpus, err := GetWumpus(event.Author.ID); err != nil {
	} else {
		UserWumpus.Credits++
		UpdateWumpus(event.Author.ID, UserWumpus)
	}
}
