package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/agilistikmal/socialview/app/events"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("DISCORD_TOKEN")
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer s.Close()

	s.Identify.Intents = discordgo.IntentsGuildMessages

	// Register Event Handler
	s.AddHandler(events.OnReady)
	s.AddHandler(events.FetchMessageUrl)

	err = s.Open()
	if err != nil {
		log.Fatal(err.Error())
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
