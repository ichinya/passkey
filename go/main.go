package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: passkey <e|d|g> [args]")
		os.Exit(1)
	}

	mode := os.Args[1]
	os.Args = append([]string{os.Args[0]}, os.Args[2:]...)

	encMode := flag.String("mode", "shell", "encryption mode: shell or safe")
	length := flag.Int("length", 16, "password length")
	level := flag.String("level", "medium", "password complexity level")
	batch := flag.Int("batch", 1, "number of passwords to generate")
	encryptGen := flag.Bool("encrypt", false, "encrypt generated passwords")

	flag.Parse()

	key := os.Getenv("PASSCRYPT_KEY")
	switch mode {
	case "e":
		if key == "" {
			fmt.Println("PASSCRYPT_KEY not set")
			os.Exit(1)
		}
		if flag.NArg() < 1 {
			fmt.Println("usage: passkey e <string>")
			os.Exit(1)
		}
		result, err := Encrypt(flag.Arg(0), key, *encMode)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		fmt.Println(result)

	case "d":
		if key == "" {
			fmt.Println("PASSCRYPT_KEY not set")
			os.Exit(1)
		}
		if flag.NArg() < 1 {
			fmt.Println("usage: passkey d <cipher>")
			os.Exit(1)
		}
		result, err := Decrypt(flag.Arg(0), key, *encMode)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		fmt.Println(result)

	case "g":
		for i := 0; i < *batch; i++ {
			pw, err := GeneratePassword(*length, *level)
			if err != nil {
				fmt.Println("error:", err)
				os.Exit(1)
			}
			if *encryptGen {
				if key == "" {
					fmt.Println("PASSCRYPT_KEY not set")
					os.Exit(1)
				}
				pw, err = Encrypt(pw, key, *encMode)
				if err != nil {
					fmt.Println("error:", err)
					os.Exit(1)
				}
			}
			fmt.Println(pw)
		}

	default:
		fmt.Println("unknown mode")
		os.Exit(1)
	}
}
