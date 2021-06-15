package leetcode

import (
	"reflect"
	"testing"
)

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	l, t, r, b := -1, -1, len(matrix[0]), len(matrix)
	for {
		t++
		for x := l + 1; x < r; x++ {
			result = append(result, matrix[t][x])
		}
		if t+1 == b {
			break
		}
		r--
		for y := t + 1; y < b; y++ {
			result = append(result, matrix[y][r])
		}
		if l+1 == r {
			break
		}
		b--
		for x := r - 1; x > l; x-- {
			result = append(result, matrix[b][x])
		}
		if t+1 == b {
			break
		}
		l++
		for y := b - 1; y > t; y-- {
			result = append(result, matrix[y][l])
		}
		if l+1 == r {
			break
		}
	}
	return result
}

func TestSpiralOrder(t *testing.T) {
	output := spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	})
	expected = []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
