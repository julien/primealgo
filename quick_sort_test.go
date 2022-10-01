package main

import "testing"

func TestQuickSort(t *testing.T) {
	t.Parallel()

	src := []int{9, 3, 7, 4, 69, 420, 42}
	quickSort(src)
	expected := []int{3, 4, 7, 9, 42, 69, 420}
	for i, e := range src {
		if e != expected[i] {
			t.Errorf("got %v, want %v", e, expected[i])
		}
	}
}
