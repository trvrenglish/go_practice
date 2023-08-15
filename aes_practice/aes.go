package main

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/eiannone/keyboard"
)

func main() {
	key := GenerateRandomKey(32)
	plaintext := ScanInput()
	// Encrypt the plaintext
	ciphertext := EncryptAES([]byte(key), plaintext)
	DecryptAES([]byte(key), ciphertext)
}

func ScanInput() string {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	var plaintext string
	fmt.Print("Enter 16 character plaintext (characters used: 0): ")

	charCount := 0
	for len(plaintext) != 16 {
		char, key, err := keyboard.GetKey()

		if key == keyboard.KeySpace {
			fmt.Println("hure")
		}
		if err != nil {
			panic(err)
		}

		// Don't forget to count space key
		if key == keyboard.KeySpace {
			char = ' '
		}

		if char != 0 && len(plaintext) < 16 {
			plaintext += string(char)
			charCount++
			fmt.Printf("\rEnter 16 character plaintext (characters used: %d): %s", charCount, plaintext)
		}
	}
	fmt.Println()
	return plaintext
}

func GenerateRandomKey(size int) []byte {
	key := make([]byte, size)
	_, err := rand.Read(key)
	CheckError(err)
	return key
}

func EncryptAES(key []byte, plaintext string) string {
	// Create the cipher
	cipher, err := aes.NewCipher(key)
	CheckError(err)

	allocByteSlice := make([]byte, len(plaintext))

	cipher.Encrypt(allocByteSlice, []byte(plaintext))
	ciphertext := hex.EncodeToString(allocByteSlice)
	fmt.Println("ENCRYPTED:", ciphertext)

	return ciphertext
}

func DecryptAES(key []byte, ciphertext string) {
	decodedCipher, _ := hex.DecodeString(ciphertext)

	c, err := aes.NewCipher(key)
	CheckError(err)

	plaintext := make([]byte, len(ciphertext))
	c.Decrypt(plaintext, decodedCipher)

	plaintextAsString := string(plaintext[:])
	fmt.Println("DECRYPTED:", plaintextAsString)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
