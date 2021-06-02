package leetcode

import "testing"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	prevMax1, max1 := 0, 0
	prevMax2, max2 := 0, 0

	for i := 0; i < len(nums)-1; i++ {
		prevMax1, max1 = max1, max(max1, nums[i]+prevMax1)
		prevMax2, max2 = max2, max(max2, nums[i+1]+prevMax2)
	}

	return max(max1, max2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestRob(t *testing.T) {
	output := rob([]int{2, 3, 2})
	expected := 3
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = rob([]int{1, 2, 3, 1})
	expected = 4
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = rob([]int{0})
	expected = 0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
