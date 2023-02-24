package strings

import (
	"math/rand"
	"time"
)

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()),
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomStringWithCharset returns a random string from a specific set of
// characters with a given length.
func RandomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// RandomString returns a random string from the standard English alphabet with
// a given length.
func RandomString(length int) string {
	return RandomStringWithCharset(length, charset)
}
