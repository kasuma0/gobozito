package model

type UserLoginRequest struct {
	User     string `form:"user"`
	Password string `form:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
