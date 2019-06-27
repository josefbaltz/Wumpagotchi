package main

import (
	"bytes"
	"fmt"
	"image/png"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// leftHandler decides which story the user will get depending on the Wumpus' age
// Also decides whether or not the user will be able to claim an egg or not
func leftHandler(UserWumpus Wumpus, event *discordgo.MessageCreate, session *discordgo.Session) {
	//Changed all from sendMessage to the built in method, the time was too short for users to read the message, may re-write the sendMessage function later to support custom sleep times
	if UserWumpus.Health <= 0 {
		session.ChannelMessageDelete(event.ChannelID, event.Message.ID)
		SentMessage, _ := session.ChannelMessageSend(event.ChannelID, "Wumpus in the Hospital\nTo start over with a new Wumpus, type 'w.adopt'")
		time.Sleep(time.Second * 20)
		session.ChannelMessageDelete(SentMessage.ChannelID, SentMessage.ID)
		return
	}
	if UserWumpus.Age > 9 {
		session.ChannelMessageDelete(event.ChannelID, event.Message.ID)
		SentMessage, _ := session.ChannelMessageSend(event.ChannelID, UserWumpus.Name+" has something important to tell you, They were accepted into Discordiversity and will be studying Wumpology, basically how to be a True Discord Wumpus, With a full ride scholarship! The wumpus shares how they loved all of the time they spent with you.\nYou can type 'w.claim' to claim your Wumpus' egg, or you can type 'w.adopt' to restart with a whole new Wumpus")
		time.Sleep(time.Second * 20)
		session.ChannelMessageDelete(SentMessage.ChannelID, SentMessage.ID)
		return
	}
	if UserWumpus.Age > 4 && UserWumpus.Age < 10 {
		session.ChannelMessageDelete(event.ChannelID, event.Message.ID)
		SentMessage, _ := session.ChannelMessageSend(event.ChannelID, UserWumpus.Name+" wants to talk with you. They enjoyed the time that they spent with you, but want to pursue greener pastures. They’ll be packing their things and leaving soon in search of one.\nTo start over with a new Wumpus, type 'w.adopt'")
		time.Sleep(time.Second * 20)
		session.ChannelMessageDelete(SentMessage.ChannelID, SentMessage.ID)
		return
	}
	session.ChannelMessageDelete(event.ChannelID, event.Message.ID)
	SentMessage, _ := session.ChannelMessageSend(event.ChannelID, "You can’t seem to find the wumpus anywhere. You don’t worry too much and head for the fridge for a quick snack. As you head towards the fridge you see a note addressed to "+event.Author.Username+". You open and read, Hey "+event.Author.Username+", I’m sorry but i’ve decided to leave without telling you first, all I wanted was a friend but I’m constantly stressed out living with you.\nTo start over with a new Wumpus, type 'w.adopt'")
	time.Sleep(time.Second * 20)
	session.ChannelMessageDelete(SentMessage.ChannelID, SentMessage.ID)
	return
}

// claimHandler checks to see if the user is eligible to claim a new Wumpus if so It generates a wumpus as normal but the name is predetermined and the color is also calculated, the user also maintains their credits
// The name is based off of the original name for example if the user's Wumpus' name was Wump the new name would be Wump Jr.
// The color is 80% of the original Wumpus and 20% of a randomly generated Base16 Number which can go as high as 0xFFFFFF
func claimHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	messageContent := strings.Split(strings.ToLower(event.Content), " ")
	if messageContent[0] == CommandPrefix+"claim" && !event.Author.Bot {
		UserWumpus, err := GetWumpus(event.Author.ID, false)
		if err != nil {
			go sendMessage(session, event, event.ChannelID, "There is nothing to claim!")
			return
		}
		if UserWumpus.Age > 9 && UserWumpus.Left && UserWumpus.Health > 0 {
			rand.Seed(time.Now().UnixNano())
			newColor := rand.Intn(0xFFFFFF + 1)
			weightedColor := int((float64(UserWumpus.Color) * 0.8) + (float64(newColor) * 0.2))
			if weightedColor > 0xFFFFFF {
				weightedColor = 0xFFFFFF
			}
			NewWumpus := Wumpus{
				Credits:   UserWumpus.Credits,
				Name:      UserWumpus.Name + " Jr.",
				Color:     weightedColor,
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
			err := png.Encode(&b, LeafedWumpus("https://orangeflare.me/imagehosting/Wumpagotchi/Happy.png", false, NewWumpus))
			if err != nil {
				fmt.Println(err)
				return
			}
			ClaimMessage := &discordgo.MessageSend{
				Embed: &discordgo.MessageEmbed{
					Color: NewWumpus.Color,
					Title: NewWumpus.Name,
					Fields: []*discordgo.MessageEmbedField{
						&discordgo.MessageEmbedField{
							Name:   "Congrats!",
							Value:  "You have claimed " + NewWumpus.Name + " as your new Wumpus!",
							Inline: false,
						},
					},
					Image: &discordgo.MessageEmbedImage{
						URL: "attachment://" + WumpusImageFile.Name,
					},
				},
				Files: []*discordgo.File{WumpusImageFile},
			}
			SentMessage, _ := session.ChannelMessageSendComplex(event.ChannelID, ClaimMessage)
			time.Sleep(15 * time.Second)
			session.ChannelMessageDelete(SentMessage.ChannelID, SentMessage.ID)
			return
		}
		go sendMessage(session, event, event.ChannelID, "There is nothing to claim!")
		return
	}
}
