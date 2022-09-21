package main

import "testing"

func TestBubbleSort(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		src  []int
		dst  []int
	}{
		{
			name: "success",
			src:  []int{1, 2, 42, 3, 4, 5},
			dst:  []int{1, 2, 3, 4, 5, 42},
		},
		{
			name: "success",
			src:  []int{1, 2, 42, 72, 36, 28},
			dst:  []int{1, 2, 28, 36, 42, 72},
		},
	}

	for _, c := range tcs {
		t.Run(c.name, func(t *testing.T) {
			bubbleSort(c.src)

			for i := range c.dst {
				if c.dst[i] != c.src[i] {
					t.Errorf("want %v, got %v", c.dst[i], c.src[i])
				}
			}
		})
	}

}
