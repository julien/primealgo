package main

type graphEdge struct {
	to, weight int
}

type weightedAdjacencyList [][]graphEdge

func dfsGraphList(graph weightedAdjacencyList, src, needle int) []int {
	seen := make([]bool, len(graph))
	for i := range seen {
		seen[i] = false
	}

	path := make([]int, 0)

	walkDFSGraphList(graph, src, needle, seen, path)

	return path
}

func walkDFSGraphList(graph weightedAdjacencyList, curr, needle int, seen []bool, path []int) bool {
	if seen[curr] {
		return false
	}

	seen[curr] = true

	// pre
	path = append(path, curr)
	if curr == needle {
		return true
	}

	// recurse
	list := graph[curr]
	for i := 0; i < len(list); i++ {
		edge := list[i]

		if walkDFSGraphList(graph, edge.to, needle, seen, path) {
			return true
		}
	}

	// post
	path = path[:len(path)-1]
	return false
}
