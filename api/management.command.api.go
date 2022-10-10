package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kasuma0/gobozito/controller"
	"github.com/kasuma0/gobozito/model"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

func CreateAndUpdateCommandHandler(ctx *gin.Context) {
	var request model.DiscordCommand
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("bad-request: %s", err.Error())})
		return
	}
	response, err := controller.CreateAndUpdateCommandController(ctx, &request)
	if err != nil {
		statusCode := http.StatusInternalServerError
		var errResp interface{} = err
		if strings.Contains(err.Error(), "statusCode") {
			statusCode, err = strconv.Atoi(strings.TrimLeft(err.Error(), "statusCode: "))
			if err != nil {
				logrus.Error(err)
			} else {
				errResp = response
			}
		}
		ctx.AbortWithStatusJSON(statusCode, errResp)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func GetCommands(ctx *gin.Context) {
	response, err := controller.GETCommandsController(ctx)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func DeleteCommandHandler(ctx *gin.Context) {
	var request model.DeleteCommandRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("bad-request: %s", err.Error())})
		return
	}

	response, err := controller.DeleteCommandController(ctx, &request)
	if err != nil {
		statusCode := http.StatusInternalServerError
		var errResp interface{} = err
		if strings.Contains(err.Error(), "statusCode") {
			statusCode, err = strconv.Atoi(strings.TrimLeft(err.Error(), "statusCode: "))
			if err != nil {
				logrus.Error(err)
			} else {
				errResp = response
			}
		}
		ctx.AbortWithStatusJSON(statusCode, errResp)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
