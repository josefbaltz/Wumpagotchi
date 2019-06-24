package main

import (
	"strconv"
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
			UpdateWumpus(event.Author.ID, NewWumpus)
			session.ChannelMessageSend(event.ChannelID, "Congrats, you have adopted "+NewWumpus.Name+" as your Wumpus!")
		} else {
			session.ChannelMessageSend(event.ChannelID, "Your Wumpus needs a name to be adopted!")
		}
	}
	if messageContent[0] == CommandPrefix+"view" {
		UserWumpus, err := GetWumpus(event.Author.ID)
		if err != nil {
			session.ChannelMessageSend(event.ChannelID, "Something went wrong, please contact the devs!")
		}
		ViewEmbed := &discordgo.MessageEmbed{
			Color: 0x669966, //Wumpus Leaf Green
			Title: UserWumpus.Name,
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Age",
					Value:  strconv.Itoa(UserWumpus.Age),
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "Health",
					Value:  strconv.Itoa(UserWumpus.Health),
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "Hunger",
					Value:  strconv.Itoa(UserWumpus.Hunger),
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "Energy",
					Value:  strconv.Itoa(UserWumpus.Energy),
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "Happiness",
					Value:  strconv.Itoa(UserWumpus.Happiness),
					Inline: false,
				},
			},
			Image: &discordgo.MessageEmbedImage{
				URL: "https://images-na.ssl-images-amazon.com/images/I/81xQBb5jRzL._SY355_.jpg",
			},
		}
		session.ChannelMessageSendEmbed(event.ChannelID, ViewEmbed)
	}
}
