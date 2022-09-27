package main

import "testing"

func TestDFSOnBST(t *testing.T) {
	tree := binaryNode[int]{
		value: 20,
		right: &binaryNode[int]{
			value: 50,
			right: &binaryNode[int]{
				value: 100,
			},
			left: &binaryNode[int]{
				value: 30,
				right: &binaryNode[int]{
					value: 45,
				},
				left: &binaryNode[int]{
					value: 29,
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

	if v := dfsOnBst(&tree, 45); !v {
		t.Errorf("got %v, want true", v)
	}
	if v := dfsOnBst(&tree, 7); !v {
		t.Errorf("got %v, want true", v)
	}
	if v := dfsOnBst(&tree, 69); v {
		t.Errorf("got %v, want false", v)
	}
}
