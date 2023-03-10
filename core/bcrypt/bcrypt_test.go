package bcrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComparePassword_Correct(t *testing.T) {
	var result bool
	firstPassword := "$2a$10$DLW5G3CwkeBDcX6cl4bYGORwHqKDaYfbTaAHTVh6nnxO8fKu4Epqm"
	secondPassword := "00000000"
	check := ComparePassword(firstPassword, secondPassword)
	if check {
		result = true
	}

	assert.Equal(t, true, result)
}

func TestComparePassword_Wrong(t *testing.T) {
	var result bool
	firstPassword := "$2a$10$DLW5G3CwkeBDcX6cl4bYGORwHqKDaYfbTaAHTVh6nnxO8fKu4Epqm"
	secondPassword := "P@ssw0rd"
	check := ComparePassword(firstPassword, secondPassword)
	if !check {
		result = true
	}

	assert.Equal(t, true, result)
}

func TestGeneratePassword(t *testing.T) {
	var result bool
	password := "00000000"
	hashPW, _ := GeneratePassword(password)
	if hashPW != "" {
		result = true
	}

	assert.Equal(t, true, result)
}
