package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/kasuma0/gobozito/conf"
	"github.com/kasuma0/gobozito/handler"
)

func main() {
	discord, err := discordgo.New("Bot " + conf.DiscordConfiguration.DiscordToken)
	if err != nil {
		log.Fatal(err)
	}
	discord.AddHandler(handler.DiscordHandler)
	discord.Open()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	discord.Close()
}
