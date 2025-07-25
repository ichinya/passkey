package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEncryptDecryptPassword(t *testing.T) {
	key := "mysecretkey"
	plaintexts := []string{
		"simple",
		"сложныйПароль123!",
		"pa$$w0rd!@#",
		"これは日本語です",
		"🙂🚀🔐💾",
	}
	modes := []string{"safe", "shell"}

	for _, plaintext := range plaintexts {
		for _, mode := range modes {
			t.Run(mode+"/"+plaintext, func(t *testing.T) {
				encrypted, err := Encrypt(plaintext, key, mode)
				if err != nil {
					t.Fatalf("Encrypt failed [%s:%s]: %v", mode, plaintext, err)
				}
				decrypted, err := Decrypt(encrypted, key, mode)
				if err != nil {
					t.Fatalf("Decrypt failed [%s]: %v", mode, err)
				}
				if decrypted != plaintext {
					t.Errorf("Mismatch [%s]: expected '%s', got '%s'", mode, plaintext, decrypted)
				}
			})
		}
	}
}

func TestCrossModeMismatch(t *testing.T) {
	key := "mysecretkey"
	text := "check compatibility"

	encryptedSafe, err := Encrypt(text, key, "safe")
	if err != nil {
		t.Fatal(err)
	}
	_, err = Decrypt(encryptedSafe, key, "shell")
	if err == nil {
		t.Error("Expected error when decrypting safe-encoded text with shell mode")
	}

	encryptedShell, err := Encrypt(text, key, "shell")
	if err != nil {
		t.Fatal(err)
	}
	_, err = Decrypt(encryptedShell, key, "safe")
	if err == nil {
		t.Error("Expected error when decrypting shell-encoded text with safe mode")
	}
}

func TestInvalidMode(t *testing.T) {
	_, err := Encrypt("test", "key", "invalid")
	if err == nil {
		t.Error("Expected error for invalid encryption mode")
	}
	_, err = Decrypt("something", "key", "invalid")
	if err == nil {
		t.Error("Expected error for invalid decryption mode")
	}
}

func TestRoundTripCrossSystem(t *testing.T) {
	// Предположим, что зашифрованный текст от Go с режимом shell
	// будет корректно расшифрован shell-скриптом (и наоборот).
	key := "test12345"
	text := "cross platform compatible 😎"

	encryptedShell, err := Encrypt(text, key, "shell")
	if err != nil {
		t.Fatal(err)
	}

	// Здесь можно вставить реальный вызов shell-скрипта (или имитацию)
	decrypted, err := Decrypt(encryptedShell, key, "shell")
	if err != nil || decrypted != text {
		t.Errorf("Shell mode roundtrip failed. Got: %s, Expected: %s", decrypted, text)
	}
}

func TestShellModeOpenSSLPass(t *testing.T) {
	key := "testkey"
	text := "openssl compat"

	enc, err := Encrypt(text, key, "shell")
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("openssl", "enc", "-aes-256-cbc", "-a", "-d", "-pbkdf2", "-pass", "pass:"+key)
	cmd.Stdin = strings.NewReader(enc + "\n")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("openssl decrypt failed: %v", err)
	}
	if strings.TrimSpace(string(out)) != text {
		t.Fatalf("openssl expected %s got %s", text, strings.TrimSpace(string(out)))
	}

	cmd = exec.Command("openssl", "enc", "-aes-256-cbc", "-a", "-salt", "-pbkdf2", "-pass", "pass:"+key)
	cmd.Stdin = strings.NewReader(text)
	out, err = cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	cipherOpen := strings.TrimSpace(string(out))
	dec, err := Decrypt(cipherOpen, key, "shell")
	if err != nil {
		t.Fatalf("Go decrypt failed: %v", err)
	}
	if dec != text {
		t.Fatalf("expected %s got %s", text, dec)
	}
}

func TestShellModeOpenSSL(t *testing.T) {
	key := "testkey"
	text := "openssl compat"

	enc, err := Encrypt(text, key, "shell")
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("openssl", "enc", "-aes-256-cbc", "-a", "-d", "-pbkdf2", "-pass", "pass:"+key)
	cmd.Stdin = strings.NewReader(enc + "\n")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("openssl decrypt failed: %v", err)
	}
	if strings.TrimSpace(string(out)) != text {
		t.Fatalf("openssl expected %s got %s", text, strings.TrimSpace(string(out)))
	}

	cmd = exec.Command("openssl", "enc", "-aes-256-cbc", "-a", "-salt", "-pbkdf2", "-pass", "pass:"+key)
	cmd.Stdin = strings.NewReader(text)
	out, err = cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	cipherOpen := strings.TrimSpace(string(out))
	dec, err := Decrypt(cipherOpen, key, "shell")
	if err != nil {
		t.Fatalf("Go decrypt failed: %v", err)
	}
	if dec != text {
		t.Fatalf("expected %s got %s", text, dec)
	}
}

func TestShellModeOpenSSLPassword(t *testing.T) {
	key := "testkey"
	text := "openssl compat"

	enc, err := Encrypt(text, key, "shell")
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("openssl", "enc", "-aes-256-cbc", "-a", "-d", "-pbkdf2", "-pass", "pass:"+key)
	cmd.Stdin = strings.NewReader(enc + "\n")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("openssl decrypt failed: %v", err)
	}
	if strings.TrimSpace(string(out)) != text {
		t.Fatalf("openssl expected %s got %s", text, strings.TrimSpace(string(out)))
	}

	cmd = exec.Command("openssl", "enc", "-aes-256-cbc", "-a", "-salt", "-pbkdf2", "-pass", "pass:"+key)
	cmd.Stdin = strings.NewReader(text)
	out, err = cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	cipherOpen := strings.TrimSpace(string(out))
	dec, err := Decrypt(cipherOpen, key, "shell")
	if err != nil {
		t.Fatalf("Go decrypt failed: %v", err)
	}
	if dec != text {
		t.Fatalf("expected %s got %s", text, dec)
	}
}
