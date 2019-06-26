package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func basicCommands(session *discordgo.Session, event *discordgo.MessageCreate) {
	messageContent := strings.Split(strings.ToLower(event.Content), " ")
	if messageContent[0] == CommandPrefix+"help" && !event.Author.Bot {
		session.ChannelMessageDelete(event.ChannelID, event.Message.ID)
		HelpEmbed := &discordgo.MessageEmbed{
			Color:       0x669966, //Wumpus Leaf Green
			Title:       "Wumpagotchi Help",
			Description: "We're here to help!",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "What is Wumpagotchi?",
					Value:  "Wumpagotchi is Tamagotchi but with Wumpus!\nStart by adopting a Wumpus with w.adopt, and then take care of your Wumpus and play with them to keep them happy!",
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
					Name:   CommandPrefix + "buy <Item>",
					Value:  "Buy items for your Wumpus! To view what items you can buy, just say w.buy",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   CommandPrefix + "view",
					Value:  "See how your Wumpus is doing",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   CommandPrefix + "play",
					Value:  "Go mining with your Wumpus!",
					Inline: true,
				},
			},
		}
		go sendEmbed(session, event, event.ChannelID, HelpEmbed)
		return
	}
	// Store commands
	if messageContent[0] == CommandPrefix+"buy" && !event.Author.Bot {
		session.ChannelMessageDelete(event.ChannelID, event.Message.ID)
		UserWumpus, err := GetWumpus(event.Author.ID, true)
		if err != nil {
			go sendMessage(session, event, event.ChannelID, "You need a Wumpus first!")
			return
		}
		if LeftCheck(UserWumpus, session, event) {
			return
		}
		messageContent := strings.Split(strings.ToLower(event.Content), " ")
		if len(messageContent) > 1 {
			// Floop buy code
			if strings.ToLower(strings.TrimPrefix(event.Content, CommandPrefix+"buy ")) == "floop" {
				checkReturn := CreditCheck(UserWumpus, 5, session, event)
				if checkReturn {
					return
				}
				UserWumpus = SleepCheck(UserWumpus, session, event)
				if UserWumpus.Sleeping {
					UpdateWumpus(event.Author.ID, UserWumpus)
					return
				}
				checkReturn = EnergyCheck(UserWumpus, 1, session, event)
				if checkReturn {
					return
				}
				UserWumpus.Credits -= 5
				UserWumpus.Hunger++
				UserWumpus.Energy--
				UserWumpus = LogicKeeper(UserWumpus)
				UpdateWumpus(event.Author.ID, UserWumpus)
				BuyEmbed := &discordgo.MessageEmbed{
					Color: UserWumpus.Color,
					Title: "Floop bought!",
					Fields: []*discordgo.MessageEmbedField{
						&discordgo.MessageEmbedField{
							Name:   "Remaining Credits",
							Value:  strconv.Itoa(UserWumpus.Credits) + "Ꞡ",
							Inline: false,
						},
						&discordgo.MessageEmbedField{
							Name:   "Stats Affected:",
							Value:  "+1 hunger\n-1 energy",
							Inline: false,
						},
					},
					Image: &discordgo.MessageEmbedImage{
						URL: "https://orangeflare.me/imagehosting/Wumpagotchi/Floop.png",
					},
				}
				go sendEmbed(session, event, event.ChannelID, BuyEmbed)
				return
			}

			// Gummy gem buy code
			if strings.ToLower(strings.TrimPrefix(event.Content, CommandPrefix+"buy ")) == "gummy" || strings.ToLower(strings.TrimPrefix(event.Content, CommandPrefix+"buy ")) == "gummy gem" {
				checkReturn := CreditCheck(UserWumpus, 10, session, event)
				if checkReturn {
					return
				}
				UserWumpus = SleepCheck(UserWumpus, session, event)
				if UserWumpus.Sleeping {
					UpdateWumpus(event.Author.ID, UserWumpus)
					return
				}
				checkReturn = EnergyCheck(UserWumpus, 1, session, event)
				if checkReturn {
					return
				}
				UserWumpus.Credits -= 10
				UserWumpus.Hunger++
				UserWumpus.Happiness++
				UserWumpus.Energy--
				UserWumpus = LogicKeeper(UserWumpus)
				UpdateWumpus(event.Author.ID, UserWumpus)
				BuyEmbed := &discordgo.MessageEmbed{
					Color: UserWumpus.Color,
					Title: "Gummy Gem bought!",
					Fields: []*discordgo.MessageEmbedField{
						&discordgo.MessageEmbedField{
							Name:   "Remaining Credits",
							Value:  strconv.Itoa(UserWumpus.Credits) + "Ꞡ",
							Inline: false,
						},
						&discordgo.MessageEmbedField{
							Name:   "Stats Affected:",
							Value:  "+1 hunger\n+1 happiness\n-1 energy",
							Inline: false,
						},
					},
					Image: &discordgo.MessageEmbedImage{
						URL: "https://orangeflare.me/imagehosting/Wumpagotchi/GummyGem.png",
					},
				}
				go sendEmbed(session, event, event.ChannelID, BuyEmbed)
				rand.Seed(time.Now().UnixNano())
				sickChance := rand.Intn(10)
				if sickChance == 1 {
					UserWumpus.Sick = true
					UpdateWumpus(event.Author.ID, UserWumpus)
					go sendMessage(session, event, event.ChannelID, UserWumpus.Name+" has gotten sick from the gummy gem!")
					return
				}
				return
			}

			// Medicine buy code
			if strings.ToLower(strings.TrimPrefix(event.Content, CommandPrefix+"buy ")) == "medicine" {
				checkReturn := CreditCheck(UserWumpus, 15, session, event)
				if checkReturn {
					return
				}
				UserWumpus.Credits -= 15
				UserWumpus.Health += 2
				UserWumpus.Sick = false
				UserWumpus = LogicKeeper(UserWumpus)
				UpdateWumpus(event.Author.ID, UserWumpus)
				BuyEmbed := &discordgo.MessageEmbed{
					Color: UserWumpus.Color,
					Title: "Medicine bought!",
					Fields: []*discordgo.MessageEmbedField{
						&discordgo.MessageEmbedField{
							Name:   "Remaining Credits",
							Value:  strconv.Itoa(UserWumpus.Credits) + "Ꞡ",
							Inline: false,
						},
						&discordgo.MessageEmbedField{
							Name:   "Stats Affected:",
							Value:  "+2 health\nRecovers Wumpus from sickness",
							Inline: false,
						},
					},
					Image: &discordgo.MessageEmbedImage{
						URL: "https://orangeflare.me/imagehosting/Wumpagotchi/Medicine.png",
					},
				}
				go sendEmbed(session, event, event.ChannelID, BuyEmbed)
				return
			}

			// Salad buy code
			if strings.ToLower(strings.TrimPrefix(event.Content, CommandPrefix+"buy ")) == "salad" {
				checkReturn := CreditCheck(UserWumpus, 30, session, event)
				if checkReturn {
					return
				}
				UserWumpus = SleepCheck(UserWumpus, session, event)
				if UserWumpus.Sleeping {
					UpdateWumpus(event.Author.ID, UserWumpus)
					return
				}
				UserWumpus.Credits -= 30
				UserWumpus.Health += 2
				UserWumpus.Hunger += 3
				UserWumpus = LogicKeeper(UserWumpus)
				UpdateWumpus(event.Author.ID, UserWumpus)
				BuyEmbed := &discordgo.MessageEmbed{
					Color: UserWumpus.Color,
					Title: "Salad bought!",
					Fields: []*discordgo.MessageEmbedField{
						&discordgo.MessageEmbedField{
							Name:   "Remaining Credits",
							Value:  strconv.Itoa(UserWumpus.Credits) + "Ꞡ",
							Inline: false,
						},
						&discordgo.MessageEmbedField{
							Name:   "Stats Affected:",
							Value:  "+3 hunger\n+2 health",
							Inline: false,
						},
					},
					Image: &discordgo.MessageEmbedImage{
						URL: "https://orangeflare.me/imagehosting/Wumpagotchi/Salad.png",
					},
				}
				go sendEmbed(session, event, event.ChannelID, BuyEmbed)
				return
			}
			return
		}
		StoreEmbed := &discordgo.MessageEmbed{
			Color: UserWumpus.Color,
			Title: "Store",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Credits",
					Value:  strconv.Itoa(UserWumpus.Credits) + "Ꞡ",
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "Floop (5Ꞡ)",
					Value:  "A basic part of every Wumpus' diet!",
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "Gummy Gem (10Ꞡ)",
					Value:  "A delicious treat for your Wumpus! Be careful though as these can be unhealthy!",
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "Medicine (15Ꞡ)",
					Value:  "Used to make your Wumpus feel all better!",
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "Salad (30Ꞡ)",
					Value:  "A healthy meal that Wumpi love to eat!",
					Inline: false,
				},
			},
		}
		go sendEmbed(session, event, event.ChannelID, StoreEmbed)
		return
	}
	return
}
