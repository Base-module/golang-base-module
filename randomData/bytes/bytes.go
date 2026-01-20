package bytes

import (
	"github.com/Base-module/golang-base-module/randomData/base"
)

// RandomBytes generates n random bytes
func RandomBytes(n int) []byte {
	if n <= 0 {
		return nil
	}
	rng := base.NewRng()
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(rng.Intn(256))
	}
	return b
}
