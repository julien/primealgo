package main

import "testing"

func TestMazeSolver(t *testing.T) {
	t.Parallel()

	maze := []string {
        "xxxxxxxxxx x",
        "x        x x",
        "x        x x",
        "x xxxxxxxx x",
        "x          x",
        "x xxxxxxxxxx",
	}

	result := maze_solver(maze, "x", point{x: 10, y: 0}, point{x: 1,  y: 5})

	expected := []*point {
		{x: 10 , y: 0},
		{x: 10, y: 1},  
		{x: 10, y: 2},  
		{x: 10, y: 3},  
		{x: 10, y: 4},  
		{x: 9, y: 4},  
		{x: 8, y: 4},  
		{x: 7, y: 4},  
		{x: 6, y: 4},  
		{x: 5, y: 4},  
		{x: 4, y: 4},  
		{x: 3, y: 4},  
		{x: 2, y: 4},  
		{x: 1, y: 4},  
		{x: 1, y: 5}, 
	}

	for i, p1 := range result {
		p2 := expected[i]
		if p1.x != p2.x && p1.y != p2.y {
			t.Errorf("got %v, want %v", p1, p2)
		}
	}
}
