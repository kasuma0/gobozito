package model

type PingResponse struct {
	Type int64 `json:"type"`
}
type InteractionRequest struct {
	ApplicationId string `json:"application_id"`
	Id            string `json:"id"`
	Token         string `json:"token"`
	Type          int    `json:"type"`
	User          struct {
		Avatar           string      `json:"avatar"`
		AvatarDecoration interface{} `json:"avatar_decoration"`
		Discriminator    string      `json:"discriminator"`
		Id               string      `json:"id"`
		PublicFlags      int         `json:"public_flags"`
		Username         string      `json:"username"`
	} `json:"user"`
	Version int `json:"version"`
}
