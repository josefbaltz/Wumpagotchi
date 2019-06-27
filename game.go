package main

import (
	"bytes"
	"fmt"
	"image/png"
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
		session.ChannelMessageDelete(event.ChannelID, event.Message.ID)
		if UserWumpus, err := GetWumpus(event.Author.ID, true); err != nil || UserWumpus.Left == true {
			if len(messageContent) > 1 {
				if len(strings.TrimPrefix(event.Content, CommandPrefix+"adopt ")) <= 15 {
					rand.Seed(time.Now().UnixNano())
					newColor := rand.Intn(0xFFFFFF + 1)
					NewWumpus := Wumpus{
						Credits:   0,
						Name:      strings.TrimPrefix(event.Content, CommandPrefix+"adopt "),
						Color:     newColor,
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
					var b bytes.Buffer
					WumpusImageFile := &discordgo.File{
						Name:        "Wumpus.png",
						ContentType: "image/png",
						Reader:      &b,
					}
					err = png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/Happy.png", false, NewWumpus))
					if err != nil {
						fmt.Println(err)
						return
					}
					AdoptMessage := &discordgo.MessageSend{
						Embed: &discordgo.MessageEmbed{
							Color: NewWumpus.Color,
							Title: NewWumpus.Name,
							Fields: []*discordgo.MessageEmbedField{
								&discordgo.MessageEmbedField{
									Name:   "Congrats!",
									Value:  "You have adopted " + NewWumpus.Name + " as your Wumpus!",
									Inline: false,
								},
							},
							Image: &discordgo.MessageEmbedImage{
								URL: "attachment://" + WumpusImageFile.Name,
							},
						},
						Files: []*discordgo.File{WumpusImageFile},
					}
					SentMessage, _ := session.ChannelMessageSendComplex(event.ChannelID, AdoptMessage)
					session.ChannelMessageDelete(SentMessage.ChannelID, SentMessage.ID)
					return
				}
				go sendMessage(session, event, event.ChannelID, "Your Wumpus' name must be 15 characters or less!")
				return
			}
			go sendMessage(session, event, event.ChannelID, "Your Wumpus needs a name to be adopted!")
			return
		} else {
			go sendMessage(session, event, event.ChannelID, "You already have a Wumpus, and their name is "+UserWumpus.Name+"!")
			return
		}
	}
	if messageContent[0] == CommandPrefix+"view" && !event.Author.Bot {
		session.ChannelMessageDelete(event.ChannelID, event.Message.ID)
		UserWumpus, err := GetWumpus(event.Author.ID, false)
		if err != nil {
			go sendMessage(session, event, event.ChannelID, "You need a Wumpus first!")
			return
		}
		if UserWumpus.Left == true {
			leftHandler(UserWumpus, event, session)
			return
		}
		var State = " "
		var b bytes.Buffer
		if UserWumpus.Sleeping {
			State = "Sleeping"
			if err := png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/Asleep.png", true, UserWumpus)); err != nil {
				fmt.Println(err)
				return
			}
		} else if UserWumpus.Sick {
			State = "Sick"
			if err := png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/Sick.png", false, UserWumpus)); err != nil {
				fmt.Println(err)
				return
			}
		} else if UserWumpus.Hunger == 0 {
			State = "Starving"
			if err := png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/Sad.png", false, UserWumpus)); err != nil {
				fmt.Println(err)
				return
			}
		} else if UserWumpus.Happiness <= 1 {
			State = "Depressed"
			if err := png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/Depressed.png", false, UserWumpus)); err != nil {
				fmt.Println(err)
				return
			}
		} else if UserWumpus.Hunger <= 3 {
			State = "Hungry"
			if err := png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/Sad.png", false, UserWumpus)); err != nil {
				fmt.Println(err)
				return
			}
		} else if UserWumpus.Happiness <= 3 {
			State = "Sad"
			if err := png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/Sad.png", false, UserWumpus)); err != nil {
				fmt.Println(err)
				return
			}
		} else if UserWumpus.Health <= 3 {
			State = "Hurt"
			err := png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/Sad.png", false, UserWumpus))
			if err != nil {
				fmt.Println(err)
				return
			}
		} else if UserWumpus.Energy <= 3 {
			State = "Tired"
			err := png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/Tired.png", false, UserWumpus))
			if err != nil {
				fmt.Println(err)
				return
			}
		} else if UserWumpus.Energy > 8 && UserWumpus.Happiness > 8 && UserWumpus.Health > 8 && UserWumpus.Hunger > 8 && UserWumpus.Sick == false && UserWumpus.Sleeping == false && UserWumpus.Age > 1 {
			State = "Joyous (+10Ꞡ every 2 hours)"
			err := png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/Joyous.png", false, UserWumpus))
			if err != nil {
				fmt.Println(err)
				return
			}
		} else if UserWumpus.Happiness > 7 {
			State = "Ecstatic"
			err := png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/Happy.png", false, UserWumpus))
			if err != nil {
				fmt.Println(err)
				return
			}
		} else if UserWumpus.Energy > 7 {
			State = "Hyper"
			err := png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/Happy.png", false, UserWumpus))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		if UserWumpus.Sick && UserWumpus.Sleeping {
			State = "Sleeping (Sick)"
		}

		WumpusImageFile := &discordgo.File{
			Name:        "Wumpus.png",
			ContentType: "image/png",
			Reader:      &b,
		}

		ms := &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Color: UserWumpus.Color,
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
					URL: "attachment://" + WumpusImageFile.Name,
				},
			},
			Files: []*discordgo.File{WumpusImageFile},
		}
		SentMessage, _ := session.ChannelMessageSendComplex(event.ChannelID, ms)
		time.Sleep(15 * time.Second)
		session.ChannelMessageDelete(SentMessage.ChannelID, SentMessage.ID)
		return
	}
	if messageContent[0] == CommandPrefix+"play" && !event.Author.Bot {
		UserWumpus, err := GetWumpus(event.Author.ID, true)
		if err != nil {
			go sendMessage(session, event, event.ChannelID, "You need a Wumpus first!")
			return
		}
		if LeftCheck(UserWumpus, session, event) {
			return
		}
		if EnergyCheck(UserWumpus, 2, session, event) {
			return
		}
		if CreditCheck(UserWumpus, 10, session, event) {
			return
		}
		UserWumpus = SleepCheck(UserWumpus, session, event)
		if UserWumpus.Sleeping {
			UpdateWumpus(event.Author.ID, UserWumpus)
			return
		}
		UserWumpus.Energy -= 2
		UserWumpus.Credits -= 10
		UpdateWumpus(event.Author.ID, UserWumpus)
		rand.Seed(time.Now().UnixNano())
		var b bytes.Buffer
		WumpusImageFile := &discordgo.File{
			Name:        "Wumpus.png",
			ContentType: "image/png",
			Reader:      &b,
		}
		gemSpot := rand.Intn(6)
		ms := &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Color: UserWumpus.Color,
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
					URL: "attachment://" + WumpusImageFile.Name,
				},
			},
			Files: []*discordgo.File{WumpusImageFile},
		}
		if err := png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/Gamer.png", false, UserWumpus)); err != nil {
			fmt.Println(err)
			return
		}
		SentMessage, err := session.ChannelMessageSendComplex(event.ChannelID, ms)
		if err != nil {
			fmt.Println("ya hecked up lol, here's the thing\n" + err.Error())
			return
		}
		for i := 0; i <= 2; i++ {
			time.Sleep(2 * time.Second)
			wumpusGuess := rand.Intn(6)
			if wumpusGuess == gemSpot {
				ms.Embed.Fields[gemSpot].Name = "❗"
				ms.Embed.Fields[gemSpot].Value = "💎"
				if err := png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/EpicGamer.png", false, UserWumpus)); err != nil {
					fmt.Println(err)
					return
				}
				msWon := &discordgo.MessageSend{
					Embed: ms.Embed,
					Files: []*discordgo.File{WumpusImageFile},
				}
				session.ChannelMessageDelete(SentMessage.ChannelID, SentMessage.ID)
				SentMessageWon, _ := session.ChannelMessageSendComplex(event.ChannelID, msWon)
				sendMessage(session, event, event.ChannelID, UserWumpus.Name+" found a gem!\n+20Ꞡ\n+2 Happiness\n-2 Energy")
				UserWumpus.Credits += 30
				UserWumpus.Happiness += 2
				if UserWumpus.Happiness > 10 {
					UserWumpus.Happiness = 10
				}
				UpdateWumpus(event.Author.ID, UserWumpus)
				time.Sleep(15 * time.Second)
				session.ChannelMessageDelete(SentMessageWon.ChannelID, SentMessageWon.ID)
				break
			}
			ms.Embed.Fields[wumpusGuess].Name = "..."
			ms.Embed.Fields[wumpusGuess].Value = "⬛"
			me := &discordgo.MessageEdit{
				Embed:   ms.Embed,
				ID:      SentMessage.ID,
				Channel: SentMessage.ChannelID,
			}
			session.ChannelMessageEditComplex(me)
			if i == 2 {
				sendMessage(session, event, event.ChannelID, "No gems found!\n-10Ꞡ\n-2 Energy")
				UpdateWumpus(event.Author.ID, UserWumpus)
				time.Sleep(15 * time.Second)
				session.ChannelMessageDelete(SentMessage.ChannelID, SentMessage.ID)
				break
			}
		}
		return
	}
}
