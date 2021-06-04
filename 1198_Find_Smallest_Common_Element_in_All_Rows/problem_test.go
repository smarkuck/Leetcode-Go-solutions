package leetcode

import "testing"

func smallestCommonElement(mat [][]int) int {
	nRow, nCol := len(mat), len(mat[0])
	indexes := make([]int, nRow)
	common := mat[0][0]
	matchCount := 0
	for {
		for i, row := range mat {
			indexes[i] = findNextIndex(row, indexes[i], common)
			if indexes[i] >= nCol {
				return -1
			}
			if row[indexes[i]] > common {
				common = row[indexes[i]]
				matchCount = 1
			} else {
				matchCount++
				if matchCount == nRow {
					return common
				}
			}
		}
	}
}

func findNextIndex(row []int, pos int, val int) int {
	d := 1
	for pos < len(row) && row[pos] < val {
		d <<= 1
		if row[min(pos+d, len(row)-1)] >= val {
			d = 1
		}
		pos += d
	}
	return pos
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestSmallestCommonElement(t *testing.T) {
	output := smallestCommonElement([][]int{
		{1, 2, 3, 4, 5},
		{2, 4, 5, 8, 10},
		{3, 5, 7, 9, 11},
		{1, 3, 5, 7, 9}})

	expected := 5
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = smallestCommonElement([][]int{
		{1, 2, 3},
		{2, 3, 4},
		{2, 3, 5}})

	expected = 2
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
