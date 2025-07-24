package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"strings"
)

func Encrypt(plaintext, key, mode string) (string, error) {
	if mode != "shell" && mode != "safe" && mode != "" {
		return "", errors.New("invalid encryption mode")
	}
	salt := make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}
	keyHash := sha256.Sum256([]byte(key + string(salt)))
	block, err := aes.NewCipher(keyHash[:])
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)
	final := append(salt, append(nonce, ciphertext...)...)

	if mode == "shell" {
		return encodeShellSafe(final), nil
	}
	return base64.StdEncoding.EncodeToString(final), nil
}

func Decrypt(ciphertext, key, mode string) (string, error) {
	if mode != "shell" && mode != "safe" && mode != "" {
		return "", errors.New("invalid decryption mode")
	}
	var raw []byte
	var err error
	if mode == "shell" {
		raw, err = decodeShellSafe(ciphertext)
	} else {
		raw, err = base64.StdEncoding.DecodeString(ciphertext)
	}
	if err != nil {
		return "", err
	}
	if len(raw) < 8 {
		return "", errors.New("invalid ciphertext")
	}
	salt := raw[:8]
	keyHash := sha256.Sum256([]byte(key + string(salt)))
	block, err := aes.NewCipher(keyHash[:])
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	if len(raw) < 8+gcm.NonceSize() {
		return "", errors.New("invalid ciphertext")
	}
	nonce := raw[8 : 8+gcm.NonceSize()]
	ciphertextData := raw[8+gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertextData, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func encodeShellSafe(b []byte) string {
	s := base64.StdEncoding.EncodeToString(b)
	s = strings.ReplaceAll(s, "+", "-")
	s = strings.ReplaceAll(s, "/", "_")
	s = strings.ReplaceAll(s, "=", "")
	return s
}

func decodeShellSafe(s string) ([]byte, error) {
	s = strings.ReplaceAll(s, "-", "+")
	s = strings.ReplaceAll(s, "_", "/")
	switch len(s) % 4 {
	case 2:
		s += "=="
	case 3:
		s += "="
	}
	return base64.StdEncoding.DecodeString(s)
}
