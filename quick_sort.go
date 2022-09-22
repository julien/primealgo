package main

func quickSort(src []int) {
	qs(src, 0, len(src)-1)
}

func qs(src []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(src, lo, hi)
	qs(src, lo, pivotIdx-1)
	qs(src, pivotIdx+1, hi)
}

func partition(src []int, lo, hi int) int {
	pivot := src[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if src[i] <= pivot {
			idx++
			tmp := src[i]
			src[i] = src[idx]
			src[idx] = tmp
		}
	}

	idx++
	src[hi] = src[idx]
	src[idx] = pivot
	return idx
}
