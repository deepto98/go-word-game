package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/crypto/scrypt"
)

func HashPassword(password string) (string, error) {
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

//Check if supplied password, when salted with same salt, produces the same value
func ComparePasswords(storedPassword string, suppliedPassword string) (bool, error) {
	passwordSalt := strings.Split(storedPassword, ".")

	//Check supplied password, salted with hash
	salt, err := hex.DecodeString(passwordSalt[1])

	if err != nil {
		return false, fmt.Errorf("Unable to verify user password")
	}

	shash, err := scrypt.Key([]byte(suppliedPassword), salt, 32768, 8, 1, 32)

	match := hex.EncodeToString(shash) == passwordSalt[0]
	return match, nil
}
