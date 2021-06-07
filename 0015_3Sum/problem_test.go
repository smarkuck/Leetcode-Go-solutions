package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for i, num := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		target, left, right := -num, i+1, len(nums)-1
		for left < right {
			num2, num3 := nums[left], nums[right]
			sum := num2 + num3
			if sum == target {
				result = append(result, []int{num, num2, num3})
				left, right = left+1, right-1
				for ; left < right && nums[left-1] == nums[left]; left++ {
				}
				for ; left < right && nums[right+1] == nums[right]; right-- {
				}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func TestThreeSum(t *testing.T) {
	output := threeSum([]int{-1, 0, 1, 2, -1, -4})
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = threeSum([]int{})
	expected = [][]int{}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = threeSum([]int{0})
	expected = [][]int{}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
