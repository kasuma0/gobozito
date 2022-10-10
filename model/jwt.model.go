package model

import "github.com/golang-jwt/jwt/v4"

type JWTManagementDiscordCommand struct {
	User string `json:"user"`
	jwt.RegisteredClaims
}
