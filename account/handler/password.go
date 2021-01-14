package handler

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd string) string {
	bytePwd := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func comparePasswords(hashedPassword string, plainPassword string) bool {
	bytePassword := []byte(plainPassword)
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
