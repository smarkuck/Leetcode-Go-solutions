package leetcode

import "testing"

type Matrix struct {
	matrix       [][]int
	cache        [][]int
	sizeX, sizeY int
}

type Field struct {
	x, y int
}

var directions = []Field{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func longestIncreasingPath(matrix [][]int) int {
	m := newMatrix(matrix)

	ans := 0
	for x := 0; x < m.sizeX; x++ {
		for y := 0; y < m.sizeY; y++ {
			ans = max(m.dfs(Field{x, y}), ans)
		}
	}

	return ans
}

func newMatrix(matrix [][]int) *Matrix {
	m := &Matrix{
		matrix: matrix, cache: nil,
		sizeX: len(matrix), sizeY: len(matrix[0])}

	m.cache = make([][]int, m.sizeX)
	for i := range m.cache {
		m.cache[i] = make([]int, m.sizeY)
	}

	return m
}

func (m *Matrix) dfs(f Field) int {
	if m.cache[f.x][f.y] > 0 {
		return m.cache[f.x][f.y]
	}

	for _, d := range directions {
		neighbour := Field{f.x + d.x, f.y + d.y}
		if 0 <= neighbour.x && neighbour.x < m.sizeX &&
			0 <= neighbour.y && neighbour.y < m.sizeY &&
			m.matrix[f.x][f.y] < m.matrix[neighbour.x][neighbour.y] {
			m.cache[f.x][f.y] = max(m.dfs(neighbour), m.cache[f.x][f.y])
		}
	}

	m.cache[f.x][f.y]++
	return m.cache[f.x][f.y]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestLongestIncreasingPath(t *testing.T) {
	output := longestIncreasingPath([][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1}})

	expected := 4
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = longestIncreasingPath([][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1}})

	expected = 4
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = longestIncreasingPath([][]int{{1}})
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
