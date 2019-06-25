package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func basicCommands(session *discordgo.Session, event *discordgo.MessageCreate) {
	messageContent := strings.Split(strings.ToLower(event.Content), " ")
	if messageContent[0] == CommandPrefix+"help" && !event.Author.Bot {
		HelpEmbed := &discordgo.MessageEmbed{
			Color:       0x669966, //Wumpus Leaf Green
			Title:       "Wumpagotchi Help",
			Description: "We're here to help!",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "What is Wumpagotchi?",
					Value:  "Wumpagotchi is Tamagotchi but with Wumpus!\nStart by adopting a wumpus with w.adopt, type messages to gain credit, and play games with w.play, you can buy things to take care of your Wumpus with w.store",
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   CommandPrefix + "help",
					Value:  "Displays this text",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   CommandPrefix + "adopt <Wumpus Name>",
					Value:  "Adopt a Wumpus!",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   CommandPrefix + "store",
					Value:  "Visit the store",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   CommandPrefix + "view",
					Value:  "View how your Wumpus is doing",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   CommandPrefix + "play",
					Value:  "Go mining with your wumpus",
					Inline: true,
				},
			},
		}
		sendEmbed(session, event, event.ChannelID, HelpEmbed)
		return
	}
	if messageContent[0] == CommandPrefix+"store" && !event.Author.Bot {
		return
	}
}
