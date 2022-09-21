package main

import "math"

func twoCrystalBalls(breaks []bool) int {
	jmpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	
	i := jmpAmount
	l := len(breaks)
	for ; i < l; i += jmpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jmpAmount

	for j := 0; j < jmpAmount; j++ {
		if i > l - 1 {
			break
		}
		if breaks[i] {
			return i
		}
		i++
	}
	
	return -1
}
