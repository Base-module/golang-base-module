package strings

import (
	"github.com/Base-module/golang-base-module/randomData/base"
)

// RandomString generates a random string of specified length
// Note: For better randomness in production, consider using crypto/rand
func RandomString(n int) string {
	rng := base.NewRng()
	s, r := make([]rune, n), []rune(base.AllChars)
	for i := range s {
		s[i] = r[rng.Intn(len(r))]
	}
	return string(s)
}

// RandomStringFromCharset generates a random string from custom charset
func RandomStringFromCharset(n int, charset string) string {
	rng := base.NewRng()
	s, r := make([]rune, n), []rune(charset)
	for i := range s {
		s[i] = r[rng.Intn(len(r))]
	}
	return string(s)
}

// RandomAlphaString generates a random alphabetic string (letters only)
func RandomAlphaString(n int) string {
	return RandomStringFromCharset(n, base.LetterChars)
}

// RandomAlphaNumericString generates a random alphanumeric string
func RandomAlphaNumericString(n int) string {
	return RandomStringFromCharset(n, base.AlphaNumChars)
}

// RandomNumericString generates a random numeric string
func RandomNumericString(n int) string {
	return RandomStringFromCharset(n, base.DigitChars)
}

// RandomLowerString generates a random lowercase string
func RandomLowerString(n int) string {
	return RandomStringFromCharset(n, base.LowerChars)
}

// RandomUpperString generates a random uppercase string
func RandomUpperString(n int) string {
	return RandomStringFromCharset(n, base.UpperChars)
}

// RandomHexString generates a random hexadecimal string
func RandomHexString(n int) string {
	return RandomStringFromCharset(n, base.HexChars)
}
