package main

func walkPreOrder[T comparable](curr *binaryNode[T], path []T) []T {
	if curr == nil {
		return path
	}
	// pre 	
	path = append(path, curr.value)

	// recurse 
	walkPreOrder(curr.left, path)
	walkPreOrder(curr.right, path)
	
	// post
	return path
}

func preOrderSearch[T comparable](head binaryNode[T]) []T {
	return walkPreOrder(&head, make([]T, 0))
}
