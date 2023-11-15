package main

import (
	"context"
	"encoding/base64"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/firestore"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// WUMPUS STRUCT
type Wumpus struct {
	name      string
	age       int
	color     int
	credits   int
	energy    int
	happiness int
	health    int
	hunger    int
	sick      bool
	sleeping  bool
	gone      bool
}

// MAIN
func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[GODOTENV] error loading .env file |", err)
		os.Exit(1)
	}

	// Build session
	ses, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal("[DISCORDGO] error creating discord session |", err)
		os.Exit(1)
	}

	// Build Firestore client
	ctx := context.Background()
	d, _ := base64.StdEncoding.DecodeString(os.Getenv("GCP_CREDS_BASE64"))
	fs, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"), option.WithCredentialsJSON(d))
	if err != nil {
		log.Fatal("[FIRESTORE] error making new client |", err)
		os.Exit(1)
	}
	defer fs.Close()

	// Register callbacks
	discord.AddHandler(ready)
	discord.AddHandler(connect)
	discord.AddHandler(disconnect)
	discord.AddHandler(resumed)
	discord.AddHandler(commandHandler)

	// Create slash commands
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "adopt",
			Description: "adopt your very own wumpus",
			Type:        discordgo.ChatApplicationCommand,
		},
		{
			Name:        "buy",
			Description: "buy an item",
			Type:        discordgo.ChatApplicationCommand,
		},
		{
			Name:        "credit",
			Description: "lists the people that have helped make this bot",
			Type:        discordgo.ChatApplicationCommand,
		},
		{
			Name:        "invite",
			Description: "sends an invite link for wumpagotchi",
			Type:        discordgo.ChatApplicationCommand,
		},
		{
			Name:        "play",
			Description: "mine with your wumpus",
			Type:        discordgo.ChatApplicationCommand,
		},
		{
			Name:        "view",
			Description: "view your wumpus",
			Type:        discordgo.ChatApplicationCommand,
		},
	}

	// Create and open websocket for Discord API using created session
	err = ses.Open()
	if err != nil {
		log.Fatal("[DISCORDGO] error opening discord session |", err)
		os.Exit(1)
	}
	defer ses.Close()

	// Write commands
	_, err := discord.ApplicationCommandBulkOverwrite(ses.State.User.ID, "", commands)
	if err != nil {
		log.Fatal("[COMMAND WRITE] error writing commands |", err)
		os.Exit(1)
	}

	// Wait for termination signal (Ctrl+C, etc)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// -- HANDLER FUNCTIONS

// CONNECT
func connect(s *discordgo.Session, event *discordgo.Connect) {
	log.Println("[CONNECTED]")
}

// READY
func ready(s *discordgo.Session, event *discordgo.Ready) {
	log.Println("[READY]")
}

// DISCONNECT
func disconnect(s *discordgo.Session, event *discordgo.Disconnect) {
	log.Println("[DISCONNECTED]")
}

// RESUMED
func resumed(s *discordgo.Session, event *discordgo.Resumed) {
	log.Println("[RESUMED]")
}

// COMMAND HANDLER
func commandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractiionApplicationCommand {
		return
	}
	switch i.ApplicationCommandData().Name {
	case "adopt":
		adopt(s, i)
	case "buy":
		buy(s, i)
	case "credit":
		credit(s, i)
	case "invite":
		invite(s, i)
	case "play":
		play(s, i)
	case "view":
		view(s, i)
	}
}

// -- BOT COMMANDS
// adopt
// buy [item]
// credit
// invite
// play
// view

// ADOPT
func adopt(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var UID
	if i.User != nil {
		UID = i.User.ID
	} else {
		UID = i.Member.User.ID
	}

	if 
	rand.Seed(time.Now().UnixNano())
	newColor := rand.Intn(0xFFFFFF + 1)
}

// BUY
func buy(s *discordgo.Session, i *discordgo.InteractionCreate) {
	item := i.ApplicationCommandData().Options[0].Options[0].Value.(string)
}

// CREDIT
func credit(s *discordgo.Session, i *discordgo.InteractionCreate) {

}

// INVITE
func invite(s *discordgo.Session, i *discordgo.InteractionCreate) {

}

// PLAY
func play(s *discordgo.Session, i *discordgo.InteractionCreate) {

}

// VIEW
func view(s *discordgo.Session, i *discordgo.InteractionCreate) {

}

// -- GCP FUNCTIONS

// TIME
func time() {
	return
}
