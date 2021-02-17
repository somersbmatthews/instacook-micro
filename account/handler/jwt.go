package handler

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func NewJWT(userID string) (string, error) {
	signingKey, ok := os.LookupEnv("JWTSIGNINGKEY")
	if ok != true {
		log.Fatalln("JWTSIGNINGKEY not found")
	}

	claims := jwt.MapClaims{
		"UserID": userID,
		"StandardClaims": jwt.StandardClaims{
			ExpiresAt: 86400000,
			Issuer:    "Authentication",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		log.Fatalln(err)
	}
	return ss, nil
}
