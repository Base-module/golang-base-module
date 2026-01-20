package base

import (
	"math/rand"
	"time"
)

const (
	LetterChars   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerChars    = "abcdefghijklmnopqrstuvwxyz"
	UpperChars    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DigitChars    = "0123456789"
	SpecialChars  = "!@#$%^&*()_+-=[]{}|;:,.<>?"
	AlphaNumChars = LetterChars + DigitChars
	AllChars      = AlphaNumChars + "_+"
	HexChars      = "0123456789abcdef"
)

// NewRng creates a new random number generator with time-based seed
func NewRng() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
