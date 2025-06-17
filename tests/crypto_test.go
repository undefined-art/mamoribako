package tests

import (
	"bytes"
	"mamoribako/internal"
	"testing"
)

func TestDeriveKeyLength(t *testing.T) {
	password := []byte("testpassword")
	salt := []byte("randomsalt123456") // 16 bytes
	key := internal.DeriveKey(password, salt)
	if len(key) != 32 {
		t.Errorf("Expected key length 32, got %d", len(key))
	}
}

func TestEncryptDecrypt(t *testing.T) {
	key := []byte("01234567890123456789012345678901") // 32 bytes
	plaintext := []byte("secret data")

	ciphertext, err := internal.Encrypt(key, plaintext)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decrypted, err := internal.Decrypt(key, ciphertext)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("Decrypted text does not match original. Got %s", decrypted)
	}
}
