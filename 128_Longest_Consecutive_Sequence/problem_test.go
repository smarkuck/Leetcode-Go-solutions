package leetcode

import "testing"

func longestConsecutive(nums []int) int {
	longest := 0
	numsSet := make(map[int]bool)
	for _, num := range nums {
		numsSet[num] = true
	}
	for num := range numsSet {
		if !numsSet[num-1] {
			current := 1
			for num++; numsSet[num]; num++ {
				current++
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}

func TestLongestConsecutive(t *testing.T) {
	output := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	expected := 4
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	expected = 9
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
