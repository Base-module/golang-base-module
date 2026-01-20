package contact

import (
	"regexp"
	"strings"
	"testing"
)

func TestRandomEmail(t *testing.T) {
	s := RandomEmail()
	expectedLength := 25 + len("@example.com")
	if len(s) != expectedLength {
		t.Errorf("wrong length random email returned; expected %d, got %d", expectedLength, len(s))
	}
	if !strings.Contains(s, "@example.com") {
		t.Error("random email does not contain @example.com domain")
	}
}

func TestRandomEmailWithDomain(t *testing.T) {
	domain := "test.org"
	s := RandomEmailWithDomain(domain)
	if !strings.HasSuffix(s, "@"+domain) {
		t.Errorf("email should end with @%s, got %s", domain, s)
	}
}

func TestRandomPhone(t *testing.T) {
	phone := RandomPhone()
	matched, _ := regexp.MatchString("^\\d{3}-\\d{3}-\\d{4}$", phone)
	if !matched {
		t.Errorf("invalid phone format: %s", phone)
	}
}

func TestRandomPhoneWithPrefix(t *testing.T) {
	phone := RandomPhoneWithPrefix("+1")
	if !strings.HasPrefix(phone, "+1 ") {
		t.Errorf("phone should start with +1, got %s", phone)
	}
}
