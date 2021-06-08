package leetcode

import "testing"

func maxArea(height []int) int {
	maxArea, l, r := 0, 0, len(height)-1
	for l < r {
		maxArea = max(maxArea, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMaxArea(t *testing.T) {
	output := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	expected := 49
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = maxArea([]int{1, 1})
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = maxArea([]int{4, 3, 2, 1, 4})
	expected = 16
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = maxArea([]int{1, 2, 1})
	expected = 2
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
