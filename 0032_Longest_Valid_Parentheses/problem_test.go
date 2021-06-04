package leetcode

import (
	"testing"
)

func longestValidParentheses(s string) int {
	l, r, max := 0, 0, 0
	for _, c := range s {
		if c == '(' {
			l++
		}
		if c == ')' {
			r++
		}
		if l == r && l+r > max {
			max = l + r
		} else if r > l {
			l, r = 0, 0
		}
	}

	l, r = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			l++
		}
		if s[i] == ')' {
			r++
		}
		if l == r && l+r > max {
			max = l + r
		} else if r < l {
			l, r = 0, 0
		}
	}
	return max
}

func TestLongestValidParentheses(t *testing.T) {
	output := longestValidParentheses("(()")
	expected := 2
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = longestValidParentheses(")()())")
	expected = 4
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = longestValidParentheses("")
	expected = 0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
