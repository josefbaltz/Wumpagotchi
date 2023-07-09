package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/bwmarrin/discordgo"
	"google.golang.org/api/option"
)

func init() {
	flag.StringVar(&DiscordToken, "t", "", "Discord API Token")
	flag.Parse()
}

// DiscordToken - Discord API Token
var DiscordToken string

// CommandPrefix - This is a string that user's have to start their messages with to actually use the bot
var CommandPrefix = "w."

// gcp - Creating the Client object for GCP Datastore Requests
var gcp *datastore.Client

// gcpErr - Creating the Error object for GCP Datastore Errors
var gcpErr error

// ctx - Creating the Context object for GCP Datastore Requests
var ctx = context.Background()

func main() {
	if DiscordToken == "" {
		fmt.Println("==Wumpagotchi Error==\nYour start command should be as follows:\nWumpagotchi -t <Discord API Token>")
		os.Exit(0)
	}

	// Create the Datastore client so we can make requests to the Datastore
	// Additionally this also Connects to the GCP Datastore
	gcp, gcpErr = datastore.NewClient(ctx, "wumpagotchi", option.WithCredentialsFile("./WumpagotchiCredentials.json"))
	if gcpErr != nil {
		fmt.Println("==Datastore Error==\nFailed to create GCP Client\n" + gcpErr.Error())
		os.Exit(1)
	}

	// Create the Discord Bot Session so we can communicate with the Discord API
	// This logs in with the DiscordToken provided by the user at startup (-t <DiscordToken>)
	wump, err := discordgo.New("Bot " + DiscordToken)
	if err != nil {
		fmt.Println("==Wumpagotchi Error==\n" + err.Error())
		os.Exit(1)
	}

	// Add Event Handlers so whenever an event happens (ie A user sends a message, Joins a voice channel, deletes a message, etc.) the bot will run code
	wump.AddHandler(loginLogic)
	wump.AddHandler(basicCommands)
	wump.AddHandler(game)
	wump.AddHandler(messageCredits)
	wump.AddHandler(claimHandler)

	// Connect to the DiscordAPI with the Object we created earlier
	err = wump.Open()
	if err != nil {
		fmt.Println("==Wumpagotchi Error==\n" + err.Error())
	}

	// Wait until a Termination signal is recieved from the OS (Ctrl+C, Alt+F4) and then gracefully shutdown (Disconnect from the Discord API & GCP Datastore)
	fmt.Println("Wumpagotchi Online\nRunning until a termination singal is recieved ...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	wump.Close()
	gcp.Close()
}

// loginLogic is a function that handles when the bot logs in
func loginLogic(session *discordgo.Session, event *discordgo.Ready) {
	for 1 == 1 {
		query := datastore.NewQuery("User")
		var wumpus []Wumpus
		var err error
		_, err = gcp.GetAll(ctx, query, &wumpus)
		if err != nil {
			fmt.Println(err.Error())
		}
		err = session.UpdateGameStatus(0, "with "+strconv.Itoa(len(wumpus))+" Wumpi")
		if err != nil {
			fmt.Println("Error updating status:", err)
		}
		time.Sleep(20 * time.Second)
	}
}

// sendMessage is a function that will delete the message that fired the event, send a message, wait 10 seconds, and then delete the message that the bot sent
func sendMessage(session *discordgo.Session, event *discordgo.MessageCreate, channel string, message string) {
	sentMessage, err := session.ChannelMessageSend(channel, message)
	if err != nil {
		fmt.Println(err)
	}
	err = session.ChannelMessageDelete(event.ChannelID, event.ID)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(10 * time.Second)
	session.ChannelMessageDelete(sentMessage.ChannelID, sentMessage.ID)
}

// sendEmbed is a function that will delete the message that fired the event, send an embeded message, wait 15 seconds, and then delete the embed that the bot sent
func sendEmbed(session *discordgo.Session, event *discordgo.MessageCreate, channel string, embed *discordgo.MessageEmbed) {
	sentMessage, err := session.ChannelMessageSendEmbed(channel, embed)
	if err != nil {
		fmt.Println(err)
	}
	err = session.ChannelMessageDelete(event.ChannelID, event.ID)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(15 * time.Second)
	session.ChannelMessageDelete(sentMessage.ChannelID, sentMessage.ID)
}
