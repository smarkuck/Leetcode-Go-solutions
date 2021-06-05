package leetcode

import "testing"

func knightDialer(n int) int {
	if n == 1 {
		return 10
	}

	current := []int{1, 1, 1, 1, 1, 0, 1, 1, 1, 1}
	next := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 1; i < n; i++ {
		next[0] = (current[4] + current[6]) % 1000000007
		next[1] = (current[6] + current[8]) % 1000000007
		next[2] = (current[7] + current[9]) % 1000000007
		next[3] = (current[4] + current[8]) % 1000000007
		next[4] = (current[3] + current[9] + current[0]) % 1000000007
		next[6] = (current[1] + current[7] + current[0]) % 1000000007
		next[7] = (current[2] + current[6]) % 1000000007
		next[8] = (current[1] + current[3]) % 1000000007
		next[9] = (current[2] + current[4]) % 1000000007
		current, next = next, current
	}

	sum := 0
	for _, v := range current {
		sum += v
	}
	return sum % 1000000007
}

func TestKnightDialer(t *testing.T) {
	output := knightDialer(1)
	expected := 10
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = knightDialer(2)
	expected = 20
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = knightDialer(3)
	expected = 46
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = knightDialer(4)
	expected = 104
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = knightDialer(3131)
	expected = 136006598
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
