package main

func walkInOrder[T comparable](curr *binaryNode[T], path []T) []T {
	if curr == nil {
		return path
	}
	// pre

	// recurse
	path = walkInOrder(curr.left, path)
	path = append(path, curr.value)
	path = walkInOrder(curr.right, path)

	// post
	return path
}

func inOrderSearch[T comparable](head binaryNode[T]) []T {
	return walkInOrder(&head, make([]T, 0))
}
