package bytes

import "testing"

func TestRandomBytes(t *testing.T) {
	bytes := RandomBytes(16)
	if len(bytes) != 16 {
		t.Errorf("expected 16 bytes, got %d", len(bytes))
	}

	// Test zero/negative
	result := RandomBytes(0)
	if result != nil {
		t.Error("zero length should return nil")
	}
	result = RandomBytes(-5)
	if result != nil {
		t.Error("negative length should return nil")
	}
}
