package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/agilistikmal/socialview/internal/infrastructure/config"
	delivery "github.com/agilistikmal/socialview/internal/socialview/delivery/discord"
	"github.com/agilistikmal/socialview/internal/socialview/delivery/discord/event"
	"github.com/agilistikmal/socialview/internal/socialview/service"
)

func main() {
	config.Load()

	tiktokService := service.NewTikTokService()
	mediaService := service.NewMediaService(tiktokService)
	discordEventHandler := event.NewDiscordEventHandler(mediaService)
	discordDelivery := delivery.NewDiscordDelivery(discordEventHandler)
	discordDelivery.RegisterEvents()

	err := discordDelivery.Session.Open()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Airhorn is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discordDelivery.Session.Close()
}
