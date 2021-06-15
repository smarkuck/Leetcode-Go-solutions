package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func rotate(matrix [][]int) {
	for y := 0; y < len(matrix)/2; y++ {
		end := len(matrix) - y - 1
		for x := y; x < end; x++ {
			r := end - x + y
			tmp := matrix[y][x]
			matrix[y][x], matrix[r][y] = matrix[r][y], matrix[end][r]
			matrix[end][r], matrix[x][end] = matrix[x][end], tmp
		}
	}
}

func TestRotate(t *testing.T) {
	output := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}

	expected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3}}

	rotate(output)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:\n%v\nexpected:\n%v", toString(output), toString(expected))
	}

	output = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16}}

	expected = [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11}}

	rotate(output)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:\n%v\nexpected:\n%v", toString(output), toString(expected))
	}

	output = [][]int{{1}}
	expected = [][]int{{1}}
	rotate(output)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:\n%v\nexpected:\n%v", toString(output), toString(expected))
	}

	output = [][]int{
		{1, 2},
		{3, 4}}

	expected = [][]int{
		{3, 1},
		{4, 2}}

	rotate(output)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:\n%v\nexpected:\n%v", toString(output), toString(expected))
	}
}

func toString(array [][]int) string {
	result := ""
	for _, r := range array {
		result += fmt.Sprintln(r)
	}
	return result
}
