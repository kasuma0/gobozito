package model

type DiscordCommand struct {
	ID                       string                  `json:"id"`
	Type                     int                     `json:"type"`
	ApplicationID            string                  `json:"application_id"`
	GuildID                  string                  `json:"guild_id"`
	Name                     string                  `json:"name"`
	NameLocalizations        []interface{}           `json:"name_localizations"`
	Description              string                  `json:"description"`
	DescriptionLocalizations []interface{}           `json:"description_localizations"`
	Options                  []DiscordCommandOptions `json:"options"`
	DefaultMemberPermissions string                  `json:"default_member_permissions"`
	DMPermission             bool                    `json:"dm_permission"`
	DefaultPermission        bool                    `json:"defaultPermission"`
	Version                  string                  `json:"version"`
}
type DiscordCommandOptions struct {
	Name                     string                         `json:"name"`
	NameLocalizations        []interface{}                  `json:"name_localizations"`
	Description              string                         `json:"description"`
	DescriptionLocalizations []interface{}                  `json:"description_localizations"`
	Type                     int                            `json:"type"`
	Required                 bool                           `json:"required"`
	Choices                  []DiscordCommandOptionsChoices `json:"choices,omitempty"`
	Options                  []DiscordCommandOptions        `json:"options"`
	ChannelTypes             []int
}
type DiscordCommandOptionsChoices struct {
	Name              string        `json:"name"`
	NameLocalizations []interface{} `json:"name_localizations"`
	Value             string        `json:"value"`
}
