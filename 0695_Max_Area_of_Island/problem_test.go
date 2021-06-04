package leetcode

import "testing"

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for x, row := range grid {
		for y := range row {
			area := getArea(grid, x, y)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func getArea(grid [][]int, x int, y int) int {
	if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) && grid[x][y] == 1 {
		grid[x][y] = 2
		return 1 + getArea(grid, x-1, y) + getArea(grid, x+1, y) +
			getArea(grid, x, y-1) + getArea(grid, x, y+1)
	}
	return 0
}

func TestMaxAreaOfIsland(t *testing.T) {
	output := maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}})

	expected := 6
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = maxAreaOfIsland([][]int{
		{0, 0, 0, 0, 0, 0, 0, 0}})

	expected = 0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
