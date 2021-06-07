package leetcode

import (
	"reflect"
	"testing"
)

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}

func TestTwoSum(t *testing.T) {
	output := twoSum([]int{2, 7, 11, 15}, 9)
	expected := []int{1, 2}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = twoSum([]int{2, 3, 4}, 6)
	expected = []int{1, 3}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = twoSum([]int{-1, 0}, -1)
	expected = []int{1, 2}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
