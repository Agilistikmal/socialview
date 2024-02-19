package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnReady(s *discordgo.Session, ready *discordgo.Ready) {
	log.Println("Bot is online")
	s.UpdateWatchStatus(0, "Social Media")
}
