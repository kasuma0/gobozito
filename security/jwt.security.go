package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kasuma0/gobozito/conf"
	"github.com/kasuma0/gobozito/model"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/scrypt"
	"io"
	"time"
)

func GenerateJWT(user string) (string, error) {
	iss, err := Encrypt(conf.DiscordConfiguration.BotID)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	userEnc, err := Encrypt(user)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	claims := model.JWTManagementDiscordCommand{
		User: userEnc,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			Issuer:    iss,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	strToken, err := token.SignedString([]byte(conf.DiscordConfiguration.JWTSecret))
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return strToken, nil
}

func Encrypt(str string) (string, error) {
	key, err := scrypt.Key([]byte(conf.DiscordConfiguration.JWTSecret), nil, 32768, 8, 1, 32)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	text := []byte(str)
	block, err := aes.NewCipher(key)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	if len(text)%aes.BlockSize != 0 {
		text, err = pkcs7padding(text, aes.BlockSize)
		if err != nil {
			logrus.Error(err)
			return "", err
		}
	}
	cipherText := make([]byte, aes.BlockSize+len(text))
	iv := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], text)
	return hex.EncodeToString(cipherText), nil
}
func Decrytp(encryptedText string) (string, error) {
	key, err := scrypt.Key([]byte(conf.DiscordConfiguration.JWTSecret), nil, 32768, 8, 1, 32)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	cipherText, err := hex.DecodeString(encryptedText)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		err = errors.New("is too short")
		logrus.Error(err)
		return "", err
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		err = errors.New("ciphertext is not multiple of blockSize")
		logrus.Error(err)
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)
	return string(cipherText), nil
}
func pkcs7padding(padtext []byte, blocksize int) ([]byte, error) {
	if blocksize < 0 || blocksize > 256 {
		return nil, fmt.Errorf("pkcs7: Invalid block size %d", blocksize)
	}
	padLen := blocksize - len(padtext)%blocksize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(padtext, padding...), nil
}
