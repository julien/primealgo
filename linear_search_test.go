package main

import "testing"

func TestLinearSearch(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name     string
		haystack []int
		needle   int
		result   bool
	}{
		{
			name:     "success",
			haystack: []int{1, 2, 42, 3, 4, 5},
			needle:   42,
			result:   true,
		},
		{
			name:     "failure",
			haystack: []int{1, 2, 42, 3, 4, 5},
			needle:   69,
			result:   false,
		},
	}

	for _, c := range tcs {
		t.Run(c.name, func(t *testing.T) {
			result := linearSearch(c.haystack, c.needle)
			if result != c.result {
				t.Errorf("got %v, want %v", result, c.result)
			}
		})
	}

}
