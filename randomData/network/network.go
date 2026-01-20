package network

import (
	"fmt"
	stdstrings "strings"

	"github.com/Base-module/golang-base-module/randomData/numbers"
	"github.com/Base-module/golang-base-module/randomData/slices"
	"github.com/Base-module/golang-base-module/randomData/strings"
)

// RandomIPv4 generates a random IPv4 address
func RandomIPv4() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		numbers.RandomInt(1, 255),
		numbers.RandomInt(0, 255),
		numbers.RandomInt(0, 255),
		numbers.RandomInt(1, 254))
}

// RandomIPv6 generates a random IPv6 address
func RandomIPv6() string {
	parts := make([]string, 8)
	for i := 0; i < 8; i++ {
		parts[i] = strings.RandomHexString(4)
	}
	return stdstrings.Join(parts, ":")
}

// RandomMAC generates a random MAC address
func RandomMAC() string {
	return fmt.Sprintf("%s:%s:%s:%s:%s:%s",
		strings.RandomHexString(2),
		strings.RandomHexString(2),
		strings.RandomHexString(2),
		strings.RandomHexString(2),
		strings.RandomHexString(2),
		strings.RandomHexString(2))
}

// RandomURL generates a random URL
func RandomURL() string {
	protocols := []string{"http", "https"}
	tlds := []string{"com", "org", "net", "io", "dev"}
	return fmt.Sprintf("%s://%s.%s/%s",
		slices.RandomSliceElement(protocols),
		strings.RandomLowerString(8),
		slices.RandomSliceElement(tlds),
		strings.RandomLowerString(10))
}
