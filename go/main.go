package main

import (
	"flag"
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  passkey e <string> [-mode safe|shell]     Encrypt a string")
	fmt.Println("  passkey d <cipher> [-mode safe|shell]     Decrypt a string")
	fmt.Println("  passkey g [flags]                          Generate password(s)")
	fmt.Println()
	fmt.Println("Flags for 'g':")
	fmt.Println("  -length int                    Length of password (default 16)")
	fmt.Println("  -level string                  Complexity level: low, medium, strong, paranoid")
	fmt.Println("  -batch int                     Number of passwords to generate")
	fmt.Println("  -encrypt                       Encrypt generated passwords")
	fmt.Println("  -mode string                   Encryption mode: safe or shell (default safe)")
	fmt.Println()
	fmt.Println("Environment:")
	fmt.Println("  PASSCRYPT_KEY must be set for encryption/decryption")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  PASSCRYPT_KEY=abc123 passkey e \"mypassword\" -mode shell")
	fmt.Println("  PASSCRYPT_KEY=abc123 passkey d \"U2FsdGVk...\" -mode safe")
	fmt.Println("  passkey g -length 20 -level strong -batch 5")
	fmt.Println("  PASSCRYPT_KEY=abc123 passkey g -length 24 -level paranoid -batch 10 -encrypt -mode shell")
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	mode := os.Args[1]
	os.Args = append([]string{os.Args[0]}, os.Args[2:]...)

	length := flag.Int("length", 16, "password length")
	level := flag.String("level", "medium", "password complexity level")
	batch := flag.Int("batch", 1, "number of passwords to generate")
	encryptGen := flag.Bool("encrypt", false, "encrypt generated passwords")
	safeMode := flag.String("mode", "safe", "encryption mode: safe or shell")

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
		result, err := Encrypt(flag.Arg(0), key, *safeMode)
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
		result, err := Decrypt(flag.Arg(0), key, *safeMode)
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
				pw, err = Encrypt(pw, key, *safeMode)
				if err != nil {
					fmt.Println("error:", err)
					os.Exit(1)
				}
			}
			fmt.Println(pw)
		}

	default:
		printHelp()
		os.Exit(1)
	}
}
