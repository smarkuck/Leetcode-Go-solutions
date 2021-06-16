package leetcode

import (
	"reflect"
	"testing"
)

type Temperature struct {
	index int
	value int
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []Temperature{}
	for i, t := range temperatures {
		for len(stack) > 0 && stack[len(stack)-1].value < t {
			result[stack[len(stack)-1].index] = i - stack[len(stack)-1].index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, Temperature{index: i, value: t})
	}
	return result
}

func TestDailyTemperatures(t *testing.T) {
	output := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	expected := []int{1, 1, 4, 2, 1, 1, 0, 0}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = dailyTemperatures([]int{30, 40, 50, 60})
	expected = []int{1, 1, 1, 0}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = dailyTemperatures([]int{30, 60, 90})
	expected = []int{1, 1, 0}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
