package conf

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Configuration add constant needed
type Configuration struct {
	DiscordToken     string
	BotID            string
	DiscordApiURL    string
	DiscordPublicKey string
	DiscordEpoch     int
	Credentials      struct {
		MongoDB struct {
			URL             string
			MaxPoolSize     uint64
			MinPoolSize     uint64
			MaxConnIdleTime time.Duration
		}
	}
}

var DiscordConfiguration Configuration

func init() {
	maxPoolSize, err := strconv.Atoi(os.Getenv("MONGOMAXPOOLSIZE"))
	if err != nil {
		log.Fatal(err)
	}
	minPoolSize, err := strconv.Atoi(os.Getenv("MONGOMINPOOLSIZE"))
	if err != nil {
		log.Fatal(err)
	}
	maxConnIdleTime, err := strconv.Atoi(os.Getenv("MONGOMAXIDLETIME"))
	if err != nil {
		log.Fatal(err)
	}
	DiscordConfiguration = Configuration{
		DiscordToken:     os.Getenv("TOKEN"),
		BotID:            os.Getenv("BOTID"),
		DiscordApiURL:    fmt.Sprintf("https://discord.com/api/v%s", os.Getenv("DISCORDVERSION")),
		DiscordPublicKey: os.Getenv("DISCORDPUBLICKEY"),
		DiscordEpoch:     1420070400000,
		Credentials: struct {
			MongoDB struct {
				URL             string
				MaxPoolSize     uint64
				MinPoolSize     uint64
				MaxConnIdleTime time.Duration
			}
		}{
			MongoDB: struct {
				URL             string
				MaxPoolSize     uint64
				MinPoolSize     uint64
				MaxConnIdleTime time.Duration
			}{
				URL:             os.Getenv("MONGODBURL"),
				MaxPoolSize:     uint64(maxPoolSize),
				MinPoolSize:     uint64(minPoolSize),
				MaxConnIdleTime: time.Duration(maxConnIdleTime),
			},
		},
	}
}
