package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// Return a hashed password
func HashPass(password string) (string, error) {
	key, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(key), nil
}

// MatchPass returns if the hash matches pass or not
func MatchPass(hash, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err != nil
}
