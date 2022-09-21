package main

func bubbleSort(src []int) {
	l := len(src)
	for i := 0; i < l; i++ {
		for j := 0; j < l - 1 - i; j++ {
			if src[j] > src[j + 1] {
				tmp := src[j]
				src[j] = src[j+1]
				src[j+1] = tmp
			}
		}
	}
}
