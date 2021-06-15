package leetcode

import "testing"

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) >> 1
		if nums[m] == target {
			return m
		} else if nums[m] >= nums[l] {
			if target >= nums[l] && nums[m] > target {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target <= nums[r] && nums[m] < target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}

func TestSearch(t *testing.T) {
	output := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	expected := 4
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = search([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	expected = -1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = search([]int{1}, 0)
	expected = -1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
