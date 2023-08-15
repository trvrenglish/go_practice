package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Set a password
	password := "VALLEYFORGE"

	// Hash the password
	hashedPassword, _ := Hash(password)
	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hashedPassword)

	// Verify the password
	var decryptState string

	if Verify(password, hashedPassword) {
		decryptState = "decrypts"
	} else {
		decryptState = "does not decrypt"
	}
	fmt.Printf("The hash %s to the correct password.\n", decryptState)

}

func Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}

func Verify(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil

}
