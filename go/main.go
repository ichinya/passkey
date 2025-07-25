package main

import (
	"flag"
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  passkey e -mode <safe|shell> <string>     Encrypt a string (default mode: shell)")
	fmt.Println("  passkey d -mode <safe|shell> <cipher>     Decrypt a string (default mode: shell)")
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
	fmt.Println("  PASSCRYPT_KEY=abc123 passkey e -mode shell \"mypassword\"")
	fmt.Println("  PASSCRYPT_KEY=abc123 passkey d -mode safe \"U2FsdGVk...\"")
	fmt.Println("  passkey g -length 20 -level strong -batch 5")
	fmt.Println("  PASSCRYPT_KEY=abc123 passkey g -length 24 -level paranoid -batch 10 -encrypt -mode shell")
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	mode := os.Args[1]

	switch mode {
	case "e":
		fs := flag.NewFlagSet("e", flag.ExitOnError)
		modeFlag := fs.String("mode", "shell", "encryption mode: safe or shell")
		fs.Parse(os.Args[2:])
		key := os.Getenv("PASSCRYPT_KEY")
		if key == "" {
			fmt.Println("PASSCRYPT_KEY not set")
			os.Exit(1)
		}
		if fs.NArg() < 1 {
			fmt.Println("usage: passkey e <string>")
			os.Exit(1)
		}
		result, err := Encrypt(fs.Arg(0), key, *modeFlag)
		if err != nil {
			os.Exit(2)
		}
		fmt.Println(result)

	case "d":
		fs := flag.NewFlagSet("d", flag.ExitOnError)
		modeFlag := fs.String("mode", "shell", "encryption mode: safe or shell")
		fs.Parse(os.Args[2:])
		key := os.Getenv("PASSCRYPT_KEY")
		if key == "" {
			fmt.Println("PASSCRYPT_KEY not set")
			os.Exit(1)
		}
		if fs.NArg() < 1 {
			fmt.Println("usage: passkey d <cipher>")
			os.Exit(1)
		}
		result, err := Decrypt(fs.Arg(0), key, *modeFlag)
		if err != nil {
			os.Exit(2)
		}
		fmt.Println(result)

	case "g":
		fs := flag.NewFlagSet("g", flag.ExitOnError)
		length := fs.Int("length", 16, "password length")
		level := fs.String("level", "medium", "password complexity level")
		batch := fs.Int("batch", 1, "number of passwords to generate")
		encryptGen := fs.Bool("encrypt", false, "encrypt generated passwords")
		modeFlag := fs.String("mode", "shell", "encryption mode: safe or shell")
		fs.Parse(os.Args[2:])
		key := os.Getenv("PASSCRYPT_KEY")
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
				pw, err = Encrypt(pw, key, *modeFlag)
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
