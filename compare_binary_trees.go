package main

func compareBinaryTrees[T comparable](a, b *binaryNode[T]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.value != b.value {
		return false
	}
	
	return compareBinaryTrees(a.left, b.left) && compareBinaryTrees(a.right, b.right)
}
