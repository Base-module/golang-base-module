package datetime

import (
	"time"

	"github.com/Base-module/golang-base-module/randomData/numbers"
)

// RandomDate generates a random date between start and end
func RandomDate(start, end time.Time) time.Time {
	if start.After(end) {
		start, end = end, start
	}
	delta := end.Unix() - start.Unix()
	sec := numbers.RandomInt64(0, delta)
	return time.Unix(start.Unix()+sec, 0)
}

// RandomDateString generates a random date string in specified format
func RandomDateString(start, end time.Time, format string) string {
	return RandomDate(start, end).Format(format)
}
