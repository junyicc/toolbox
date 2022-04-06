package crypto

import (
	"fmt"
	"toolbox/typeconv"

	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword generates "hash" password
func EncryptPassword(input string) (string, error) {
	if input == "" {
		return "", ErrInvalidData
	}
	// generate "hash" password
	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("bcrypt.GenerateFromPassword error: %w", err)
	}
	// encrypted password
	return typeconv.BytesToString(hash), nil
}

// ComparePassword compares input password and hash password
func ComparePassword(inputPwd, hashPwd string) error {
	return bcrypt.CompareHashAndPassword(typeconv.StringToBytes(hashPwd), typeconv.StringToBytes(inputPwd))
}
