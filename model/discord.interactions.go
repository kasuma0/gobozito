package model

type DiscordCommand struct {
	Name        string                  `json:"name"`
	Type        int                     `json:"type"`
	Description string                  `json:"description"`
	Options     []DiscordCommandOptions `json:"options"`
}

type DiscordCommandOptions struct {
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	Type        int                            `json:"type"`
	Required    bool                           `json:"required"`
	Choices     []DiscordCommandOptionsChoices `json:"choices,omitempty"`
}

type DiscordCommandOptionsChoices struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
