package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kasuma0/gobozito/api"
)

func DiscordRoutes(engine *gin.Engine) {
	interactions := engine.Group("/interactions")
	{
		interactions.POST("/discord", api.DiscordInteractions)
	}
}

func ManagementRoutes(engine *gin.Engine) {
	management := engine.Group("/management")
	{
		command := management.Group("/discord/command")
		{
			command.POST("", api.CreateAndUpdateCommandHandler)
			command.GET("", api.GetCommands)
			command.DELETE("", api.DeleteComandHandler)
			command.PATCH("", api.CreateAndUpdateCommandHandler)
		}
		login := management.Group("/login")
		{
			login.GET("/private", api.LoginPrivateHandler)
		}
	}

}
