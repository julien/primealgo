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

func benchmarkBinarySearch(i int, b *testing.B) {
	haystack := []int{1, 2, 42, 3, 4, 5}
	for n := 0; n < b.N; n++ {
		binarySearch(haystack, i)
	}
}

func BenchmarkBinarySearch10(b *testing.B)    { benchmarkBinarySearch(10, b) }
func BenchmarkBinarySearch100(b *testing.B)   { benchmarkBinarySearch(100, b) }
func BenchmarkBinarySearch1000(b *testing.B)  { benchmarkBinarySearch(1000, b) }
func BenchmarkBinarySearch10000(b *testing.B) { benchmarkBinarySearch(10000, b) }
