package main

import "testing"

func TestEncryptDecrypt(t *testing.T) {
	key := "my-super-key"
	plaintext := "secret123"
	modes := []string{"shell", "safe"}
	for _, m := range modes {
		enc, err := Encrypt(plaintext, key, m)
		if err != nil {
			t.Fatalf("Encrypt failed in %s: %v", m, err)
		}
		dec, err := Decrypt(enc, key, m)
		if err != nil {
			t.Fatalf("Decrypt failed in %s: %v", m, err)
		}
		if dec != plaintext {
			t.Errorf("Mode %s mismatch: got %s, want %s", m, dec, plaintext)
		}
	}
}
