package leetcode

import "testing"

type Land struct {
	grid         [][]byte
	sizeX, sizeY int
}

func numIslands(grid [][]byte) int {
	return newLand(grid).getNumOfIslands()
}

func newLand(grid [][]byte) *Land {
	return &Land{grid: grid, sizeX: len(grid), sizeY: len(grid[0])}
}

func (l *Land) getNumOfIslands() int {
	sum := 0
	for x := 0; x < l.sizeX; x++ {
		for y := 0; y < l.sizeY; y++ {
			if l.grid[x][y] == '1' {
				sum++
				l.setIslandAsVisited(x, y)
			}
		}
	}
	return sum
}

func (l *Land) setIslandAsVisited(x, y int) {
	if 0 <= x && x < l.sizeX && 0 <= y && y < l.sizeY && l.grid[x][y] == '1' {
		l.grid[x][y] = '2'
		l.setIslandAsVisited(x-1, y)
		l.setIslandAsVisited(x+1, y)
		l.setIslandAsVisited(x, y-1)
		l.setIslandAsVisited(x, y+1)
	}
}

func TestNumIslands(t *testing.T) {
	output := numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}})

	expected := 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}})

	expected = 3
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
