package accounts

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func comparePasswords(hashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
