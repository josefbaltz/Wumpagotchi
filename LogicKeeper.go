package main

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// LogicKeeper performs numerous checks on a Wumpus to make sure their stats
// are within range
//
// Requires 2 Arguments
// UserWumpus Wumpus should be the Wumpus you want to check and correct
//
// Returns CorrectedWumpus Wumpus which is the original Wumpus with values
// that are within range
func LogicKeeper(UserWumpus Wumpus) (CorrectedWumpus Wumpus) {
	if UserWumpus.Age > 14 {
		CorrectedWumpus.Age = 14
	} else if UserWumpus.Age < 0 {
		CorrectedWumpus.Age = 0
	} else {
		CorrectedWumpus.Age = UserWumpus.Age
	}

	if UserWumpus.Health > 10 {
		CorrectedWumpus.Health = 10
	} else if UserWumpus.Health < 0 {
		CorrectedWumpus.Health = 0
	} else {
		CorrectedWumpus.Health = UserWumpus.Health
	}

	if UserWumpus.Energy > 10 {
		CorrectedWumpus.Energy = 10
	} else if UserWumpus.Energy < 0 {
		CorrectedWumpus.Energy = 0
	} else {
		CorrectedWumpus.Energy = UserWumpus.Energy
	}

	if UserWumpus.Happiness > 10 {
		CorrectedWumpus.Happiness = 10
	} else if UserWumpus.Happiness < 0 {
		CorrectedWumpus.Happiness = 0
	} else {
		CorrectedWumpus.Happiness = UserWumpus.Happiness
	}

	if UserWumpus.Credits < 0 {
		CorrectedWumpus.Credits = 0
	} else {
		CorrectedWumpus.Credits = UserWumpus.Credits
	}

	CorrectedWumpus.Color = UserWumpus.Color
	CorrectedWumpus.Sick = UserWumpus.Sick
	CorrectedWumpus.Sleeping = UserWumpus.Sleeping
	CorrectedWumpus.Left = UserWumpus.Left
	return CorrectedWumpus
}

// LeftCheck checks if the Wumpus has left
func LeftCheck(UserWumpus Wumpus, session *discordgo.Session, event *discordgo.MessageCreate) (Left bool) {
	if UserWumpus.Age >= 14 {
		UserWumpus.Left = true
		UserWumpus.Age = 14
		UpdateWumpus(event.Author.ID, UserWumpus)
		sendMessage(session, event, event.ChannelID, UserWumpus.Name+" has something important to tell you!\nPlease run w.view")
		return true
	} else if UserWumpus.Age > 9 && UserWumpus.Left {
		sendMessage(session, event, event.ChannelID, UserWumpus.Name+" has something important to tell you!\nPlease run w.view")
		return true
	} else if UserWumpus.Age > 4 && UserWumpus.Age < 10 && UserWumpus.Left {
		sendMessage(session, event, event.ChannelID, UserWumpus.Name+" wants to talk\nPlease run w.view")
		return true
	} else if UserWumpus.Left {
		sendMessage(session, event, event.ChannelID, "You can't seem to find "+UserWumpus.Name+" anywhere ...\nPlease run w.view")
		return true
	}
	//Wumpus hasn't left yet :D
	return false
}

// WumpusCheck checks if the user has a Wumpus
func WumpusCheck(err error, session *discordgo.Session, event *discordgo.MessageCreate) (noPass bool) {
	if err != nil {
		sendMessage(session, event, event.ChannelID, "You need a Wumpus first!")
		return true
	}
	return false
}

// EnergyCheck checks if the Wumpus has enough energy
func EnergyCheck(UserWumpus Wumpus, requiredEnergy int, session *discordgo.Session, event *discordgo.MessageCreate) (noPass bool) {
	if UserWumpus.Energy < requiredEnergy {
		sendMessage(session, event, event.ChannelID, UserWumpus.Name+" is too tired!")
		return true
	}
	return false
}

// SleepCheck checks if the Wumpus is sleeping
// Can possibly wake up the Wumpus if conditions allow for it
func SleepCheck(UserWumpus Wumpus, canWake bool, session *discordgo.Session, event *discordgo.MessageCreate) (noPass bool) {
	if UserWumpus.Sleeping {
		if canWake && UserWumpus.Energy > 0 {
			UserWumpus.Sleeping = false
			UpdateWumpus(event.Author.ID, UserWumpus)
			sendMessage(session, event, event.ChannelID, UserWumpus.Name+" has woken from sleep!")
			return false
		}
		sendMessage(session, event, event.ChannelID, UserWumpus.Name+" is sleeping!")
		return true
	}
	return false
}

// CreditCheck checks if the user has enough credits
func CreditCheck(UserWumpus Wumpus, creditsRequired int, session *discordgo.Session, event *discordgo.MessageCreate) (noPass bool) {
	if UserWumpus.Credits < creditsRequired {
		sendMessage(session, event, event.ChannelID, "You need "+strconv.Itoa(creditsRequired)+"Ꞡ!")
		return true
	}
	return false
}
