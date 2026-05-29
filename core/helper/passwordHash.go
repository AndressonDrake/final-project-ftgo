package helper

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	memory      = 64 * 1024
	iterations  = 3
	parallelism = 2
	saltLength  = 16
	keyLength   = 32
)

func generateSalt() []byte {
	salt := make([]byte, saltLength)
	rand.Read(salt)
	return salt
}

func HashPassword(password string) string {
	salt := generateSalt()

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		iterations,
		memory,
		parallelism,
		keyLength,
	)

	return fmt.Sprintf(
		"%s.%s",
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)
}

func VerifyPassword(password, stored string) bool {
	parts := strings.Split(stored, ".")
	if len(parts) != 2 {
		return false
	}

	salt, _ := base64.RawStdEncoding.DecodeString(parts[0])
	expectedHash, _ := base64.RawStdEncoding.DecodeString(parts[1])

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		iterations,
		memory,
		parallelism,
		keyLength,
	)

	return subtle.ConstantTimeCompare(hash, expectedHash) == 1
}