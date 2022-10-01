package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kasuma0/gobozito/handler"
)

func Routes(engine *gin.Engine) {
	engine.POST("/", handler.DiscordInteractions)
}
