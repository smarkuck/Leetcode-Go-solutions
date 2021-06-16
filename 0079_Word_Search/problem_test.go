package leetcode

import "testing"

type Finder struct {
	board        [][]byte
	visited      [][]bool
	word         string
	sizeX, sizeY int
}

func exist(board [][]byte, word string) bool {
	f := newFinder(board, word)
	for x := 0; x < f.sizeX; x++ {
		for y := 0; y < f.sizeY; y++ {
			if f.dfs(0, x, y) == true {
				return true
			}
		}
	}
	return false
}

func newFinder(board [][]byte, word string) *Finder {
	sizeX, sizeY := len(board), len(board[0])
	visited := make([][]bool, sizeX)
	for i := range visited {
		visited[i] = make([]bool, sizeY)
	}
	return &Finder{board: board, visited: visited, word: word, sizeX: sizeX, sizeY: sizeY}
}

func (f *Finder) dfs(i int, x, y int) bool {
	if 0 <= x && x < f.sizeX && 0 <= y && y < f.sizeY && !f.visited[x][y] {
		if f.board[x][y] != f.word[i] {
			return false
		}
		if i == len(f.word)-1 {
			return true
		}
		f.visited[x][y] = true
		result := f.dfs(i+1, x-1, y) || f.dfs(i+1, x+1, y) || f.dfs(i+1, x, y-1) || f.dfs(i+1, x, y+1)
		f.visited[x][y] = false
		return result
	}
	return false
}

func TestExist(t *testing.T) {
	output := exist([][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE")}, "ABCCED")

	expected := true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = exist([][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE")}, "SEE")

	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = exist([][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE")}, "ABCB")

	expected = false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
