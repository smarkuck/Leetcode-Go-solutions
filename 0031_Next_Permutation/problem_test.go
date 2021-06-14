package leetcode

import (
	"reflect"
	"testing"
)

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	if i >= 0 {
		j := len(nums) - 1
		for ; nums[j] <= nums[i]; j-- {
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums, i+1)
}

func reverse(nums []int, start int) {
	for stop := len(nums) - 1; start < stop; start, stop = start+1, stop-1 {
		nums[start], nums[stop] = nums[stop], nums[start]
	}
}

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	expected := []int{1, 3, 2}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", nums, expected)
	}

	nums = []int{3, 2, 1}
	nextPermutation(nums)
	expected = []int{1, 2, 3}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", nums, expected)
	}

	nums = []int{1, 1, 5}
	nextPermutation(nums)
	expected = []int{1, 5, 1}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", nums, expected)
	}

	nums = []int{1}
	nextPermutation(nums)
	expected = []int{1}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", nums, expected)
	}
}
