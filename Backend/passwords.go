package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func getPass() {
	// get password from user
	// hashed password will be stored in the database
}

func hashedPass(password string) (string, error) {
	// generates a new hashed password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err
}

func compareHash(password string, hash string) bool {
	// compare the real pass with the hashed one
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
