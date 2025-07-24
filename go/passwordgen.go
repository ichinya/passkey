package main

import (
	"crypto/rand"
	"errors"
	"math/big"
)

var charsets = map[string]string{
	"low":      "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	"medium":   "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?",
	"strong":   "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/\\`~",
	"paranoid": "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/\\`~☺☻♥♦♣♠•◘○◙♂♀♪♫☼►◄↕‼¶§▬↨↑↓→←∟↔▲▼",
}

func GeneratePassword(length int, level string) (string, error) {
	charset, ok := charsets[level]
	if !ok {
		return "", errors.New("invalid complexity level")
	}

	// Преобразуем в срез рун
	runes := []rune(charset)
	password := make([]rune, length)
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(runes))))
		if err != nil {
			return "", err
		}
		password[i] = runes[num.Int64()]
	}

	return string(password), nil
}
