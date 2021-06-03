package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{intervals[0]}
	for _, interval := range intervals {
		if interval[0] <= result[len(result)-1][1] {
			if interval[1] > result[len(result)-1][1] {
				result[len(result)-1][1] = interval[1]
			}
		} else {
			result = append(result, interval)
		}
	}
	return result
}

func TestMerge(t *testing.T) {
	output := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = merge([][]int{{1, 4}, {4, 5}})
	expected = [][]int{{1, 5}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
