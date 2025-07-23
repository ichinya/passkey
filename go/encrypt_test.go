package main

import "testing"

func TestEncryptDecrypt(t *testing.T) {
	key := "my-super-key"
	plaintext := "secret123"

	enc, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	dec, err := Decrypt(enc, key)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	if dec != plaintext {
		t.Errorf("Mismatch: got %s, want %s", dec, plaintext)
	}
}
