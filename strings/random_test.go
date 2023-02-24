package strings

import (
	str "strings"
	"testing"
)

func TestRandomStringWithCharset(t *testing.T) {
	// Test length 0
	if RandomStringWithCharset(0, "") != "" {
		t.Errorf("RandomStringWithCharset(0, \"\") = %s; want \"\"", RandomStringWithCharset(0, ""))
	}

	// Test valid length and charset
	length := 10
	charset := "abc"
	result := RandomStringWithCharset(length, charset)
	if len(result) != length {
		t.Errorf("RandomStringWithCharset(%d, \"%s\") returned a string of length %d; want %d", length, charset, len(result), length)
	}
	for _, c := range result {
		if !str.Contains(charset, string(c)) {
			t.Errorf("RandomStringWithCharset(%d, \"%s\") returned invalid character %s", length, charset, string(c))
		}
	}
}

func TestRandomString(t *testing.T) {
	// Test length 0
	if RandomString(0) != "" {
		t.Errorf("RandomString(0) = %s; want \"\"", RandomString(0))
	}

	// Test valid length
	length := 10
	result := RandomString(length)
	if len(result) != length {
		t.Errorf("RandomString(%d) returned a string of length %d; want %d", length, len(result), length)
	}
	for _, c := range result {
		if !str.Contains(charset, string(c)) {
			t.Errorf("RandomString(%d) returned invalid character %s", length, string(c))
		}
	}
}
