package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/kasuma0/gobozito/conf"
	"github.com/kasuma0/gobozito/handler"
	"github.com/kasuma0/gobozito/routes"
)

func main() {
	discord, err := discordgo.New("Bot " + conf.DiscordConfiguration.DiscordToken)
	if err != nil {
		log.Fatal(err)
	}
	discord.AddHandler(handler.DiscordHandler)
	discord.Open()
	defer discord.Close()
	gin.ForceConsoleColor()
	engine := gin.Default()
	routes.Routes(engine)
	engine.Run(":8080")
}
