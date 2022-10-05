package middleware

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ed25519"
	"io"
)

func RequestValidation() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
func verifySignature(ctx *gin.Context) bool {
	bytesRequest, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logrus.Error(err)
		return false
	}
	defer func(ctx *gin.Context) {
		ctx.Request.Body = io.NopCloser(bytes.NewReader(bytesRequest))
	}(ctx)
	xSignatureEd25519 := ctx.Request.Header.Get("X-Signature-Ed25519")
	xSignatureTimestamp := ctx.Request.Header.Get("X-Signature-Timestamp")
	signature, err := hex.DecodeString(xSignatureEd25519)
	if err != nil {

	}
	fmt.Println(xSignatureTimestamp)
	if len(signature) != ed25519.SignatureSize || signature[63]&224 != 0 {

	}

	//return ed25519.Verify()
	return false
}
