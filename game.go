package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Wumpus :D
type Wumpus struct {
	Credits   int
	Name      string
	Color     int
	Age       int
	Health    int
	Hunger    int
	Energy    int
	Happiness int
	Sick      bool
}

func game(session *discordgo.Session, event *discordgo.MessageCreate) {
	messageContent := strings.Split(strings.ToLower(event.Content), " ")
	if messageContent[0] == CommandPrefix+"adopt" {
		if len(messageContent) > 1 {
			NewWumpus := Wumpus{
				Credits:   0,
				Name:      strings.TrimPrefix(event.Content, CommandPrefix+"adopt "),
				Color:     0,
				Age:       0,
				Health:    10,
				Hunger:    10,
				Energy:    10,
				Happiness: 10,
				Sick:      false,
			}
			AddWumpus(event.Author.ID, NewWumpus)
			session.ChannelMessageSend(event.ChannelID, "Congrats, you have adopted "+NewWumpus.Name+" as your Wumpus!")
		} else {
			session.ChannelMessageSend(event.ChannelID, "Your Wumpus needs a name to be adopted!")
		}
	}
}
