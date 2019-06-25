package main

import (
	"strings"
	"time"

	"github.com/Necroforger/dgwidgets"

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
					Name:	CommandPrefix + "play",
					Value:	"Go mining with your wumpus",
					Inline: true,
				}
			},
		}
		sendEmbed(session, event, event.ChannelID, HelpEmbed)
		return
	}
	if messageContent[0] == CommandPrefix+"store" && !event.Author.Bot {
		store := dgwidgets.NewPaginator(session, event.Message.ChannelID)
		store.Add(&discordgo.MessageEmbed{
			Color:       0x669966, //Wumpus Leaf Green
			Title:       "Store",
			Description: "Buy trinkets and things for your Wumpus!",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Floop (5Ꞡ)",
					Value:  "A standard diet for any Wumpus!",
					Inline: false,
				},
			},
			Image: &discordgo.MessageEmbedImage{
				URL: "https://orangeflare.me/imagehosting/Wumpagotchi/Floop.png",
			},
		},
			&discordgo.MessageEmbed{
				Color:       0x669966, //Wumpus Leaf Green
				Title:       "Store",
				Description: "Buy trinkets and things for your Wumpus!",
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   "Gummy Gem (10Ꞡ)",
						Value:  "A yummy snack! Has some warnings on the side, probably not important.",
						Inline: false,
					},
				},
				Image: &discordgo.MessageEmbedImage{
					URL: "https://orangeflare.me/imagehosting/Wumpagotchi/Gummy.png",
				},
			},
			&discordgo.MessageEmbed{
				Color:       0x669966, //Wumpus Leaf Green
				Title:       "Store",
				Description: "Buy trinkets and things for your Wumpus!",
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   "Medicine (15Ꞡ)",
						Value:  "A healthy boost for your Wumpus!",
						Inline: false,
					},
				},
				Image: &discordgo.MessageEmbedImage{
					URL: "https://orangeflare.me/imagehosting/Wumpagotchi/Salad.png",
				},
			},
			&discordgo.MessageEmbed{
				Color:       0x669966, //Wumpus Leaf Green
				Title:       "Store",
				Description: "Buy trinkets and things for your Wumpus!",
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   "Salad (30Ꞡ)",
						Value:  "A tasty nutritious salad! Wumpi are known to love these!",
						Inline: false,
					},
				},
				Image: &discordgo.MessageEmbedImage{
					URL: "https://orangeflare.me/imagehosting/Wumpagotchi/Salad.png",
				},
			})
		store.SetPageFooters()
		store.ColourWhenDone = 0x36393f
		store.Widget.Timeout = time.Minute * 1
		store.Spawn()
		return
	}
}
