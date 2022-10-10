package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kasuma0/gobozito/middleware"
	"github.com/kasuma0/gobozito/routes"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {
	signalChannel := make(chan os.Signal, 2)
	go startDiscordEngine(signalChannel)
	go startManagementEngine(signalChannel)
	<-signalChannel
}

func startDiscordEngine(sgnChan chan os.Signal) {
	gin.ForceConsoleColor()
	engine := gin.Default()
	engine.Use(middleware.RequestValidation())
	routes.DiscordRoutes(engine)
	logrus.Fatal(engine.Run("0.0.0.0:8080"))
	signal.Notify(sgnChan, os.Interrupt, os.Kill)
}
func startManagementEngine(sgnChan chan os.Signal) {
	gin.ForceConsoleColor()
	engine := gin.Default()
	routes.ManagementRoutes(engine)
	logrus.Fatal(engine.Run("0.0.0.0:8081"))
	signal.Notify(sgnChan, os.Interrupt, os.Kill)
}
