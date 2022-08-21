package handler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/kasuma0/gobozito/controller"
	"github.com/kasuma0/gobozito/model"
	"github.com/sirupsen/logrus"
)

func DiscordHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	command := strings.Split(m.Content, " ")
	if command[0] != "g=" {
		return
	}
	switch command[1] {
	case "birthday":
		switch command[2] {
		case "add":
			fmt.Println(m.ID)
			birthdate, err := time.Parse("2006/01/02", command[3])
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "ingrese fecha en formato 2006/01/02")
				return
			}
			birthObj := model.BirthayADD{
				UserID:    m.Author.ID,
				Name:      m.Author.Username,
				Birthdate: int(birthdate.Unix()),
				Birthday:  birthdate.Format("01/02"),
			}

			err = controller.BirthdayADD(context.Background(), birthObj)
			if err != nil {
				logrus.Error(err)
			}

		}
	default:
		return

	}
}
