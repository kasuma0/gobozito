package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kasuma0/gobozito/controller"
	"github.com/kasuma0/gobozito/model"
	"net/http"
)

func DiscordInteractions(ctx *gin.Context) {
	var request model.InteractionRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := controller.PingController(&request)
	if response != nil {
		ctx.JSON(http.StatusOK, response)

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "no ping"})
	}

}
