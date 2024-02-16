package utils

import (
	"math/rand"
	"time"
)

// init seeds the random number generator; do this only once to ensure true randomness.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateRandomString generates a random string of 6 characters.
func GenerateRandomString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
