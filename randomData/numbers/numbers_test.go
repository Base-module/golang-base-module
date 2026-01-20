package numbers

import "testing"

func TestRandomInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := RandomInt(5, 10)
		if n < 5 || n > 10 {
			t.Errorf("RandomInt(5,10) returned %d, expected between 5 and 10", n)
		}
	}
	// Test with swapped min/max
	n := RandomInt(10, 5)
	if n < 5 || n > 10 {
		t.Errorf("RandomInt(10,5) returned %d, expected between 5 and 10", n)
	}
}

func TestRandomInt64(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := RandomInt64(100, 200)
		if n < 100 || n > 200 {
			t.Errorf("RandomInt64(100,200) returned %d", n)
		}
	}
}

func TestRandomFloat64(t *testing.T) {
	for i := 0; i < 100; i++ {
		f := RandomFloat64(1.0, 5.0)
		if f < 1.0 || f > 5.0 {
			t.Errorf("RandomFloat64(1.0,5.0) returned %f", f)
		}
	}
}

func TestRandomBool(t *testing.T) {
	trueCount := 0
	for i := 0; i < 100; i++ {
		if RandomBool() {
			trueCount++
		}
	}
	// Should have roughly 50% true, allow wide margin
	if trueCount == 0 || trueCount == 100 {
		t.Error("RandomBool seems biased")
	}
}
