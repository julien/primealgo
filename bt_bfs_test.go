package main

import "testing"

func TestBinaryTreeBFS(t *testing.T) {
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

	tcs := []struct{
		name string
		needle int
		result bool
	}{
		{
			name: "Returns true for 45",
			needle: 45,
			result: true,
		},
		{
			name: "Returns true for 7",
			needle: 7,
			result: true,
		},
		{
			name: "Returns false for 69",
			needle: 69,
			result: false,
		},
	}

	for _, c := range tcs {
		t.Run(c.name, func(t *testing.T) {
			result := bfs(&tree, c.needle)
			if result != c.result {
				t.Errorf("got %v, want %v", result, c.result)
			}
		})
	}

}
