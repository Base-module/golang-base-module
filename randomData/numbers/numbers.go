package numbers

import (
	"github.com/Base-module/golang-base-module/randomData/base"
)

// RandomInt generates a random integer between min and max (inclusive)
func RandomInt(min, max int) int {
	if min > max {
		min, max = max, min
	}
	rng := base.NewRng()
	return rng.Intn(max-min+1) + min
}

// RandomInt64 generates a random int64 between min and max (inclusive)
func RandomInt64(min, max int64) int64 {
	if min > max {
		min, max = max, min
	}
	rng := base.NewRng()
	return rng.Int63n(max-min+1) + min
}

// RandomFloat64 generates a random float64 between min and max
func RandomFloat64(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	rng := base.NewRng()
	return min + rng.Float64()*(max-min)
}

// RandomBool generates a random boolean
func RandomBool() bool {
	rng := base.NewRng()
	return rng.Intn(2) == 1
}
