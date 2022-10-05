package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kasuma0/gobozito/middleware"
	"github.com/kasuma0/gobozito/process"
	"github.com/kasuma0/gobozito/routes"
	"github.com/sirupsen/logrus"
)

func main() {
	process.DiscordStart()
	gin.ForceConsoleColor()
	engine := gin.Default()
	engine.Use(middleware.RequestValidation())
	routes.Routes(engine)
	logrus.Fatal(engine.Run("0.0.0.0:8080"))

}
