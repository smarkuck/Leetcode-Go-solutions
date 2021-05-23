package leetcode

import (
	"bytes"
	"reflect"
	"testing"
)

func solveSudoku(board [][]byte) {
	NewSudoku(board).solve()
}

type Sudoku struct {
	board  [][]byte
	rowMap [81]bool
	colMap [81]bool
	boxMap [81]bool
}

func NewSudoku(board [][]byte) *Sudoku {
	s := &Sudoku{
		board: board,
	}
	for i, row := range board {
		for j, digit := range row {
			if digit != '.' {
				s.insert(i, j, digit)
			}
		}
	}
	return s
}

func (s *Sudoku) solve() {
	s.trySolveFrom(0, 0)
}

func (s *Sudoku) trySolveFrom(x, y int) bool {
	if y == 9 {
		x = x + 1
		y = 0
	}
	if x == 9 {
		return true
	}
	if s.board[x][y] == '.' {
		for _, digit := range []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'} {
			if !s.canInsert(x, y, digit) {
				continue
			}
			s.insert(x, y, digit)

			if s.trySolveFrom(x, y+1) {
				return true
			}
			s.remove(x, y)
		}
	} else {
		return s.trySolveFrom(x, y+1)
	}
	return false
}

func (s *Sudoku) canInsert(x, y int, digit byte) bool {
	if s.board[x][y] != '.' {
		return false
	}
	idx := int(digit - '1')
	return !(s.rowMap[x*9+idx] || s.colMap[y*9+idx] || s.boxMap[((x/3)*3+(y/3))*9+idx])
}

func (s *Sudoku) insert(x, y int, digit byte) {
	s.board[x][y] = digit
	idx := int(digit - '1')
	s.rowMap[x*9+idx] = true
	s.colMap[y*9+idx] = true
	s.boxMap[((x/3)*3+(y/3))*9+idx] = true
}

func (s *Sudoku) remove(x, y int) {
	idx := int(s.board[x][y] - '1')
	s.rowMap[x*9+idx] = false
	s.colMap[y*9+idx] = false
	s.boxMap[((x/3)*3+(y/3))*9+idx] = false
	s.board[x][y] = '.'
}

func TestSolveSudoku(t *testing.T) {
	sudoku := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	expected := [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'}}

	solveSudoku(sudoku)
	if !reflect.DeepEqual(sudoku, expected) {
		t.Errorf("\noutput:\n%c\n\nexpected:\n%c",
			bytes.Join(sudoku, []byte("\n")),
			bytes.Join(expected, []byte("\n")))
	}
}
