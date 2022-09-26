package main

import "testing"

func TestCompareBinaryTrees(t *testing.T) {
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

	tree2 := binaryNode[int]{
		value: 20,
		right: &binaryNode[int]{
			value: 50,
			left: &binaryNode[int]{
				value: 30,
				right: &binaryNode[int]{
					value: 45,
					right: &binaryNode[int]{
						value: 49,
					},
				},
				left: &binaryNode[int]{
					value: 29,
					left: &binaryNode[int]{
						value: 21,
					},
				},
			},
		},
		left: &binaryNode[int]{
			value: 10,
			right: &binaryNode[int]{
				value: 15,
			},
			left: &binaryNode[int]{
				value: 5,
				right: &binaryNode[int]{
					value: 7,
				},
			},
		},
	}

	if result := compareBinaryTrees(&tree, &tree); !result {
		t.Errorf("got %v, want true", result)
	}

	if result := compareBinaryTrees(&tree, &tree2); result {
		t.Errorf("got %v, want false", result)
	}
}
