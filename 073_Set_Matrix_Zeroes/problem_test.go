package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func setZeroes(matrix [][]int) {
	isCol := false
	sizeX, sizeY := len(matrix), len(matrix[0])

	for y := 0; y < sizeY; y++ {
		if matrix[0][y] == 0 {
			isCol = true
		}
	}

	for x := 1; x < sizeX; x++ {
		for y := 0; y < sizeY; y++ {
			if matrix[x][y] == 0 {
				matrix[x][0] = 0
				matrix[0][y] = 0
			}
		}
	}

	for x := 1; x < sizeX; x++ {
		for y := 1; y < sizeY; y++ {
			if matrix[x][0] == 0 || matrix[0][y] == 0 {
				matrix[x][y] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for x := 0; x < sizeX; x++ {
			matrix[x][0] = 0
		}
	}

	if isCol {
		for y := 0; y < sizeY; y++ {
			matrix[0][y] = 0
		}
	}
}

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1}}

	expected := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1}}

	setZeroes(matrix)
	if !reflect.DeepEqual(matrix, expected) {
		t.Errorf("\noutput:\n%v\nexpected:\n%v", toString(matrix), toString(expected))
	}

	matrix = [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5}}

	expected = [][]int{
		{0, 0, 0, 0},
		{0, 4, 5, 0},
		{0, 3, 1, 0}}

	setZeroes(matrix)
	if !reflect.DeepEqual(matrix, expected) {
		t.Errorf("\noutput:\n%v\nexpected:\n%v", toString(matrix), toString(expected))
	}
}

func toString(array [][]int) string {
	result := ""
	for _, r := range array {
		result += fmt.Sprintln(r)
	}
	return result
}
