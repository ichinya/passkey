package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	LevelLow      = "low"
	LevelMedium   = "medium"
	LevelStrong   = "strong"
	LevelParanoid = "paranoid"
)

func GeneratePassword(length int, level string) (string, error) {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	basic := "@#-_+="
	advanced := "!@#$%^&*()[]{}<>?~"
	emoji := "🔐💡🧠🌐🎲🚀你好†±λ@#€"

	chars := letters + numbers
	switch level {
	case LevelLow:
		chars = letters + numbers
	case LevelMedium:
		chars = letters + numbers + basic
	case LevelStrong:
		chars = letters + numbers + basic + advanced
	case LevelParanoid:
		chars = letters + numbers + basic + advanced + emoji
	default:
		return "", fmt.Errorf("unknown level: %s", level)
	}

	var result strings.Builder
	runes := []rune(chars)
	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(runes))))
		if err != nil {
			return "", err
		}
		result.WriteRune(runes[idx.Int64()])
	}
	return result.String(), nil
}
