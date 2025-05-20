package auth

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint64) (string, error) {
	perm := jwt.MapClaims{}
	perm["authorized"] = true
	perm["exp"] = time.Now().Add(time.Hour * 6).Unix()
	perm["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perm)

	return token.SignedString(config.SecretKey)
}
