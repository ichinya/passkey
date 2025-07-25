package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(plaintext, key, mode string) (string, error) {
	switch mode {
	case "shell":
		return encryptShell(plaintext, key)
	case "safe", "":
		// PBKDF2 вместо SHA256
		salt := make([]byte, 8)
		if _, err := io.ReadFull(rand.Reader, salt); err != nil {
			return "", err
		}
		dk := pbkdf2.Key([]byte(key), salt, 100000, 32, sha256.New)
		block, err := aes.NewCipher(dk)
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
		return encodeShellSafe(final), nil
	default:
		return "", errors.New("invalid encryption mode")
	}
}

func Decrypt(ciphertext, key, mode string) (string, error) {
	switch mode {
	case "shell":
		return decryptShell(ciphertext, key)
	case "safe", "":
		// PBKDF2 вместо SHA256
		var raw []byte
		var err error
		raw, err = decodeShellSafe(ciphertext)
		if err != nil {
			return "", err
		}
		if len(raw) < 8 {
			return "", errors.New("invalid ciphertext")
		}
		salt := raw[:8]
		dk := pbkdf2.Key([]byte(key), salt, 100000, 32, sha256.New)
		block, err := aes.NewCipher(dk)
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
	default:
		return "", errors.New("invalid decryption mode")
	}
}

func encodeShellSafe(b []byte) string {
	s := base64.StdEncoding.EncodeToString(b)
	s = strings.ReplaceAll(s, "+", "-")
	s = strings.ReplaceAll(s, "/", "_")
	s = strings.ReplaceAll(s, "=", "")
	return s
}

func decodeShellSafe(s string) ([]byte, error) {
	if strings.ContainsAny(s, "+/=") {
		return nil, errors.New("invalid safe encoded data")
	}
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

func pkcs7Pad(data []byte, blockSize int) []byte {
	pad := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(pad)}, pad)
	return append(data, padding...)
}

func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 || len(data)%blockSize != 0 {
		return nil, errors.New("invalid padding size")
	}
	n := int(data[len(data)-1])
	if n == 0 || n > blockSize || n > len(data) {
		return nil, errors.New("invalid padding")
	}
	for _, b := range data[len(data)-n:] {
		if int(b) != n {
			return nil, errors.New("invalid padding")
		}
	}
	return data[:len(data)-n], nil
}

func encryptShell(plaintext, key string) (string, error) {
	salt := make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}
	dk := pbkdf2.Key([]byte(key), salt, 10000, 48, sha256.New)
	block, err := aes.NewCipher(dk[:32])
	if err != nil {
		return "", err
	}
	padded := pkcs7Pad([]byte(plaintext), aes.BlockSize)
	encrypted := make([]byte, len(padded))
	cipher.NewCBCEncrypter(block, dk[32:]).CryptBlocks(encrypted, padded)
	out := append([]byte("Salted__"), salt...)
	out = append(out, encrypted...)
	return base64.StdEncoding.EncodeToString(out), nil
}

func decryptShell(ciphertext, key string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	if len(data) < 16 || string(data[:8]) != "Salted__" {
		return "", errors.New("invalid ciphertext")
	}
	salt := data[8:16]
	enc := data[16:]
	if len(enc)%aes.BlockSize != 0 {
		return "", errors.New("invalid ciphertext")
	}
	dk := pbkdf2.Key([]byte(key), salt, 10000, 48, sha256.New)
	block, err := aes.NewCipher(dk[:32])
	if err != nil {
		return "", err
	}
	decrypted := make([]byte, len(enc))
	cipher.NewCBCDecrypter(block, dk[32:]).CryptBlocks(decrypted, enc)
	unpadded, err := pkcs7Unpad(decrypted, aes.BlockSize)
	if err != nil {
		return "", err
	}
	return string(unpadded), nil
}
