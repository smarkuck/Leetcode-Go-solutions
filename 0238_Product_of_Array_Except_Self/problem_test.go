package leetcode

import (
	"reflect"
	"testing"
)

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}
	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i], r = ans[i]*r, r*nums[i]
	}
	return ans
}

func TestProductExceptSelf(t *testing.T) {
	output := productExceptSelf([]int{1, 2, 3, 4})
	expected := []int{24, 12, 8, 6}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = productExceptSelf([]int{-1, 1, 0, -3, 3})
	expected = []int{0, 0, 9, 0, 0}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
