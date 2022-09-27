package main

func dfsOnBst[T Ordered](head *binaryNode[T], needle T) bool {
	return searchBst(head, needle) 
}

func searchBst[T Ordered](curr *binaryNode[T], needle T) bool {
	if curr == nil {
		return false
	}
	
	if curr.value == needle {
		return true
	}

	if curr.value < needle {
		return searchBst(curr.right, needle)
	}
	return searchBst(curr.left, needle)
}
