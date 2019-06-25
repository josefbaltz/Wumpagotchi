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
	if messageContent[0] == CommandPrefix+"adopt" && !event.Author.Bot {
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
	if messageContent[0] == CommandPrefix+"view" && !event.Author.Bot {
		UserWumpus, err := GetWumpus(event.Author.ID)
		if err != nil {
			sendMessage(session, event, event.ChannelID, "You need a Wumpus first!")
			return
		}
		var StateURL = " "
		var State = " "
		if UserWumpus.Energy > 7 {
			State = "Hyper"
			StateURL = "https://orangeflare.me/imagehosting/Wumpagotchi/Happy.png"
		}
		if UserWumpus.Happiness > 7 {
			State = "Ecstatic"
			StateURL = "https://orangeflare.me/imagehosting/Wumpagotchi/Happy.png"
		}
		if UserWumpus.Energy > 8 && UserWumpus.Happiness > 8 && UserWumpus.Health > 8 && UserWumpus.Hunger > 8 && UserWumpus.Sick == false && UserWumpus.Sleeping == false && UserWumpus.Age > 1 {
			State = "Joyous (+10Ꞡ every 2 hours)"
			StateURL = "https://orangeflare.me/imagehosting/Wumpagotchi/Glorious.png"
		}
		if UserWumpus.Energy < 4 {
			State = "Hurt"
			StateURL = "https://orangeflare.me/imagehosting/Wumpagotchi/Tired.png"
		}
		if UserWumpus.Health < 4 {
			State = "Hurt"
			StateURL = "https://orangeflare.me/imagehosting/Wumpagotchi/Sad.png"
		}
		if UserWumpus.Happiness < 4 {
			State = "Depressed"
			StateURL = "https://orangeflare.me/imagehosting/Wumpagotchi/Sad.png"
		}
		if UserWumpus.Hunger < 4 {
			State = "Hungry"
			StateURL = "https://orangeflare.me/imagehosting/Wumpagotchi/Sad.png"
		}
		if UserWumpus.Happiness < 2 {
			State = "Depressed"
			StateURL = "https://orangeflare.me/imagehosting/Wumpagotchi/Depressed.png"
		}
		if UserWumpus.Hunger == 0 {
			State = "Starving"
			StateURL = "https://orangeflare.me/imagehosting/Wumpagotchi/Sad.png"
		}
		if UserWumpus.Sick {
			State = "Sick"
			StateURL = "https://orangeflare.me/imagehosting/Wumpagotchi/Sick.png"
		}
		if UserWumpus.Sleeping {
			State = "Sleeping"
			StateURL = "https://orangeflare.me/imagehosting/Wumpagotchi/Asleep.png"
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
			Image: &discordgo.MessageEmbedImage{
				URL: StateURL,
			},
		}
		sendEmbed(session, event, event.ChannelID, ViewEmbed)
		return
	}
	if messageContent[0] == CommandPrefix+"play" && !event.Author.Bot {
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
				URL: "https://orangeflare.me/imagehosting/Wumpagotchi/Gamer.png",
			},
		}
		SentMessage, err := session.ChannelMessageSendEmbed(event.ChannelID, GameEmbed)
		if err != nil {
			fmt.Println("ya hecked up lol, here's the thing\n" + err.Error())
			return
		}
		for i := 0; i <= 2; i++ {
			time.Sleep(2 * time.Second)
			wumpusGuess := rand.Intn(6)
			if wumpusGuess == gemSpot {
				GameEmbed.Fields[gemSpot].Name = "❗"
				GameEmbed.Fields[gemSpot].Value = "💎"
				GameEmbed.Image.URL = "https://orangeflare.me/imagehosting/Wumpagotchi/EpicGamer.png"
				session.ChannelMessageEditEmbed(SentMessage.ChannelID, SentMessage.ID, GameEmbed)
				sendMessage(session, event, event.ChannelID, UserWumpus.Name+" found a gem!\n+20Ꞡ\n+2 Happiness\n-2 Energy")
				UserWumpus.Credits += 30
				UserWumpus.Happiness += 2
				UpdateWumpus(event.Author.ID, UserWumpus)
				break
			}
			GameEmbed.Fields[wumpusGuess].Name = "..."
			GameEmbed.Fields[wumpusGuess].Value = "⬛"
			session.ChannelMessageEditEmbed(SentMessage.ChannelID, SentMessage.ID, GameEmbed)
			if i == 2 {
				sendMessage(session, event, event.ChannelID, "No gems found!\n-10Ꞡ\n-2 Energy")
				UpdateWumpus(event.Author.ID, UserWumpus)
				break
			}
		}
		return
	}
}
