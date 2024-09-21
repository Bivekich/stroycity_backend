package service

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var signingKey = os.Getenv("SIGN_KEY_STRING")

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func CreateToken(userId string) (token string) {
	params := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
	})

	token, _ = params.SignedString([]byte(signingKey))
	return token
}
