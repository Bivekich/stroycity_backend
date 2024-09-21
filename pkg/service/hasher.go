package service

import (
	"crypto/sha1"
	"fmt"
	"os"
)

var salt = os.Getenv("SALT")

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
