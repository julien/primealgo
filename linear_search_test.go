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

func benchmarkLinearSearch(i int, b *testing.B) {
	haystack := []int{1, 2, 42, 3, 4, 5}
	for n := 0; n < b.N; n++ {
		linearSearch(haystack, i)
	}
}

func BenchmarkLinearSearch10(b *testing.B)    { benchmarkLinearSearch(10, b) }
func BenchmarkLinearSearch100(b *testing.B)   { benchmarkLinearSearch(100, b) }
func BenchmarkLinearSearch1000(b *testing.B)  { benchmarkLinearSearch(1000, b) }
func BenchmarkLinearSearch10000(b *testing.B) { benchmarkLinearSearch(10000, b) }
