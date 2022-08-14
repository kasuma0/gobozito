package conf

import "os"

//Configuration add constant needed
type Configuration struct {
	DiscordToken string
	BotID        string
}

var DiscordConfiguration Configuration

func init() {
	DiscordConfiguration = Configuration{
		DiscordToken: os.Getenv("TOKEN"),
		BotID:        os.Getenv("BOTID"),
	}
}
