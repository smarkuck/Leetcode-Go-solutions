package leetcode

import (
	"reflect"
	"testing"
)

func getRow(rowIndex int) []int {
	row := []int{1}
	for i := 0; i < rowIndex; i++ {
		oldPrev := row[0]
		for j := 1; j < len(row); j++ {
			newPrev := row[j]
			row[j] = oldPrev + row[j]
			oldPrev = newPrev
		}
		row = append(row, 1)
	}
	return row
}

func TestGetRow(t *testing.T) {
	output := getRow(3)
	expected := []int{1, 3, 3, 1}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = getRow(0)
	expected = []int{1}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = getRow(1)
	expected = []int{1, 1}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
