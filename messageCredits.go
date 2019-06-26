package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

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
