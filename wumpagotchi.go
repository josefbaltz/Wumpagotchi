package main

import (
	"context"
	"encoding/base64"
	"log"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/firestore"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// MAIN
func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Build session
	ses, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Build Firestore client
	ctx := context.Background()
	d, _ := base64.StdEncoding.DecodeString(os.Getenv("GCP_CREDS_BASE64"))
	fs, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"), option.WithCredentialsJSON(d))
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close()

	// TODO: Add handlers to session

	// Create and open websocket for Discord API using created session
	err = ses.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer ses.Close()

	// Wait for termination signal (Ctrl+C, etc)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// -- HANDLER FUNCTIONS

// CONNECT

// DISCONNECT

// READY

// RESUMED

// COMMANDHANDLER

// -- COMMAND FUNCTIONS

// -- DATABASE FUNCTIONS

// -- LOGIC FUNCTIONS
