package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Declare and gather program arguments
var discordToken string

func init() {
	flag.StringVar(&discordToken, "t", "", "Discord API Token")
	flag.Parse()
}

func main() {
	// Terminate if no token provided
	if discordToken == "" {
		log.Fatal("No token provided")
	}

	// Build session
	session, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Add handlers to session

	// Create and open websocket for Discord API using created session
	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[Wumpagotchi] Online")

	// Wait for termination signal (Ctrl+C, etc)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Close websocket for Discord API
	session.Close()
	fmt.Println("[Wumpagotchi] Offline")
}
