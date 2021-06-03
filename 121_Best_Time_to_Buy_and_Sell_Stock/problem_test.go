package leetcode

import "testing"

func maxProfit(prices []int) int {
	tmp, max := 0, 0
	for i := 1; i < len(prices); i++ {
		tmp += prices[i] - prices[i-1]
		if tmp < 0 {
			tmp = 0
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func TestMaxProfit(t *testing.T) {
	output := maxProfit([]int{7, 1, 5, 3, 6, 4})
	expected := 5
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = maxProfit([]int{7, 6, 4, 3, 1})
	expected = 0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
