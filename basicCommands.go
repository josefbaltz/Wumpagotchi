package main

import (
	"strings"
	"time"

	"github.com/Necroforger/dgwidgets"

	"github.com/bwmarrin/discordgo"
)

func basicCommands(session *discordgo.Session, event *discordgo.MessageCreate) {
	messageContent := strings.Split(strings.ToLower(event.Content), " ")
	if messageContent[0] == CommandPrefix+"help" {
		HelpEmbed := &discordgo.MessageEmbed{
			Color:       0x669966, //Wumpus Leaf Green
			Title:       "Wumpagotchi Help",
			Description: "We're here to help!",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   CommandPrefix + "help",
					Value:  "Displays this text",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   CommandPrefix + "adopt",
					Value:  "Adopt a Wumpus!",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   CommandPrefix + "store",
					Value:  "Visit the store",
					Inline: true,
				},
			},
		}
		session.ChannelMessageSendEmbed(event.ChannelID, HelpEmbed)
		return
	}
	if messageContent[0] == CommandPrefix+"store" {
		store := dgwidgets.NewPaginator(session, event.Message.ChannelID)
		store.Add(&discordgo.MessageEmbed{
			Color:       0x669966, //Wumpus Leaf Green
			Title:       "Store",
			Description: "Buy trinkets and things for your Wumpus!",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Floop",
					Value:  "A standard diet for any Wumpus!",
					Inline: false,
				},
			},
			Image: &discordgo.MessageEmbedImage{
				URL: "https://images-na.ssl-images-amazon.com/images/I/81xQBb5jRzL._SY355_.jpg",
			},
		},
			&discordgo.MessageEmbed{
				Color:       0x669966, //Wumpus Leaf Green
				Title:       "Store",
				Description: "Buy trinkets and things for your Wumpus!",
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   "Gummy Gem",
						Value:  "Some hardened! Has some warnings on the side, probably not important.",
						Inline: false,
					},
				},
				Image: &discordgo.MessageEmbedImage{
					URL: "https://images-na.ssl-images-amazon.com/images/I/81XtDkc7vKL._SL1500_.jpg",
				},
			},
			&discordgo.MessageEmbed{
				Color:       0x669966, //Wumpus Leaf Green
				Title:       "Store",
				Description: "Buy trinkets and things for your Wumpus!",
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   "Salad",
						Value:  "A tasty nutritious salad! Wumpi are known to love these!",
						Inline: false,
					},
				},
				Image: &discordgo.MessageEmbedImage{
					URL: "https://media1.s-nbcnews.com/i/newscms/2018_42/1378147/sandra-lee-food-today-main-181018-02_28c1f1d7033c651ae8bd93a89f929201.jpg",
				},
			})
		store.SetPageFooters()
		store.ColourWhenDone = 0x36393f
		store.Widget.Timeout = time.Minute * 1
		store.Spawn()
		return
	}
}
