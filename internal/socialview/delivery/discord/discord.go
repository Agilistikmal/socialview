package delivery

import (
	"log"

	"github.com/agilistikmal/socialview/internal/socialview/delivery/discord/event"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

type DiscordDelivery struct {
	Session             *discordgo.Session
	DiscordEventHandler *event.DiscordEventHandler
}

func NewDiscordDelivery(discordEventHandler *event.DiscordEventHandler) *DiscordDelivery {
	session, err := discordgo.New("Bot " + viper.GetString("discord.token"))
	if err != nil {
		log.Fatal(err)
	}

	session.Identify.Intents = discordgo.IntentGuilds |
		discordgo.IntentMessageContent |
		discordgo.IntentGuildMessages | discordgo.IntentsAll

	return &DiscordDelivery{
		Session:             session,
		DiscordEventHandler: discordEventHandler,
	}
}

func (d *DiscordDelivery) RegisterEvents() {
	d.Session.AddHandler(d.DiscordEventHandler.Ready)
	d.Session.AddHandler(d.DiscordEventHandler.WatchURLs)
}
