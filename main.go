package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
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
