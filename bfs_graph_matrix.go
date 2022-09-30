package main

func bfsGraphMatrix(graph [][]int, src int, needle int) []int {
	n := len(graph)
	seen := make([]bool, n)
	for i := range seen {
		seen[i] = false
	}
	prev := make([]int, n)
	for i := range prev {
		prev[i] = -1
	}

	seen[src] = true
	q := make([]int, 0)
	q = append(q, src)

	l := len(q)
	for l > 0 {
		curr := q[0]
		q = q[1:]
		l = len(q)

		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i := 0; i < len(adjs); i++ {
			if adjs[i] == 0 {
				continue
			}
			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			q = append(q, i)
			l = len(q)
		}
	}

	out := make([]int, 0)
	if prev[needle] == -1 {
		return out
	}

	// build backwards
	curr := needle

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	n = len(out) + 1
	res := make([]int, 0, n)
	res = append(res, src)
	for i := 0; i < n-1; i++ {
		res = append(res, out[i])
	}
	return res
}
