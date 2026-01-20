package identifier

import (
	"fmt"

	"github.com/Base-module/golang-base-module/randomData/base"
	"github.com/Base-module/golang-base-module/randomData/strings"
)

// RandomUUID generates a random UUID v4 string
func RandomUUID() string {
	rng := base.NewRng()
	uuid := make([]byte, 16)
	for i := range uuid {
		uuid[i] = byte(rng.Intn(256))
	}
	// Set version (4) and variant (2) bits
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// RandomUsername generates a random username
func RandomUsername() string {
	return strings.RandomLowerString(3) + "_" + strings.RandomAlphaNumericString(5)
}

// RandomPassword generates a random password with mixed characters
func RandomPassword(length int) string {
	if length < 4 {
		length = 4
	}
	rng := base.NewRng()
	// Ensure at least one of each type
	password := make([]byte, length)
	password[0] = base.LowerChars[rng.Intn(len(base.LowerChars))]
	password[1] = base.UpperChars[rng.Intn(len(base.UpperChars))]
	password[2] = base.DigitChars[rng.Intn(len(base.DigitChars))]
	password[3] = base.SpecialChars[rng.Intn(len(base.SpecialChars))]

	allPasswordChars := base.LowerChars + base.UpperChars + base.DigitChars + base.SpecialChars
	for i := 4; i < length; i++ {
		password[i] = allPasswordChars[rng.Intn(len(allPasswordChars))]
	}

	// Shuffle the password
	for i := len(password) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		password[i], password[j] = password[j], password[i]
	}

	return string(password)
}
