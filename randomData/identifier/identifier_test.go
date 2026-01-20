package identifier

import (
	"regexp"
	"strings"
	"testing"
)

func TestRandomUUID(t *testing.T) {
	uuid := RandomUUID()
	matched, _ := regexp.MatchString("^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$", uuid)
	if !matched {
		t.Errorf("invalid UUID format: %s", uuid)
	}
}

func TestRandomUsername(t *testing.T) {
	username := RandomUsername()
	if !strings.Contains(username, "_") {
		t.Errorf("username should contain underscore: %s", username)
	}
	if len(username) != 9 { // 3 + 1 + 5
		t.Errorf("wrong username length: %s", username)
	}
}

func TestRandomPassword(t *testing.T) {
	password := RandomPassword(12)
	if len(password) != 12 {
		t.Errorf("password should be 12 chars, got %d", len(password))
	}
	// Should contain at least one of each type
	hasLower, _ := regexp.MatchString("[a-z]", password)
	hasUpper, _ := regexp.MatchString("[A-Z]", password)
	hasDigit, _ := regexp.MatchString("[0-9]", password)
	hasSpecial, _ := regexp.MatchString("[!@#$%^&*()_+\\-=\\[\\]{}|;:,.<>?]", password)

	if !hasLower || !hasUpper || !hasDigit || !hasSpecial {
		t.Errorf("password missing required character types: %s", password)
	}

	// Test minimum length
	shortPassword := RandomPassword(2)
	if len(shortPassword) < 4 {
		t.Error("minimum password length should be 4")
	}
}
