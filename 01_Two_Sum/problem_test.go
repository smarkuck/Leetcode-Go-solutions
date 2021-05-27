package leetcode

import (
	"reflect"
	"testing"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return []int{0, 0}
}

func TestTwoSum(t *testing.T) {
	output := twoSum([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = twoSum([]int{3, 2, 4}, 6)
	expected = []int{1, 2}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = twoSum([]int{3, 3}, 6)
	expected = []int{0, 1}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
