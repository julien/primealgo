package main

import "testing"

func TestBinarySearch(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name     string
		haystack []int
		needle   int
		result   bool
	}{
		{
			name:     "success",
			haystack: []int{1, 2, 3, 4, 5},
			needle:   4,
			result:   true,
		},
		{
			name:     "failure",
			haystack: []int{1, 2, 3, 4, 5},
			needle:   69,
			result:   false,
		},
	}

	for _, c := range tcs {
		t.Run(c.name, func(t *testing.T) {
			result := binarySearch(c.haystack, c.needle)
			if result != c.result {
				t.Errorf("got %v, want %v", result, c.result)
			}
		})
	}
}

