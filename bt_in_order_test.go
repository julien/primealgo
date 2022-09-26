package main

import "testing"

func TestBinaryTreeInOrder(t *testing.T) {
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

	result := inOrderSearch(tree)
	expected := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}
	for i, r := range result {
		if r != expected[i] {
			t.Errorf("got %v, want %v", r, expected[i])
		}
	}
}
