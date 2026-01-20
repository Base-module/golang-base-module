package contact

import (
	"fmt"

	"github.com/Base-module/golang-base-module/randomData/base"
	"github.com/Base-module/golang-base-module/randomData/strings"
)

// RandomEmail generates a random email address
func RandomEmail() string {
	rng := base.NewRng()
	s, r := make([]rune, 25), []rune(base.AllChars)
	for i := range s {
		s[i] = r[rng.Intn(len(r))]
	}
	return string(s) + "@example.com"
}

// RandomEmailWithDomain generates a random email with specified domain
func RandomEmailWithDomain(domain string) string {
	username := strings.RandomLowerString(10)
	return username + "@" + domain
}

// RandomPhone generates a random phone number (format: XXX-XXX-XXXX)
func RandomPhone() string {
	return fmt.Sprintf("%s-%s-%s",
		strings.RandomNumericString(3),
		strings.RandomNumericString(3),
		strings.RandomNumericString(4))
}

// RandomPhoneWithPrefix generates a random phone number with country prefix
func RandomPhoneWithPrefix(prefix string) string {
	return fmt.Sprintf("%s %s-%s-%s",
		prefix,
		strings.RandomNumericString(3),
		strings.RandomNumericString(3),
		strings.RandomNumericString(4))
}
