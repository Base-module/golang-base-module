package network

import (
	"regexp"
	"strings"
	"testing"
)

func TestRandomIPv4(t *testing.T) {
	ip := RandomIPv4()
	matched, _ := regexp.MatchString("^\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}$", ip)
	if !matched {
		t.Errorf("invalid IPv4 format: %s", ip)
	}
}

func TestRandomIPv6(t *testing.T) {
	ip := RandomIPv6()
	parts := strings.Split(ip, ":")
	if len(parts) != 8 {
		t.Errorf("IPv6 should have 8 parts, got %d: %s", len(parts), ip)
	}
}

func TestRandomMAC(t *testing.T) {
	mac := RandomMAC()
	matched, _ := regexp.MatchString("^[0-9a-f]{2}(:[0-9a-f]{2}){5}$", mac)
	if !matched {
		t.Errorf("invalid MAC format: %s", mac)
	}
}

func TestRandomURL(t *testing.T) {
	url := RandomURL()
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		t.Errorf("URL should start with http(s)://: %s", url)
	}
	if !strings.Contains(url, ".") {
		t.Errorf("URL should contain dot: %s", url)
	}
}
