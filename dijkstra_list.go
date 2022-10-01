package main

import "math"

func dijkstraList(source, sink int, arr weightedAdjacencyList) []int {
	n := len(arr)
	seen := make([]bool, n)
	for i := range seen {
		seen[i] = false
	}

	prev := make([]int, n)
	for i := range prev {
		prev[i] = -1
	}

	inf := math.Inf(1)
	dists := make([]float64, n)
	for i := range dists {
		dists[i] = inf
	}
	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := lowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := arr[curr]
		for i := range adjs {
			edge := adjs[i]
			if seen[edge.to] {
				continue
			}

			dist := int(dists[curr]) + edge.weight
			if dist < int(dists[edge.to]) {
				dists[edge.to] = float64(dist)
				prev[edge.to] = curr
			}
		}
	}

	out := make([]int, 0)
	curr := sink
	
	for prev[curr] != -1 {
		 out = append(out, curr)
		 curr = prev[curr]
	}

	out = append(out, source)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}

func hasUnvisited(seen []bool, dists []float64) bool {
	inf := math.Inf(1)
	for i := range seen {
		if !seen[i] && dists[i] < inf {
			return true
		}
	}
	return false
}

func lowestUnvisited(seen []bool, dists []float64) int {
	idx := -1
	inf := math.Inf(1)
	lowest := inf

	for i := range seen {
		if seen[i] {
			continue
		}
		if lowest > dists[i] {
			lowest = dists[i]
			idx = i
		}
	}
	return idx
}
