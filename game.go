package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

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
	Sleeping  bool
	Left      bool
}

func game(session *discordgo.Session, event *discordgo.MessageCreate) {
	messageContent := strings.Split(strings.ToLower(event.Content), " ")
	if messageContent[0] == CommandPrefix+"adopt" {
		if UserWumpus, err := GetWumpus(event.Author.ID); err != nil {
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
					Sleeping:  false,
					Left:      false,
				}
				UpdateWumpus(event.Author.ID, NewWumpus)
				sendMessage(session, event, event.ChannelID, "Congrats, you have adopted "+NewWumpus.Name+" as your Wumpus!")
				return
			} else {
				sendMessage(session, event, event.ChannelID, "Your Wumpus needs a name to be adopted!")
				return
			}
		} else {
			sendMessage(session, event, event.ChannelID, "You already have a Wumpus, and their name is "+UserWumpus.Name+"!")
			return
		}
	}
	if messageContent[0] == CommandPrefix+"view" {
		UserWumpus, err := GetWumpus(event.Author.ID)
		if err != nil {
			sendMessage(session, event, event.ChannelID, "Something went wrong, please contact the devs!")
			return
		}
		var State string
		WumpusImage := &discordgo.MessageEmbedImage{
			URL: "https://orangeflare.me/imagehosting/Wumpagotchi/Happy.png",
		}
		if UserWumpus.Energy > 7 {
			State = "Hyper"
			WumpusImage.URL = "https://orangeflare.me/imagehosting/Wumpagotchi/Happy.png"
			return
		}
		if UserWumpus.Happiness > 7 {
			State = "Ecstatic"
			WumpusImage.URL = "https://orangeflare.me/imagehosting/Wumpagotchi/Happy.png"
			return
		}
		if UserWumpus.Energy < 8 && UserWumpus.Happiness < 8 && UserWumpus.Health < 8 && UserWumpus.Hunger < 8 && UserWumpus.Sick == false && UserWumpus.Sleeping == false && UserWumpus.Age > 1 {
			State = "Joyous (+10Ꞡ every 2 hours)"
			WumpusImage.URL = "https://orangeflare.me/imagehosting/Wumpagotchi/Glorious.png"
			return
		}
		if UserWumpus.Energy < 4 {
			State = "Hurt"
			WumpusImage.URL = "https://orangeflare.me/imagehosting/Wumpagotchi/Tired.png"
			return
		}
		if UserWumpus.Health < 4 {
			State = "Hurt"
			WumpusImage.URL = "https://orangeflare.me/imagehosting/Wumpagotchi/Sad.png"
			return
		}
		if UserWumpus.Happiness < 4 {
			State = "Depressed"
			WumpusImage.URL = "https://orangeflare.me/imagehosting/Wumpagotchi/Sad.png"
			return
		}
		if UserWumpus.Hunger < 4 {
			State = "Hungry"
			WumpusImage.URL = "https://orangeflare.me/imagehosting/Wumpagotchi/Sad.png"
			return
		}
		if UserWumpus.Happiness < 2 {
			State = "Depressed"
			WumpusImage.URL = "https://orangeflare.me/imagehosting/Wumpagotchi/Depressed.png"
			return
		}
		if UserWumpus.Hunger == 0 {
			State = "Starving"
			WumpusImage.URL = "https://orangeflare.me/imagehosting/Wumpagotchi/Sad.png"
			return
		}
		if UserWumpus.Sick {
			State = "Sick"
			WumpusImage.URL = "https://orangeflare.me/imagehosting/Wumpagotchi/Sick.png"
			return
		}
		if UserWumpus.Sleeping {
			State = "Sleeping"
			WumpusImage.URL = "https://orangeflare.me/imagehosting/Wumpagotchi/Asleep.png"
			return
		}
		ViewEmbed := &discordgo.MessageEmbed{
			Color: 0x669966, //Wumpus Leaf Green
			Title: UserWumpus.Name,
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Credits",
					Value:  strconv.Itoa(UserWumpus.Credits) + "Ꞡ",
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "Status",
					Value:  State,
					Inline: false,
				},
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
			Image: WumpusImage,
		}
		sendEmbed(session, event, event.ChannelID, ViewEmbed)
		return
	}
	if messageContent[0] == CommandPrefix+"play" {
		UserWumpus, err := GetWumpus(event.Author.ID)
		if err != nil {
			sendMessage(session, event, event.ChannelID, "You need a Wumpus first, they are always looking for a friend!")
			return
		}
		if UserWumpus.Energy <= 2 {
			sendMessage(session, event, event.ChannelID, UserWumpus.Name+" doesn't have enough energy to play!")
			return
		}
		if UserWumpus.Credits < 10 {
			sendMessage(session, event, event.ChannelID, "You need 10Ꞡ to play!")
			return
		}
		UserWumpus.Energy -= 2
		UserWumpus.Credits -= 10
		rand.Seed(time.Now().UnixNano())
		gemSpot := rand.Intn(6)
		GameEmbed := &discordgo.MessageEmbed{
			Color: 0x669966, //Wumpus Leaf Green
			Title: "Find the gem!",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "⛏️",
					Value:  "⬜",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "⛏️",
					Value:  "⬜",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "⛏️",
					Value:  "⬜",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "⛏️",
					Value:  "⬜",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "⛏️",
					Value:  "⬜",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "⛏️",
					Value:  "⬜",
					Inline: true,
				},
			},
			Image: &discordgo.MessageEmbedImage{
				URL: "https://i.redd.it/vj6r64pcee711.gif",
			},
		}
		SentMessage, err := session.ChannelMessageSendEmbed(event.ChannelID, GameEmbed)
		if err != nil {
			fmt.Println("ya hecked up lol, here's the thing\n" + err.Error())
			return
		}
		for i := 0; i <= 2; i++ {
			wumpusGuess := rand.Intn(6)
			if wumpusGuess == gemSpot {
				GameEmbed.Fields[gemSpot].Name = "❗"
				GameEmbed.Fields[gemSpot].Value = "💎"
				session.ChannelMessageEditEmbed(SentMessage.ChannelID, SentMessage.ID, GameEmbed)
				sendMessage(session, event, event.ChannelID, UserWumpus.Name+" found a gem!")
				UserWumpus.Credits += 30
				UserWumpus.Happiness += 2
				sendMessage(session, event, event.ChannelID, "+20Ꞡ")
				sendMessage(session, event, event.ChannelID, "+2 Happiness")
				sendMessage(session, event, event.ChannelID, "-2 Energy")
				UpdateWumpus(event.Author.ID, UserWumpus)
				break
			}
			GameEmbed.Fields[wumpusGuess].Name = "..."
			GameEmbed.Fields[wumpusGuess].Value = "⬛"
			session.ChannelMessageEditEmbed(SentMessage.ChannelID, SentMessage.ID, GameEmbed)
			time.Sleep(2 * time.Second)
			if i == 2 {
				sendMessage(session, event, event.ChannelID, "No gems found!")
				sendMessage(session, event, event.ChannelID, "-10Ꞡ")
				sendMessage(session, event, event.ChannelID, "-2 Energy")
				UpdateWumpus(event.Author.ID, UserWumpus)
				break
			}
		}
		return
	}
}
