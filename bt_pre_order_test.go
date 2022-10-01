package main

import "testing"

func TestBinaryTreePreOrder(t *testing.T) {
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
				value: 15,
				right: &binaryNode[int] {
					value: 7,
				},
			},
		},
	}

	result := preOrderSearch(tree)
	expected := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	for i, r := range result {
		if r != expected[i] {
			t.Errorf("got %v, want %v", r, expected[i])
		}
	}
}
