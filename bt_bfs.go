package main

func bfs[T comparable](head *binaryNode[T], needle T) bool {
	q := make([]*binaryNode[T], 0)
	q = append(q, head)

	l := len(q)
	for l > 0 {
		curr := q[0]

		q = q[1:]
		l = len(q)

		// search
		if curr.value == needle {
			return true
		}

		if curr.left != nil {
			q = append(q, curr.left)
			l = len(q)
		}
		if curr.right != nil {
			q = append(q, curr.right)
			l = len(q)
		}
	}

	return false
}
