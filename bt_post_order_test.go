package main

import "testing"

func TestBinaryTreePostOrder(t *testing.T) {
	t.Parallel()

	tree := binaryNode[int]{
		value: 20,
		right: &binaryNode[int]{
			value: 50,
			right: &binaryNode[int] {
				value: 100,
			},
			left: &binaryNode[int] {
				value: 30,
				right: &binaryNode[int] {
					value: 45,
				},
				left: &binaryNode[int] {
					value: 29,
				},
			},
		},
		left: &binaryNode[int] {
			value: 10,
			right: &binaryNode[int] {
				value: 15,
			},
			left: &binaryNode[int] {
				value: 5,
				right: &binaryNode[int] {
					value: 7,
				},
			},
		},
	}

	result := postOrderSearch(tree)
	expected := []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}
	for i, r := range result {
		if r != expected[i] {
			t.Errorf("got %v, want %v", r, expected[i])
		}
	}
}
