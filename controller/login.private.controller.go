package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kasuma0/gobozito/model"
	"github.com/kasuma0/gobozito/security"
)

func LoginPrivateController(ctx *gin.Context, request *model.UserLoginRequest) (interface{}, error) {
	token, err := security.GenerateJWT(request.User)
	if err != nil {
		return nil, err
	}
	resp := model.UserLoginResponse{Token: token}
	return resp, nil
}
