package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

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

/*

// testing
func correctPassTest(t *testing.T) {
	// should return true as correct pass is given
	pass := "Password1"
	hash, err := hashedPass(pass)
	if err != nil {
		t.Fatalf("Generating hashed password failed", err)
	}

	checking := compareHash(pass, hash)
	fmt.Println("Passwords Match: ", checking)
}

func incorrectPassTest(t *testing.T) {
	// this should return false as an incorrect pass is given
	pass := "Password1"
	hash, err := hashedPass(pass)
	if err != nil {
		t.Fatalf("Generating hashed password failed", err)
	}

	pass2 := "Password10"
	checking := compareHash(pass2, hash)
	fmt.Println("Passwords Match: ", checking)
}
*/
