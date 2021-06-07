package leetcode

import "testing"

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i < len(dp); i++ {
			dp[i] = dp[i] + dp[i-coin]
		}
	}
	return dp[amount]
}

func TestChange(t *testing.T) {
	output := change(5, []int{1, 2, 5})
	expected := 4
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = change(3, []int{2})
	expected = 0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = change(10, []int{10})
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
