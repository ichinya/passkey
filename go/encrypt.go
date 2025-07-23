package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: passkey [e|d] <string>")
		os.Exit(1)
	}

	mode := os.Args[1]
	data := os.Args[2]
	key := os.Getenv("PASSCRYPT_KEY")
	if key == "" {
		fmt.Println("Error: PASSCRYPT_KEY is not set")
		os.Exit(1)
	}

	switch mode {
	case "e":
		result, err := Encrypt(data, key)
		if err != nil {
			fmt.Println("Encryption error:", err)
			os.Exit(1)
		}
		fmt.Println(result)
	case "d":
		result, err := Decrypt(data, key)
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
