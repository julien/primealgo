package main

func walkPostOrder[T comparable](curr *binaryNode[T], path []T) []T {
	if curr == nil {
		return path
	}
	// pre 	

	// recurse 
	path = walkPostOrder(curr.left, path)
	path = walkPostOrder(curr.right, path)
	
	// post
	path = append(path, curr.value)
	return path
}

func postOrderSearch[T comparable](head binaryNode[T]) []T {
	return walkPostOrder(&head, make([]T, 0))
}
