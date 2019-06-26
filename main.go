package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
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

//DiscordToken - Discord API Token
var DiscordToken string
var CommandPrefix = "w."
var gcp *datastore.Client
var gcpErr error
var ctx = context.Background()

func main() {
	if DiscordToken == "" {
		fmt.Println("==Wumpagotchi Error==\nYour start command should be as follows:\nWumpagotchi -t <Discord API Token>")
		os.Exit(0)
	}

	gcp, gcpErr = datastore.NewClient(ctx, "wumpagotchi", option.WithCredentialsFile("./WumpagotchiCredentials.json"))
	if gcpErr != nil {
		fmt.Println("==Datastore Error==\nFailed to create GCP Client\n" + gcpErr.Error())
		os.Exit(1)
	}

	wump, err := discordgo.New("Bot " + DiscordToken)
	if err != nil {
		fmt.Println("==Wumpagotchi Error==\n" + err.Error())
		os.Exit(1)
	}

	wump.AddHandler(basicCommands)
	wump.AddHandler(game)
	wump.AddHandler(messageCredits)
	wump.AddHandler(claimHandler)

	err = wump.Open()
	if err != nil {
		fmt.Println("==Wumpagotchi Error==\n" + err.Error())
	}

	fmt.Println("Wumpagotchi Online\nRunning until a termination singal is recieved ...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	wump.Close()
}

func sendMessage(session *discordgo.Session, event *discordgo.MessageCreate, channel string, message string) {
	sentMessage, _ := session.ChannelMessageSend(channel, message)
	time.Sleep(10 * time.Second)
	err := session.ChannelMessageDelete(event.ChannelID, event.ID)
	if err != nil {
		fmt.Println(err)
	}
	session.ChannelMessageDelete(sentMessage.ChannelID, sentMessage.ID)
}

func sendEmbed(session *discordgo.Session, event *discordgo.MessageCreate, channel string, embed *discordgo.MessageEmbed) {
	sentMessage, err := session.ChannelMessageSendEmbed(channel, embed)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(15 * time.Second)
	err = session.ChannelMessageDelete(event.ChannelID, event.ID)
	if err != nil {
		fmt.Println(err)
	}
	session.ChannelMessageDelete(sentMessage.ChannelID, sentMessage.ID)
}
