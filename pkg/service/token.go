package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var signingKey = os.Getenv("SIGN_KEY_STRING")

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
	Role   string `json:"role"`
}

func CreateToken(userId, role string) (token string) {
	params := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
		Role:   role,
	})

	token, _ = params.SignedString([]byte(signingKey))
	return token
}

func ParseToken(accessToken string) (string, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", "", errors.New("invalid token claims")
	}
	return claims.UserId, claims.Role, nil
}
