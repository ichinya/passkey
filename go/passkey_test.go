package main

import (
	"os"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestEncryptDecrypt(t *testing.T) {
	key := "my-secret-key"
	plain := "superSecure!"

	cipher, err := Encrypt(plain, key, "safe")
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
	}

	result, err := Decrypt(cipher, key, "safe")
	if err != nil {
		t.Fatalf("decryption failed: %v", err)
	}

	if result != plain {
		t.Errorf("expected %q, got %q", plain, result)
	}
}

func TestPasswordGenerationUniqueness(t *testing.T) {
	seen := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		pw, err := GeneratePassword(16, "medium")
		if err != nil {
			t.Fatalf("generate failed: %v", err)
		}
		if seen[pw] {
			t.Errorf("duplicate password: %q", pw)
		}
		seen[pw] = true
	}
}

func TestComplexityLevels(t *testing.T) {
	types := []string{"low", "medium", "strong", "paranoid"}
	for _, level := range types {
		pw, err := GeneratePassword(16, level)
		if err != nil {
			t.Errorf("%s failed: %v", level, err)
		}
		if utf8.RuneCountInString(pw) != 16 {
			t.Errorf("%s length mismatch: got %d", level, len(pw))
		}
	}
}

func TestIntegrationBetweenEncryptors(t *testing.T) {
	key := "abc123"
	plain := "P@ssw0rd2025!"

	os.Setenv("PASSCRYPT_KEY", key)
	cipher, err := Encrypt(plain, key, "safe")
	if err != nil {
		t.Fatalf("Encrypt error: %v", err)
	}

	dec, err := Decrypt(cipher, key, "safe")
	if err != nil {
		t.Fatalf("Decrypt error: %v", err)
	}

	if dec != plain {
		t.Errorf("decrypted mismatch: got %q, want %q", dec, plain)
	}

	if strings.Contains(cipher, plain) {
		t.Errorf("ciphertext leaks plaintext")
	}
}
