package leetcode

import "testing"

func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	sums := map[int]int{0: 1}
	for _, num := range nums {
		sum += num
		count += sums[sum-k]
		sums[sum]++
	}
	return count
}

func TestSubarraySum(t *testing.T) {
	output := subarraySum([]int{1, 1, 1}, 2)
	expected := 2
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = subarraySum([]int{1, 2, 3}, 3)
	expected = 2
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
