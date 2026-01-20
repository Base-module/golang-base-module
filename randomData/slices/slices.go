package slices

import (
	"github.com/Base-module/golang-base-module/randomData/base"
)

// RandomSliceElement returns a random element from a slice
func RandomSliceElement[T any](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}
	rng := base.NewRng()
	return slice[rng.Intn(len(slice))]
}

// RandomSliceElements returns n random elements from a slice (with possible duplicates)
func RandomSliceElements[T any](slice []T, n int) []T {
	if len(slice) == 0 || n <= 0 {
		return nil
	}
	rng := base.NewRng()
	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = slice[rng.Intn(len(slice))]
	}
	return result
}

// RandomUniqueSliceElements returns n unique random elements from a slice
func RandomUniqueSliceElements[T any](slice []T, n int) []T {
	if len(slice) == 0 || n <= 0 {
		return nil
	}
	if n > len(slice) {
		n = len(slice)
	}
	rng := base.NewRng()
	// Create a copy to avoid modifying original
	copied := make([]T, len(slice))
	copy(copied, slice)

	// Fisher-Yates shuffle for first n elements
	for i := 0; i < n; i++ {
		j := rng.Intn(len(copied)-i) + i
		copied[i], copied[j] = copied[j], copied[i]
	}
	return copied[:n]
}

// ShuffleSlice returns a shuffled copy of the slice
func ShuffleSlice[T any](slice []T) []T {
	if len(slice) == 0 {
		return nil
	}
	rng := base.NewRng()
	result := make([]T, len(slice))
	copy(result, slice)
	for i := len(result) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		result[i], result[j] = result[j], result[i]
	}
	return result
}
