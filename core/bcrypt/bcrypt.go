package bcrypt

import (
	"cashier-api/core/logger"

	"golang.org/x/crypto/bcrypt"
)

// ComparePassword compare password
func ComparePassword(passwordHash, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return false
	}

	return true
}

// GeneratePassword generate password
func GeneratePassword(password string) (string, error) {
	passwordHashByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		logger.Logger.Errorf("[GeneratePassword] generate password error:%s", err)
		return "", err
	}

	return string(passwordHashByte), nil
}
