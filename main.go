package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

func main() {
	if DiscordToken == "" {
		fmt.Println("==Wumpagotchi Error==\nYour start command should be as follows:\nWumpagotchi -t <Discord API Token>")
		os.Exit(0)
	}

	ctx = context.Background()
	gcp, gcpErr := datastore.NewClient(ctx, "wumpagotchi", option.WithCredentialsFile("./WumpagotchiCredentials.json"))
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
