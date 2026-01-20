package strings

import (
	"regexp"
	stdstrings "strings"
	"testing"
)

func TestRandomString(t *testing.T) {
	s := RandomString(10)
	if len(s) != 10 {
		t.Error("wrong length random string returned")
	}
}

func TestRandomStringFromCharset(t *testing.T) {
	charset := "abc123"
	s := RandomStringFromCharset(10, charset)
	if len(s) != 10 {
		t.Error("wrong length random string returned")
	}
	for _, c := range s {
		if !stdstrings.ContainsRune(charset, c) {
			t.Errorf("character %c not in charset", c)
		}
	}
}

func TestRandomAlphaString(t *testing.T) {
	s := RandomAlphaString(10)
	if len(s) != 10 {
		t.Error("wrong length")
	}
	matched, _ := regexp.MatchString("^[a-zA-Z]+$", s)
	if !matched {
		t.Error("string contains non-alpha characters")
	}
}

func TestRandomAlphaNumericString(t *testing.T) {
	s := RandomAlphaNumericString(10)
	if len(s) != 10 {
		t.Error("wrong length")
	}
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]+$", s)
	if !matched {
		t.Error("string contains non-alphanumeric characters")
	}
}

func TestRandomNumericString(t *testing.T) {
	s := RandomNumericString(10)
	if len(s) != 10 {
		t.Error("wrong length")
	}
	matched, _ := regexp.MatchString("^[0-9]+$", s)
	if !matched {
		t.Error("string contains non-numeric characters")
	}
}

func TestRandomLowerString(t *testing.T) {
	s := RandomLowerString(10)
	if len(s) != 10 {
		t.Error("wrong length")
	}
	matched, _ := regexp.MatchString("^[a-z]+$", s)
	if !matched {
		t.Error("string contains non-lowercase characters")
	}
}

func TestRandomUpperString(t *testing.T) {
	s := RandomUpperString(10)
	if len(s) != 10 {
		t.Error("wrong length")
	}
	matched, _ := regexp.MatchString("^[A-Z]+$", s)
	if !matched {
		t.Error("string contains non-uppercase characters")
	}
}

func TestRandomHexString(t *testing.T) {
	s := RandomHexString(16)
	if len(s) != 16 {
		t.Error("wrong length")
	}
	matched, _ := regexp.MatchString("^[0-9a-f]+$", s)
	if !matched {
		t.Error("string contains non-hex characters")
	}
}
