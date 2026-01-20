package slices

import "testing"

func TestRandomSliceElement(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	for i := 0; i < 50; i++ {
		elem := RandomSliceElement(slice)
		found := false
		for _, s := range slice {
			if s == elem {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("element %s not in slice", elem)
		}
	}

	// Test empty slice
	var emptySlice []string
	result := RandomSliceElement(emptySlice)
	if result != "" {
		t.Error("empty slice should return zero value")
	}
}

func TestRandomSliceElements(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	elements := RandomSliceElements(slice, 3)
	if len(elements) != 3 {
		t.Errorf("expected 3 elements, got %d", len(elements))
	}

	// Test empty slice
	result := RandomSliceElements([]int{}, 3)
	if result != nil {
		t.Error("empty slice should return nil")
	}
}

func TestRandomUniqueSliceElements(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	elements := RandomUniqueSliceElements(slice, 3)
	if len(elements) != 3 {
		t.Errorf("expected 3 elements, got %d", len(elements))
	}
	// Check uniqueness
	seen := make(map[int]bool)
	for _, e := range elements {
		if seen[e] {
			t.Error("duplicate element found")
		}
		seen[e] = true
	}

	// Test requesting more than available
	elements = RandomUniqueSliceElements(slice, 10)
	if len(elements) != 5 {
		t.Errorf("expected 5 elements (max available), got %d", len(elements))
	}
}

func TestShuffleSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	shuffled := ShuffleSlice(slice)
	if len(shuffled) != len(slice) {
		t.Error("shuffled slice has wrong length")
	}

	// Check all elements present
	sum := 0
	for _, v := range shuffled {
		sum += v
	}
	if sum != 55 {
		t.Error("shuffled slice missing elements")
	}

	// Test empty slice
	result := ShuffleSlice([]int{})
	if result != nil {
		t.Error("empty slice should return nil")
	}
}
