package util

import (
	"github.com/kasuma0/gobozito/conf"
	"time"
)

func MakeSnowflake() int {
	return (int(time.Now().UnixMilli()) - conf.DiscordConfiguration.DiscordEpoch) << 22

}
