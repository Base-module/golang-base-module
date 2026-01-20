package datetime

import (
	"testing"
	"time"
)

func TestRandomDate(t *testing.T) {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
	for i := 0; i < 100; i++ {
		d := RandomDate(start, end)
		if d.Before(start) || d.After(end) {
			t.Errorf("date %v outside range", d)
		}
	}
}

func TestRandomDateString(t *testing.T) {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)
	s := RandomDateString(start, end, "2006-01-02")
	_, err := time.Parse("2006-01-02", s)
	if err != nil {
		t.Errorf("invalid date format: %s", s)
	}
}
