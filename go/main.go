package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/scrypt"
)

func getEnvOrPrompt(key string) string {
	val := os.Getenv(key)
	if val == "" {
		fmt.Printf("%s: ", key)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	}
	return val
}

func deriveKey(password string, salt []byte) []byte {
	key, _ := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	return key
}

func encrypt(plaintext, password string) string {
	salt := make([]byte, 16)
	rand.Read(salt)
	key := deriveKey(password, salt)

	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, gcm.NonceSize())
	rand.Read(nonce)

	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)

	out := append(salt, nonce...)
	out = append(out, ciphertext...)

	return base64.StdEncoding.EncodeToString(out)
}

func decrypt(encoded, password string) string {
	data, _ := base64.StdEncoding.DecodeString(encoded)
	salt := data[:16]
	nonce := data[16:28]
	ciphertext := data[28:]

	key := deriveKey(password, salt)
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)

	plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)
	return string(plaintext)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: passcrypt encrypt|decrypt")
		return
	}

	command := args[1]
	switch command {
	case "encrypt":
		pass := getEnvOrPrompt("PASSCRYPT_PASS")
		key := getEnvOrPrompt("PASSCRYPT_KEY")
		enc := encrypt(pass, key)
		fmt.Println("Encrypted:", enc)
	case "decrypt":
		cipher := getEnvOrPrompt("PASSCRYPT_CIPHER")
		key := getEnvOrPrompt("PASSCRYPT_KEY")
		dec := decrypt(cipher, key)
		fmt.Println("Decrypted:", dec)
	default:
		fmt.Println("Unknown command:", command)
	}
}
