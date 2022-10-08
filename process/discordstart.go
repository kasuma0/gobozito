package process

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/kasuma0/gobozito/conf"
	"github.com/kasuma0/gobozito/model"
	"github.com/kasuma0/gobozito/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	commandURL = conf.DiscordConfiguration.DiscordApiURL + fmt.Sprintf("/applications/%s/commands", conf.DiscordConfiguration.BotID)
)

// DiscordStart start discord command validation.
// if commando is not detected, creates command
func DiscordStart() {
	ctx := context.Background()
	if !getCommands(ctx, "gobozito") {
		err := createCommandGobozito(ctx)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

func getCommands(ctx context.Context, command string) bool {
	request, err := http.NewRequest(http.MethodGet, commandURL, nil)
	request.Header.Set("Authorization", "Bot "+conf.DiscordConfiguration.DiscordToken)
	if err != nil {
		logrus.Fatal(err)
	}
	byteResponse, statusCode, err := util.HttpCustomClient(ctx, request)
	if err != nil {
		logrus.Fatal(err)
	}
	if statusCode != http.StatusOK {
		logrus.Fatalf("response no OK: %s", string(byteResponse))
	}
	var commands []model.DiscordCommand
	err = json.Unmarshal(byteResponse, &commands)
	if err != nil {
		logrus.Fatal(err)
	}
	for _, discordCommand := range commands {
		if discordCommand.Name == command {
			return true
		}
	}
	return false
}

func createCommandGobozito(ctx context.Context) error {
	commandRequest := model.DiscordCommand{
		Name:        "gobozito",
		Type:        1,
		Description: "command gobozito",
		Options: []model.DiscordCommandOptions{{
			Name:        "birthday",
			Description: "birthday option, saves your birthday",
			Type:        1,
			Required:    false,
			Choices: []model.DiscordCommandOptionsChoices{{
				Name:  "add",
				Value: "add_birthday",
			}},
		}},
	}
	return CreateAndUpdateCommand(ctx, http.MethodPost, &commandRequest)
}

// CreateAndUpdateCommand create discord command or update discord command
//
//	urlMethod: accept method POST(create command) Patch(update command)
func CreateAndUpdateCommand(ctx context.Context, urlMethod string, command *model.DiscordCommand) error {
	bytesRequest, err := json.Marshal(command)
	if err != nil {
		logrus.Error(err)
		return err
	}
	commandRequest, err := http.NewRequest(urlMethod, commandURL, bytes.NewReader(bytesRequest))
	if err != nil {
		logrus.Error(err)
		return err
	}
	commandRequest.Header.Set("Authorization", "Bot "+conf.DiscordConfiguration.DiscordToken)
	commandRequest.Header.Set("Content-Type", "application/json")
	bytesCommandResp, statusCode, err := util.HttpCustomClient(ctx, commandRequest)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if statusCode != http.StatusOK && statusCode != http.StatusCreated {
		logrus.Errorf("response no OK: %s", string(bytesCommandResp))
		return err
	}
	var commandResponse interface{}
	err = json.Unmarshal(bytesCommandResp, &commandResponse)
	if err != nil {
		logrus.Error(err)
		return err
	}
	fmt.Println(commandResponse)
	return nil
}
