package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kasuma0/gobozito/conf"
	"github.com/kasuma0/gobozito/model"
	"github.com/kasuma0/gobozito/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	commandURL = conf.DiscordConfiguration.DiscordApiURL + fmt.Sprintf("/applications/%s/commands", conf.DiscordConfiguration.BotID)
)

// CreateAndUpdateCommandController create discord command or update discord command
//
//	urlMethod: accept method POST(create command) Patch(update command)
func CreateAndUpdateCommandController(ctx *gin.Context, command *model.DiscordCommand) (interface{}, error) {
	url := commandURL
	if ctx.Request.Method == http.MethodPatch {
		url = url + "/" + command.ID
		command.ID = ""
	}
	bytesRequest, err := json.Marshal(command)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	commandRequest, err := http.NewRequest(ctx.Request.Method, url, bytes.NewReader(bytesRequest))
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	commandRequest.Header.Set("Authorization", "Bot "+conf.DiscordConfiguration.DiscordToken)
	commandRequest.Header.Set("Content-Type", "application/json")
	bytesCommandResp, statusCode, err := util.HttpCustomClient(ctx.Request.Context(), commandRequest)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if statusCode != http.StatusOK && statusCode != http.StatusCreated {
		logrus.Errorf("response no OK: %s", string(bytesCommandResp))
		var resp interface{}
		if err = json.Unmarshal(bytesCommandResp, &resp); err != nil {
			logrus.Error(err)
			return nil, err
		}
		return &resp, errors.New(fmt.Sprintf("statusCode: %d", statusCode))
	}
	var commandResponse interface{}
	err = json.Unmarshal(bytesCommandResp, &commandResponse)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &commandResponse, nil
}

func GETCommandsController(ctx *gin.Context) (commands []model.DiscordCommand, err error) {
	request, err := http.NewRequest(http.MethodGet, commandURL, nil)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	request.Header.Set("Authorization", "Bot "+conf.DiscordConfiguration.DiscordToken)
	byteResponse, statusCode, err := util.HttpCustomClient(ctx.Request.Context(), request)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if statusCode != http.StatusOK {
		err = errors.New("response no OK: " + string(byteResponse))
		logrus.Error(err)
		return nil, err
	}
	err = json.Unmarshal(byteResponse, &commands)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return commands, nil
}
func DeleteCommandController(ctx *gin.Context, request *model.DeleteCommandRequest) (interface{}, error) {
	url := commandURL + "/" + request.CommandID
	newRequest, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	newRequest.Header.Set("Authorization", "Bot "+conf.DiscordConfiguration.DiscordToken)
	byteResponse, statusCode, err := util.HttpCustomClient(ctx.Request.Context(), newRequest)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if statusCode != http.StatusOK && statusCode != http.StatusNoContent {
		var resp interface{}
		err = json.Unmarshal(byteResponse, &resp)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		err = errors.New(fmt.Sprintf("statusCode: %d", statusCode))
		logrus.Errorf("%s", string(byteResponse))
		return &resp, err
	}
	var response interface{}
	err = json.Unmarshal(byteResponse, &response)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &response, nil
}
