package event

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h *DiscordEventHandler) Ready(s *discordgo.Session, _ *discordgo.Ready) {
	s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: "Social Media",
				Type: discordgo.ActivityTypeWatching,
				URL:  "https://agil.zip",
			},
		},
	})
	log.Println("Discord bot online as " + s.State.User.Username)
}
