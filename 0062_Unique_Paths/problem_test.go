package leetcode

import "testing"

func uniquePaths(m int, n int) int {
	m, n = m-1, n-1
	if m < n {
		m, n = n, m
	}
	result := 1
	for i := m + n; i > m; i-- {
		result *= i
	}
	for i := 2; i <= n; i++ {
		result /= i
	}
	return result
}

func TestUniquePaths(t *testing.T) {
	output := uniquePaths(3, 7)
	expected := 28
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = uniquePaths(3, 2)
	expected = 3
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = uniquePaths(7, 3)
	expected = 28
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = uniquePaths(3, 3)
	expected = 6
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
