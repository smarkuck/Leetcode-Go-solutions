package leetcode

import (
	"sort"
	"testing"
)

func deleteAndEarn(nums []int) int {
	sort.Ints(nums)
	t1, t2, next := 0, 0, 0
	for i := range nums {
		if next <= nums[i] {
			t1 = max(t1, t2)
			if next == nums[i] {
				t1, t2 = t2, t1
			} else {
				t2 = t1
			}
			next = nums[i] + 1
		}
		t1 += nums[i]
	}
	return max(t1, t2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestDeleteAndEarn(t *testing.T) {
	output := deleteAndEarn([]int{3, 4, 2})
	expected := 6
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = deleteAndEarn([]int{2, 2, 3, 3, 3, 4})
	expected = 9
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
