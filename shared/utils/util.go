package utils

import (
	"crypto"
	"encoding/hex"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretJwt []byte

func EncryptSession(id string, expired int) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		return "", errors.New("jwt secret cannot be empty")
	}

	secretJwt = []byte(secret)

	encrypt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Second * time.Duration(expired)).Unix(),
	})

	token, err := encrypt.SignedString(secretJwt)

	if err != nil {
		return "", err
	}

	return token, nil
}

func Sha256(text string) string {
	algorithm := crypto.SHA256.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
