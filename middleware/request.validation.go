package middleware

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kasuma0/gobozito/conf"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ed25519"
	"io"
	"net/http"
)

func RequestValidation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !verifySignature(ctx) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized"})
		}
		ctx.Next()
	}
}
func verifySignature(ctx *gin.Context) bool {
	bytesRequest, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logrus.Error(err)
		return false
	}
	discordPublicKey, err := hex.DecodeString(conf.DiscordConfiguration.DiscordPublicKey)
	if err != nil {
		logrus.Error(err)
		return false
	}
	publicKey := ed25519.PublicKey(discordPublicKey)
	defer func(ctx *gin.Context) {
		ctx.Request.Body = io.NopCloser(bytes.NewReader(bytesRequest))
	}(ctx)
	xSignatureEd25519 := ctx.Request.Header.Get("X-Signature-Ed25519")
	xSignatureTimestamp := ctx.Request.Header.Get("X-Signature-Timestamp")
	signature, err := hex.DecodeString(xSignatureEd25519)
	if err != nil {
		logrus.Error(err)
		return false
	}
	fmt.Println(xSignatureTimestamp)
	if len(signature) != ed25519.SignatureSize || signature[63]&224 != 0 {
		return false
	}
	var msg bytes.Buffer
	msg.WriteString(xSignatureTimestamp)
	_, err = io.Copy(&msg, bytes.NewReader(bytesRequest))
	if err != nil {
		logrus.Error(err)
		return false
	}
	return ed25519.Verify(publicKey, msg.Bytes(), signature)
}
