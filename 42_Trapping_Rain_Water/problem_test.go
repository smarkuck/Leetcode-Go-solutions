package leetcode

import "testing"

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	sum := 0
	for left <= right {
		if leftMax < rightMax {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				sum += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				sum += rightMax - height[right]
			}
			right--
		}
	}
	return sum
}

func TestTrap(t *testing.T) {
	output := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	expected := 6
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = trap([]int{4, 2, 0, 3, 2, 5})
	expected = 9
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
