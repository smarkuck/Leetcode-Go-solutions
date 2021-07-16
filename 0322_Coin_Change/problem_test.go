package leetcode

import "testing"

func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = max
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin > i {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == max {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestCoinChange(t *testing.T) {
	output := coinChange([]int{1, 2, 5}, 11)
	expected := 3
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = coinChange([]int{2}, 3)
	expected = -1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = coinChange([]int{1}, 0)
	expected = 0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = coinChange([]int{1}, 1)
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = coinChange([]int{1}, 2)
	expected = 2
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
