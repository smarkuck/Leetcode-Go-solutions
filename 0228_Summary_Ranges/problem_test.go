package leetcode

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func summaryRanges(nums []int) []string {
	res := []string{}
	head := 0
	for i := range nums {
		if i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			continue
		}
		if head == i {
			res = append(res, strconv.Itoa(nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%v->%v", nums[head], nums[i]))
		}
		head = i + 1
	}
	return res
}

func TestSummaryRanges(t *testing.T) {
	output := summaryRanges([]int{0, 1, 2, 4, 5, 7})
	expected := []string{"0->2", "4->5", "7"}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})
	expected = []string{"0", "2->4", "6", "8->9"}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = summaryRanges([]int{})
	expected = []string{}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = summaryRanges([]int{-1})
	expected = []string{"-1"}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = summaryRanges([]int{0})
	expected = []string{"0"}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
