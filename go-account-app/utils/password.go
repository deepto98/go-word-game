package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func hashPassword(password string) (string, error) {
	//-------------------
	// Salting password
	//-------------------

	//Create salt
	salt := make([]byte, 32)

	//Add random values to salt
	_, err := rand.Read(salt)

	if err != nil {
		return "", err
	}

	//Creating hash with Cost Parameters
	shash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)

	if err != nil {
		return "", err
	}

	//Return hex-encoded string with salt appended to password
	hashedPassword := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))
	return hashedPassword, nil
}
