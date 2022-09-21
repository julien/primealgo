package main

func linearSearch(haystack []int, needle int) bool {
	for _, n := range haystack {
		if n == needle {
			return true
		}
	}
	return false
}
