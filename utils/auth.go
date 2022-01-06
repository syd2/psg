package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenAuthToken(userid string, secret_token string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userid"] = userid
	claims["ExpiresAt"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret_token))
}
