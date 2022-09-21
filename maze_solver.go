package main

type point struct {
	x, y int
}

var dir = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func maze_solver(maze []string, wall string, start, end point) []*point {
	rows := len(maze)
	cols := len(maze[0])

	seen := make([][]bool, rows)

	for i := range seen {
		seen[i] = make([]bool, cols)
		for j := range seen[i] {
			seen[i][j] = false
		}
	}

	path := make([]*point, 0)
	_, path = walk(maze, wall, start, end, seen, path)
	return path
}

func walk(maze []string, wall string, curr, end point, seen [][]bool, path []*point) (bool, []*point) {
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false, path
	}
	s := string(maze[curr.y][curr.x])
	if s == wall {
		return false, path
	}
	if curr.x == end.x && curr.y == end.y {
		path = append(path, &end)
		return true, path
	}
	if seen[curr.y][curr.x] {
		return false, path
	}

	seen[curr.y][curr.x] = true
	path = append(path, &curr)

	for i := 0; i < len(dir); i++ {
		next := dir[i]
		np := point{x: curr.x + next[0], y: curr.y + next[1]}
		if r, p := walk(maze, wall, np, end, seen, path); r {
			return true, p
		}
	}

	path = path[:len(path)-1]
	return false, path
}
