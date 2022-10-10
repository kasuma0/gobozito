package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kasuma0/gobozito/controller"
	"github.com/kasuma0/gobozito/model"
	"net/http"
)

func LoginPrivateHandler(ctx *gin.Context) {
	var request model.UserLoginRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("bad-request: %s", err.Error())})
		return
	}
	response, err := controller.LoginPrivateController(ctx, &request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
