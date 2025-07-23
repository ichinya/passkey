package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

	"golang.org/x/crypto/scrypt"
)

func deriveKey(password string, salt []byte) []byte {
	key, _ := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	return key
}

func encrypt(plaintext, password string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	key := deriveKey(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)
	out := append(salt, nonce...)
	out = append(out, ciphertext...)

	return base64.StdEncoding.EncodeToString(out), nil
}

func decrypt(encoded, password string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	salt := data[:16]
	nonce := data[16:28]
	ciphertext := data[28:]

	key := deriveKey(password, salt)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: passkey [e|d] <string>")
		os.Exit(1)
	}

	mode := os.Args[1]
	input := os.Args[2]
	key := os.Getenv("PASSCRYPT_KEY")

	if key == "" {
		fmt.Println("Error: set PASSCRYPT_KEY environment variable.")
		os.Exit(1)
	}

	switch mode {
	case "e":
		result, err := encrypt(input, key)
		if err != nil {
			fmt.Println("Encryption error:", err)
			os.Exit(1)
		}
		fmt.Println(result)
	case "d":
		result, err := decrypt(input, key)
		if err != nil {
			fmt.Println("Decryption error:", err)
			os.Exit(1)
		}
		fmt.Println(result)
	default:
		fmt.Println("Unknown mode:", mode)
		os.Exit(1)
	}
}
