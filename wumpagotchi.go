package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Build session
	session, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
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
