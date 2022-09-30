package main

import "testing"

func TestBFSGraphMatrix(t *testing.T) {
	matrix := [][]int{
		{0, 3, 1, 0, 0, 0, 0}, // 0
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 5, 0, 2, 0},
		{0, 0, 18, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 0, 1},
	}

	result := bfsGraphMatrix(matrix, 0, 6)
	expected := []int{0, 1, 4, 5, 6}

	for i, r := range result {
		if r != expected[i] {
			t.Errorf("got %v, want %v", r, expected[i])
		}
	}
}
